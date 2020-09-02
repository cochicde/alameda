/*
Copyright 2020 The Alameda Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	datahub_client_application "github.com/containers-ai/alameda/operator/datahub/client/application"
	datahub_namespace "github.com/containers-ai/alameda/operator/datahub/client/namespace"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	datahub_controller "github.com/containers-ai/alameda/operator/datahub/client/controller"
	datahub_pod "github.com/containers-ai/alameda/operator/datahub/client/pod"
	alamedascaler_reconciler "github.com/containers-ai/alameda/operator/pkg/reconciler/alamedascaler"
	"github.com/containers-ai/alameda/operator/pkg/utils"
	utilsresource "github.com/containers-ai/alameda/operator/pkg/utils/resources"
	alamutils "github.com/containers-ai/alameda/pkg/utils"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"

	corev1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var listCandidatesDefaultAlamedaScaler = func(
	ctx context.Context,
	k8sClient client.Client,
	objectMeta metav1.ObjectMeta,
) ([]autoscalingv1alpha1.AlamedaScaler, error) {

	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err := k8sClient.List(ctx, &alamedaScalerList,
		&client.ListOptions{Namespace: objectMeta.Namespace})
	if err != nil {
		return nil, errors.Wrap(err, "list AlamedaScalers failed")
	}

	candidates := make([]autoscalingv1alpha1.AlamedaScaler, 0)
	for _, alamedaScaler := range alamedaScalerList.Items {
		if !(alamedaScaler.GetType() == autoscalingv1alpha1.AlamedaScalerTypeNotDefine ||
			alamedaScaler.GetType() == autoscalingv1alpha1.AlamedaScalerTypeDefault) {
			continue
		}
		if alamedaScaler.Spec.Selector == nil {
			continue
		}
		if ok := isLabelsSelectedBySelector(*alamedaScaler.Spec.Selector,
			objectMeta.GetLabels()); ok {
			candidates = append(candidates, alamedaScaler)
		}
	}
	return candidates, nil
}

func init() {
	RegisterAlamedaScalerController(
		autoscalingv1alpha1.AlamedaScalerTypeDefault, listCandidatesDefaultAlamedaScaler)
}

var (
	scope = logUtil.RegisterScope("operator_controllers", "operator controllers", 0)

	onceCheckHasOpenshiftAPIAppsV1 = sync.Once{}
	hasOpenshiftAPIAppsV1          = false

	requeueAfter = 3 * time.Second
)

var alamedascalerFirstSynced = false

// AlamedaScalerReconciler reconciles a AlamedaScaler object
type AlamedaScalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	ClusterUID            string
	DatahubClient         *datahubpkg.Client
	DatahubControllerRepo *datahub_controller.ControllerRepository
	DatahubPodRepo        *datahub_pod.PodRepository
	//onceForceReconcile     sync.Once
	ReconcileTimeout       time.Duration
	ForceReconcileInterval time.Duration
}

// Reconcile reads that state of the cluster for a AlamedaScaler object and makes changes based on the state read
// and what is in the AlamedaScaler .Spec
func (r *AlamedaScalerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	if !alamedascalerFirstSynced {
		time.Sleep(5 * time.Second)
	}
	alamedascalerFirstSynced = true

	onceCheckHasOpenshiftAPIAppsV1.Do(
		func() {
			exist, err := utils.ServerHasOpenshiftAPIAppsV1()
			if err != nil {
				panic(errors.Wrap(err, "Check if apiServer has openshift apps v1 api failed"))
			}
			hasOpenshiftAPIAppsV1 = exist
		})

	// Delete resources relative to AlamedaScaler
	ctx := context.TODO()
	instance := autoscalingv1alpha1.AlamedaScaler{}
	err := r.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.Name}, &instance)
	if err != nil && k8sErrors.IsNotFound(err) {
		scope.Infof("Handling deletion of AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
		if err := r.handleAlamedaScalerDeletion(req.Namespace, req.Name); err != nil {
			scope.Warnf("Handle deletion of AlamedaScaler(%s/%s) failed, retry reconciling: %s",
				req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		scope.Infof("Handle deletion of AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
		return ctrl.Result{Requeue: false}, nil
	} else if err != nil {
		scope.Errorf("Get AlamedaScaler(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	err = r.DatahubClient.Create(&[]entities.ResourceClusterStatusNamespace{
		{
			Name:        req.Namespace,
			ClusterName: r.ClusterUID,
		},
	})
	if err != nil {
		scope.Errorf("create namespace for alamedascaler (%s/%s) failed: %s",
			req.Namespace, req.Name, err.Error())
	}

	alamedaScaler := autoscalingv1alpha1.AlamedaScaler{}
	instance.DeepCopyInto(&alamedaScaler)

	if !r.isAlamedaScalerNeedToBeReconciled(context.TODO(), alamedaScaler) {
		scope.Infof("AlamedaScale(%s/%s) type(%s), skip reconciling.", req.Namespace, req.Name, alamedaScaler.GetType())
		if err := r.syncDatahubApplicationsByAlamedaScaler(ctx, alamedaScaler); err != nil {
			scope.Errorf("sync alamedascaler(%s/%s) with datahub failed.", req.Namespace, req.Name)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		return ctrl.Result{Requeue: false}, nil
	}

	if alamedaScaler.GetDeletionTimestamp() != nil {
		scope.Infof("Handling deletion of AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
		if err := r.handleAlamedaScalerDeletion(alamedaScaler.Namespace, alamedaScaler.Name); err != nil {
			scope.Warnf("Handle deletion of AlamedaScaler(%s/%s) failed, retry reconciling: %s",
				req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		scope.Infof("Handle deletion of AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
		return ctrl.Result{Requeue: false}, nil
	}

	scope.Infof("Reconciling AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
	alamedaScaler = r.setDefaultAlamedaScaler(alamedaScaler)
	alamedaScaler.Status.AlamedaController = autoscalingv1alpha1.NewAlamedaController()
	if !alamedaScaler.IsScalingToolTypeCA() {
		if alamedaScaler, err = r.listAndAddDeploymentsIntoAlamedaScalerStatue(context.TODO(),
			alamedaScaler); err != nil {
			scope.Warnf("List and add Deployments into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v",
				req.Namespace, req.Name, requeueAfter.Seconds(), err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		if hasOpenshiftAPIAppsV1 {
			if alamedaScaler, err = r.listAndAddDeploymentConfigsIntoAlamedaScalerStatus(context.TODO(),
				alamedaScaler); err != nil {
				scope.Warnf("List and add DeploymentConfigs into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v",
					req.Namespace, req.Name, requeueAfter.Seconds(), err)
				return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
			}
		}
		if alamedaScaler, err = r.listAndAddStatefulSetsIntoAlamedaScalerStatue(context.TODO(),
			alamedaScaler); err != nil {
			scope.Warnf("List and add StatefulSets into AlamedaScaler(%s/%s) failed, retry after %f seconds: %+v",
				req.Namespace, req.Name, requeueAfter.Seconds(), err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
	} else {
		mgName := alamedaScaler.Spec.ScalingTool.MachineGroupScaler
		mgIns := autoscalingv1alpha1.AlamedaMachineGroupScaler{}
		err := r.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: mgName}, &mgIns)
		if err != nil && k8sErrors.IsNotFound(err) {
			scope.Errorf("AlamedaMachineGroupScaler (%s/%s) does not exist in alamedascaler (%s/%s)",
				mgIns.GetNamespace(), mgIns.GetName(), req.Namespace, req.Name)
			return ctrl.Result{Requeue: false}, nil
		} else if err != nil {
			scope.Errorf("Get AlamedaMachineGroupScaler(%s/%s) failed: %s",
				req.Namespace, mgName, err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}

		if err = r.syncCAInfoWithScalerAndMachineGroup(ctx, alamedaScaler, mgIns); err != nil {
			scope.Errorf("sync CA info with alamedascaler (%s/%s) and machinegroupscaler (%s/%s) failed: %s",
				alamedaScaler.GetNamespace(), alamedaScaler.GetName(), mgIns.GetNamespace(), mgIns.GetName(),
				err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
	}

	if err := r.Update(context.TODO(), &alamedaScaler); err != nil {
		scope.Errorf("Update AlamedaScaler(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	// after updating AlamedaPod in AlamedaScaler, start create AlamedaRecommendation if necessary and register alameda pod to datahub
	scope.Debugf("Start syncing AlamedaScaler(%s/%s) to datahub. %s",
		req.Namespace, req.Name, alamutils.InterfaceToString(alamedaScaler))
	if err := r.syncAlamedaScalerWithDepResources(&alamedaScaler); err != nil {
		scope.Error(err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	scope.Infof("Reconcile AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
	return ctrl.Result{}, nil
}

func (r AlamedaScalerReconciler) setDefaultAlamedaScaler(
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) autoscalingv1alpha1.AlamedaScaler {
	alamedaScaler.SetDefaultValue()
	alamedaScaler.Spec.Type = autoscalingv1alpha1.AlamedaScalerTypeDefault
	return alamedaScaler
}

func (r AlamedaScalerReconciler) isAlamedaScalerNeedToBeReconciled(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) bool {
	return alamedaScaler.GetType() == autoscalingv1alpha1.AlamedaScalerTypeNotDefine ||
		alamedaScaler.GetType() == autoscalingv1alpha1.AlamedaScalerTypeDefault
}

func (r AlamedaScalerReconciler) isWorkloadControllerCanBeMonitoredByAlamedaScaler(
	ctx context.Context, workloadController metav1.ObjectMeta,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) (bool, error) {
	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err := r.List(ctx, &alamedaScalerList, &client.ListOptions{Namespace: workloadController.GetNamespace()})
	if err != nil {
		return false, errors.Wrap(err, "list AlamedaScalers failed")
	}

	alamedaScalersObjectMeta := make([]metav1.ObjectMeta, 0, len(alamedaScalerList.Items))
	for _, alamedaScaler := range alamedaScalerList.Items {
		if alamedaScaler.Spec.Selector == nil {
			continue
		}
		if ok := isLabelsSelectedBySelector(*alamedaScaler.Spec.Selector,
			workloadController.GetLabels()); !ok {
			continue
		}
		alamedaScalersObjectMeta = append(alamedaScalersObjectMeta, alamedaScaler.ObjectMeta)
	}

	oldestAlamedaScaler := getFirstCreatedObjectMeta(alamedaScalersObjectMeta)
	if oldestAlamedaScaler.GetNamespace() != alamedaScaler.GetNamespace() ||
		oldestAlamedaScaler.GetName() != alamedaScaler.GetName() {
		return false, nil
	}
	return true, nil
}

func (r AlamedaScalerReconciler) listAndAddDeploymentsIntoAlamedaScalerStatue(ctx context.Context,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) (autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(r.Client, &alamedaScaler)
	deployments, err := listResources.ListDeploymentsByNamespaceLabels(alamedaScaler.Namespace,
		alamedaScaler.Spec.Selector.MatchLabels)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list Deployments failed")
	}
	for _, deployment := range deployments {
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(ctx, deployment.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if Deployment can be monitored failed")
		}

		ok = ok && IsMonitoredByAlamedaScalerController(deployment.ObjectMeta,
			autoscalingv1alpha1.AlamedaScalerTypeDefault)
		if ok {
			_, err = alamedascalerReconciler.UpdateStatusByDeployment(&deployment)
			if err != nil {
				return alamedaScaler, errors.Wrap(err, "update status of AlamedaScaler(%s/%s) by Deployment failed")
			}
		} else {
			ar := r.getIneffectiveAlamedaResource(deployment.ObjectMeta)
			alamedaScaler.AddAlamedaResourceIntoStatus(autoscalingv1alpha1.DeploymentController, ar)
		}
	}
	return alamedaScaler, nil
}

func (r AlamedaScalerReconciler) listAndAddDeploymentConfigsIntoAlamedaScalerStatus(ctx context.Context,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) (autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(r.Client, &alamedaScaler)
	dploymentConfigs, err := listResources.ListDeploymentConfigsByNamespaceLabels(alamedaScaler.Namespace,
		alamedaScaler.Spec.Selector.MatchLabels)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list Deployments failed")
	}
	for _, deploymentConfig := range dploymentConfigs {
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(context.TODO(),
			deploymentConfig.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if DeploymentConfig can be monitored failed")
		}

		ok = ok && IsMonitoredByAlamedaScalerController(deploymentConfig.ObjectMeta,
			autoscalingv1alpha1.AlamedaScalerTypeDefault)
		if ok {
			_, err = alamedascalerReconciler.UpdateStatusByDeploymentConfig(&deploymentConfig)
			if err != nil {
				return alamedaScaler, errors.Wrap(err, "update status of AlamedaScaler(%s/%s) by DeploymentConfig failed")
			}
		} else {
			ar := r.getIneffectiveAlamedaResource(deploymentConfig.ObjectMeta)
			alamedaScaler.AddAlamedaResourceIntoStatus(autoscalingv1alpha1.DeploymentConfigController, ar)
		}
	}
	return alamedaScaler, nil
}

func (r AlamedaScalerReconciler) listAndAddStatefulSetsIntoAlamedaScalerStatue(ctx context.Context,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) (autoscalingv1alpha1.AlamedaScaler, error) {
	listResources := utilsresource.NewListResources(r.Client)
	alamedascalerReconciler := alamedascaler_reconciler.NewReconciler(r.Client, &alamedaScaler)
	statefulSets, err := listResources.ListStatefulSetsByNamespaceLabels(alamedaScaler.Namespace,
		alamedaScaler.Spec.Selector.MatchLabels)
	if err != nil {
		return alamedaScaler, errors.Wrap(err, "list Deployments failed")
	}
	for _, statefulSet := range statefulSets {
		ok, err := r.isWorkloadControllerCanBeMonitoredByAlamedaScaler(context.TODO(),
			statefulSet.ObjectMeta, alamedaScaler)
		if err != nil {
			return alamedaScaler, errors.Wrap(err, "check if StatefulSet can be monitored failed")
		}

		ok = ok && IsMonitoredByAlamedaScalerController(statefulSet.ObjectMeta,
			autoscalingv1alpha1.AlamedaScalerTypeDefault)
		if ok {
			_, err = alamedascalerReconciler.UpdateStatusByStatefulSet(&statefulSet)
			if err != nil {
				return alamedaScaler, errors.Wrap(err, "update status of AlamedaScaler(%s/%s) by StatefulSet failed")
			}
		} else {
			ar := r.getIneffectiveAlamedaResource(statefulSet.ObjectMeta)
			alamedaScaler.AddAlamedaResourceIntoStatus(autoscalingv1alpha1.StatefulSetController, ar)
		}
	}
	return alamedaScaler, nil
}

func (r AlamedaScalerReconciler) getIneffectiveAlamedaResource(
	workloadController metav1.ObjectMeta) autoscalingv1alpha1.AlamedaResource {
	empty := int32(0)
	return autoscalingv1alpha1.AlamedaResource{
		Namespace:    workloadController.GetNamespace(),
		Name:         workloadController.GetName(),
		UID:          string(workloadController.GetUID()),
		SpecReplicas: &empty,
		Effective:    false,
		Message:      "Is monitoring by other AlamedaScaler.",
	}
}

func (r *AlamedaScalerReconciler) syncAlamedaScalerWithDepResources(
	alamedaScaler *autoscalingv1alpha1.AlamedaScaler) error {

	existingPodsMap := make(map[autoscalingv1alpha1.NamespacedName]bool)
	existingPods := alamedaScaler.GetMonitoredPods()
	for _, pod := range existingPods {
		existingPodsMap[pod.GetNamespacedName()] = true
	}

	wg := errgroup.Group{}
	wg.Go(func() error {
		return r.syncDatahubResourceByAlamedaScaler(context.TODO(), *alamedaScaler)
	})
	if err := wg.Wait(); err != nil {
		return errors.Wrapf(err, "sync AlamedaScaler %s/%s with dependent resources failed",
			alamedaScaler.Namespace, alamedaScaler.Name)
	}

	return nil
}

func (r *AlamedaScalerReconciler) syncDatahubResourceByAlamedaScaler(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) error {

	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		if err := r.syncDatahubApplicationsByAlamedaScaler(ctx, alamedaScaler); err != nil {
			return errors.Wrap(err, "sync applications with Datahub failed")
		}
		return nil
	})
	wg.Go(func() error {
		if err := r.syncDatahubControllersByAlamedaScaler(ctx, alamedaScaler); err != nil {
			return errors.Wrap(err, "sync controllers with Datahub failed")
		}
		return nil
	})
	wg.Go(func() error {
		if err := r.syncDatahubPodsByAlamedaScaler(ctx, alamedaScaler); err != nil {
			return errors.Wrap(err, "sync pods with Datahub failed")
		}
		return nil
	})
	return wg.Wait()
}

func (r *AlamedaScalerReconciler) syncDatahubApplicationsByAlamedaScaler(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) error {

	namespace := alamedaScaler.Namespace
	name := alamedaScaler.Name
	scalingToolStr := datahub_client_application.GetAlamedaScalerDatahubScalingTypeStr(alamedaScaler)
	scope.Debugf(
		"Creating applications to datahub. AlamedaScaler(application): %s/%s. Scaling tool: %s",
		namespace, name, scalingToolStr)
	entity := entities.ResourceClusterStatusApplication{
		ClusterName: r.ClusterUID,
		Namespace:   namespace,
		Name:        name,
	}

	err := r.DatahubClient.Create(&[]entities.ResourceClusterStatusApplication{
		entity,
	})
	if err != nil {
		return errors.Wrapf(err, "create Application(%s/%s) to Datahub failed", namespace, name)
	}
	scope.Debugf(
		"Create applications to datahub success. AlamedaScaler: %s/%s. Scaling tool: %s",
		namespace, name, scalingToolStr)
	return nil
}
func (r *AlamedaScalerReconciler) syncDatahubControllersByAlamedaScaler(ctx context.Context,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) error {

	namespace := alamedaScaler.Namespace
	name := alamedaScaler.Name
	if err := r.createAlamedaWatchedResourcesToDatahub(&alamedaScaler); err != nil {
		return errors.Wrapf(err, "create AlamedaScaler(%s/%s) watched resources to datahub failed",
			namespace, name)
	}

	// list all controller with namespace same as alamedaScaler
	controllers, err := r.listAlamedaWatchedResourcesToDatahub(&alamedaScaler)
	if err != nil {
		return errors.Wrapf(err, "list AlamedaScaler(%s/%s) watched resources to datahub failed",
			namespace, name)
	}
	err = r.deleteAlamedaWatchedResourcesToDatahub(context.TODO(), &alamedaScaler, controllers)
	if err != nil {
		return errors.Wrapf(err, "delete AlamedaScaler(%s/%s) watched resources to datahub failed",
			namespace, name)
	}
	return nil
}

func (r *AlamedaScalerReconciler) syncDatahubPodsByAlamedaScaler(ctx context.Context,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) error {

	// // When AlamedaScaler type is not vpa, delete all pods monitored by the AlamedaScaler from Datahub
	// if alamedaScaler.Spec.ScalingTool.Type != autoscalingv1alpha1.ScalingToolTypeVPA {
	// 	if err := r.deletePodsFromDatahubByAlamedaScaler(ctx, alamedaScaler.Namespace, alamedaScaler.Name); err != nil {
	// 		return errors.Wrap(err, "delete pods from Datahub by AlamedaScaler failed")
	// 	}
	// 	return nil
	// }

	// Delete pods from Datahub which are no longer monitered by AlamedaScaler.
	monitoringPodMap := map[string]bool{}
	for _, pod := range alamedaScaler.GetMonitoredPods() {
		if pod == nil {
			continue
		}
		namespace := pod.Namespace
		name := pod.Name
		monitoringPodMap[fmt.Sprintf("%s/%s", namespace, name)] = true
	}
	podsNeedToBeDeleted := make([]*datahub_resources.ObjectMeta, 0)
	pods, err := r.DatahubPodRepo.ListAlamedaPodsByAlamedaScaler(context.TODO(),
		alamedaScaler.Namespace, alamedaScaler.Name)
	if err != nil {
		return errors.Wrapf(err, "list pods monitored by AlamedaScaler(%s/%s) from Datahub failed",
			alamedaScaler.Namespace, alamedaScaler.Name)
	}
	for _, pod := range pods {
		if pod == nil || pod.ObjectMeta == nil {
			continue
		}
		if ok, exist := monitoringPodMap[fmt.Sprintf("%s/%s",
			pod.ObjectMeta.Namespace, pod.ObjectMeta.Name)]; !ok || !exist {
			podsNeedToBeDeleted = append(podsNeedToBeDeleted, pod.ObjectMeta)
		}
	}
	if len(podsNeedToBeDeleted) == 0 {
		return nil
	}
	scope.Debugf("Deleting pods from datahub. AlamedaScaler: %s/%s. Pods: %+v",
		alamedaScaler.GetNamespace(), alamedaScaler.GetName(), podsNeedToBeDeleted)
	if err := r.DatahubPodRepo.DeletePods(context.TODO(), podsNeedToBeDeleted); err != nil {
		return errors.Wrap(err, "delete pods from Datahub failed")
	}
	scope.Debugf("Delete pods from datahub success. AlamedaScaler: %s/%s. Pods: %+v",
		alamedaScaler.GetNamespace(), alamedaScaler.GetName(), podsNeedToBeDeleted)

	return nil
}

func (r *AlamedaScalerReconciler) listAlamedaWatchedResourcesToDatahub(
	scaler *autoscalingv1alpha1.AlamedaScaler) ([]*datahub_resources.Controller, error) {

	controllers, err := r.DatahubControllerRepo.ListControllersByApplication(context.TODO(),
		scaler.Namespace, scaler.Name)
	if err != nil {
		return nil, errors.Wrap(err, "list controllers by application from Datahub failed")
	}
	return controllers, nil
}

func (r *AlamedaScalerReconciler) createAlamedaWatchedResourcesToDatahub(scaler *autoscalingv1alpha1.AlamedaScaler) error {
	controllers := r.getDatahubControllersFromAlamedaScalerStatus(*scaler)
	scope.Debugf("Creating controllers to datahub. AlamedaScaler: %s/%s. Controllers: %+v",
		scaler.GetNamespace(), scaler.GetName(), controllers)
	err := r.DatahubClient.Create(&controllers)
	if err != nil {
		return err
	}
	scope.Debugf("Create controllers to datahub success. AlamedaScaler: %s/%s. Controllers: %+v",
		scaler.GetNamespace(), scaler.GetName(), controllers)
	return nil
}

func (r *AlamedaScalerReconciler) getDatahubControllersFromAlamedaScalerStatus(
	scaler autoscalingv1alpha1.AlamedaScaler) []entities.ResourceClusterStatusController {
	policy := datahub_resources.RecommendationPolicy_RECOMMENDATION_POLICY_UNDEFINED
	switch scaler.Spec.Policy {
	case autoscalingv1alpha1.RecommendationPolicySTABLE:
		policy = datahub_resources.RecommendationPolicy_STABLE
	case autoscalingv1alpha1.RecommendationPolicyCOMPACT:
		policy = datahub_resources.RecommendationPolicy_COMPACT
	}
	isScalerEnableExecution := scaler.IsEnableExecution()
	scalingTool := datahub_client_application.GetAlamedaScalerDatahubScalingTypeStr(scaler)

	datahubKindToAlamedaResourceMap := map[datahub_resources.Kind]map[autoscalingv1alpha1.NamespacedName]autoscalingv1alpha1.AlamedaResource{
		datahub_resources.Kind_DEPLOYMENT:       scaler.Status.AlamedaController.Deployments,
		datahub_resources.Kind_DEPLOYMENTCONFIG: scaler.Status.AlamedaController.DeploymentConfigs,
		datahub_resources.Kind_STATEFULSET:      scaler.Status.AlamedaController.StatefulSets,
	}
	controllers := []entities.ResourceClusterStatusController{}
	for kind, alamedaResourceMap := range datahubKindToAlamedaResourceMap {
		for _, alamedaResource := range alamedaResourceMap {
			if alamedaResource.Effective == false {
				continue
			}
			replicas := len(alamedaResource.Pods)
			specReplicas := int32(-1)
			minReplicas := int32(-1)
			maxReplicas := int32(-1)
			if alamedaResource.SpecReplicas != nil {
				specReplicas = *alamedaResource.SpecReplicas
			}
			if scaler.Spec.ScalingTool.MinReplicas != nil {
				minReplicas = *scaler.Spec.ScalingTool.MinReplicas
			}
			if scaler.Spec.ScalingTool.MaxReplicas != nil {
				maxReplicas = *scaler.Spec.ScalingTool.MaxReplicas
			}
			controllers = append(controllers, entities.ResourceClusterStatusController{
				Namespace:                alamedaResource.Namespace,
				Name:                     alamedaResource.Name,
				ClusterName:              r.ClusterUID,
				AlamedaScalerName:        scaler.Name,
				Kind:                     entities.Kind(datahub_resources.Kind_name[int32(kind)]),
				Policy:                   entities.Policy(datahub_resources.RecommendationPolicy_name[int32(policy)]),
				EnableExecution:          isScalerEnableExecution,
				AlamedaScalerScalingTool: entities.ScalingTool(scalingTool),
				Replicas:                 int32(replicas),
				SpecReplicas:             specReplicas,
				ResourceK8sMinReplicas:   minReplicas,
				ResourceK8sMaxReplicas:   maxReplicas,
			})
		}
	}
	return controllers
}

func (r *AlamedaScalerReconciler) deleteAlamedaWatchedResourcesToDatahub(ctx context.Context,
	scaler *autoscalingv1alpha1.AlamedaScaler, ctlrsFromDH []*datahub_resources.Controller) error {

	controllerMap := map[datahub_resources.Kind][]*datahub_resources.ObjectMeta{
		datahub_resources.Kind_DEPLOYMENT:       {},
		datahub_resources.Kind_DEPLOYMENTCONFIG: {},
		datahub_resources.Kind_STATEFULSET:      {},
	}
	for _, ctlr := range ctlrsFromDH {
		if !r.isControllerHasAlamedaScalerInfo(ctlr, *scaler) {
			continue
		}
		if r.isControllerExistsInAlamedaScalerStatus(ctlr, *scaler) {
			continue
		}
		controllerMap[ctlr.Kind] = append(controllerMap[ctlr.Kind], ctlr.ObjectMeta)
	}
	for kind, controllers := range controllerMap {
		if len(controllers) == 0 {
			continue
		}
		scope.Debugf("Deleting controllers from datahub. AlamedaScaler: %s/%s. Controllers: %+v",
			scaler.GetNamespace(), scaler.GetName(), controllers)
		err := r.DatahubControllerRepo.DeleteControllers(ctx, controllers, kind)
		if err != nil {
			return errors.Wrap(err, "delete controllers from Datahub failed")
		}
		scope.Debugf("Delete controllers from datahub success. AlamedaScaler: %s/%s. Controllers: %+v",
			scaler.GetNamespace(), scaler.GetName(), controllers)
	}
	return nil
}

func (r *AlamedaScalerReconciler) setResourceForContainers(
	corePod *corev1.Pod, containers []*datahub_resources.Container) {
	for _, podContainer := range corePod.Spec.Containers {
		for containerIdx := range containers {
			if podContainer.Name == containers[containerIdx].GetName() {
				r.setContainerResource(
					corePod, podContainer, containers[containerIdx])
				break
			}
		}
	}
}

func (r *AlamedaScalerReconciler) setContainerResource(corePod *corev1.Pod,
	podContainer corev1.Container, container *datahub_resources.Container) {
	for _, resourceType := range []corev1.ResourceName{
		corev1.ResourceCPU, corev1.ResourceMemory,
	} {
		mt := datahub_common.ResourceName_CPU
		if resourceType == corev1.ResourceCPU {
			mt = datahub_common.ResourceName_CPU
		} else if resourceType == corev1.ResourceMemory {
			mt = datahub_common.ResourceName_MEMORY
		}
		if &podContainer.Resources != nil && podContainer.Resources.Limits != nil {
			resVal, ok := podContainer.Resources.Limits[resourceType]
			if ok && mt == datahub_common.ResourceName_CPU {
				container.Resources.Limits[int32(mt)] =
					strconv.FormatInt(resVal.MilliValue(), 10)
			} else if ok && mt == datahub_common.ResourceName_MEMORY {
				container.Resources.Limits[int32(mt)] =
					strconv.FormatInt(resVal.Value(), 10)
			} else {
				scope.Errorf("Recommendation %s will not generate if container %s of pod (%s/%s) does not set %s limit",
					mt, container.GetName(), corePod.GetNamespace(),
					corePod.GetName(), mt)
			}
		} else {
			scope.Errorf("Recommendation will not generate if container %s of pod (%s/%s) does not set any limit",
				container.GetName(), corePod.GetNamespace(), corePod.GetName())
		}

		if &podContainer.Resources != nil && podContainer.Resources.Requests != nil {
			resVal, ok := podContainer.Resources.Requests[resourceType]
			if ok && mt == datahub_common.ResourceName_CPU {
				container.Resources.Requests[int32(mt)] =
					strconv.FormatInt(resVal.MilliValue(), 10)
			} else if ok && mt == datahub_common.ResourceName_MEMORY {
				container.Resources.Requests[int32(mt)] =
					strconv.FormatInt(resVal.Value(), 10)
			} else {
				scope.Errorf("Recommendation %s will not generate if container %s of pod (%s/%s) does not set %s request",
					mt, container.GetName(), corePod.GetNamespace(),
					corePod.GetName(), mt)
			}
		} else {
			scope.Errorf("Recommendation will not generate if container %s of pod (%s/%s) does not set any request",
				container.GetName(), corePod.GetNamespace(), corePod.GetName())
		}
	}
}

func (r *AlamedaScalerReconciler) handleAlamedaScalerDeletion(namespace, name string) error {

	ctx := context.TODO()
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
		err := r.List(ctx, &alamedaScalerList, client.InNamespace(namespace))
		if err != nil {
			return fmt.Errorf("list alamedascaler in namespace %s for alamedascaler (%s/%s) remove failed: %s",
				namespace, namespace, name, err.Error())
		}
		if datahub_namespace.IsNSExcluded(namespace, alamedaScalerList.Items) {
			err = r.DatahubClient.Delete(&[]entities.ResourceClusterStatusNamespace{
				{
					Name:        namespace,
					ClusterName: r.ClusterUID,
				},
			})
			if err != nil {
				scope.Errorf("delete namespace for alamedascaler (%s/%s) remove failed: %s",
					namespace, name, err.Error())
			}
		}
		return nil
	})
	wg.Go(func() error {
		if err := r.deleteControllersFromDatahubByAlamedaScaler(ctx, namespace, name); err != nil {
			return errors.Wrapf(err, "delete controllers from datahub by AlamedaScaler(%s/%s) failed",
				namespace, name)
		}
		return nil
	})
	wg.Go(func() error {
		if err := r.deletePodsFromDatahubByAlamedaScaler(ctx, namespace, name); err != nil {
			return errors.Wrapf(err, "delete pods from datahub by AlamedaScaler(%s/%s) failed",
				namespace, name)
		}
		return nil
	})
	wg.Go(func() error {
		scope.Debugf("Deleting applications from datahub. AlamedaScaler (application): %s/%s.",
			namespace, name)

		if err := r.DatahubClient.DeleteByOpts(&entities.ResourceClusterStatusApplication{}, datahubpkg.Option{
			Entity: entities.ResourceClusterStatusApplication{
				ClusterName: r.ClusterUID,
				Namespace:   namespace,
				Name:        name,
			},
			Fields: []string{"ClusterName", "Namespace", "Name"},
		}); err != nil {
			return errors.Wrapf(err, "delete Application(%s/%s) from datahub failed", namespace, name)
		}
		scope.Debugf("Delete applications from datahub success. AlamedaScaler(application): %s/%s.",
			namespace, name)
		return nil
	})

	return wg.Wait()
}

func (r *AlamedaScalerReconciler) deleteControllersFromDatahubByAlamedaScaler(
	ctx context.Context, namespace, name string) error {
	wg, _ := errgroup.WithContext(ctx)
	wg.Go(func() error {
		scope.Debugf("Deleting controllers of AlamedaScaler: %s/%s from datahub.",
			namespace, name)

		err := r.DatahubClient.DeleteByOpts(&entities.ResourceClusterStatusController{}, datahubpkg.Option{
			Entity: entities.ResourceClusterStatusController{
				ClusterName:       r.ClusterUID,
				Namespace:         namespace,
				AlamedaScalerName: name,
			},
			Fields: []string{"ClusterName", "Namespace", "AlamedaScalerName"},
		})
		if err != nil {
			return err
		}
		scope.Debugf("Deleting controllers of AlamedaScaler: %s/%s from datahub success.",
			namespace, name)
		return nil
	})
	return wg.Wait()
}

func (r *AlamedaScalerReconciler) deletePodsFromDatahubByAlamedaScaler(
	ctx context.Context, namespace, name string) error {

	pods, err := r.DatahubPodRepo.ListAlamedaPodsByAlamedaScaler(ctx, namespace, name)
	if err != nil {
		return errors.Wrapf(err, "list pods by AlamedaScaler(%s/%s) failed", namespace, name)
	}

	podsNeedDeleting := make([]*datahub_resources.ObjectMeta, len(pods))
	for i, pod := range pods {
		podsNeedDeleting[i] = pod.ObjectMeta
	}
	if len(podsNeedDeleting) == 0 {
		return nil
	}
	scope.Debugf("Deleting pods from datahub. AlamedaScaler: %s/%s. Pods: %+v",
		namespace, name, podsNeedDeleting)
	err = r.DatahubPodRepo.DeletePods(ctx, podsNeedDeleting)
	if err != nil {
		return errors.Wrapf(err, "delete pods by AlamedaScaler(%s/%s) failed",
			namespace, name)
	}
	scope.Debugf("Delete pods from datahub success. AlamedaScaler: %s/%s. Pods: %+v",
		namespace, name, podsNeedDeleting)
	return nil
}

func (r *AlamedaScalerReconciler) isControllerHasAlamedaScalerInfo(
	controller *datahub_resources.Controller, alamedaScaler autoscalingv1alpha1.AlamedaScaler) bool {
	if controller.AlamedaControllerSpec == nil || controller.AlamedaControllerSpec.AlamedaScaler == nil {
		return false
	}
	// TODO: Might compare namespace if Datahub return non empty controller.AlamedaControllerSpec.AlamedaScaler.Namespace
	// if controller.AlamedaControllerSpec.AlamedaScaler.Namespace == alamedaScaler.Namespace && controller.AlamedaControllerSpec.AlamedaScaler.Name == alamedaScaler.Name {
	// 	return true
	// }
	if controller.AlamedaControllerSpec.AlamedaScaler.Name == alamedaScaler.Name {
		return true
	}
	return false
}

func (r *AlamedaScalerReconciler) isControllerExistsInAlamedaScalerStatus(
	controller *datahub_resources.Controller, alamedaScaler autoscalingv1alpha1.AlamedaScaler) bool {

	isInAlamedaScaler := false
	switch controller.Kind {
	case datahub_resources.Kind_DEPLOYMENTCONFIG:
		for _, dc := range alamedaScaler.Status.AlamedaController.DeploymentConfigs {
			if controller.ObjectMeta.Name == dc.Name {
				isInAlamedaScaler = true
				break
			}
		}
	case datahub_resources.Kind_DEPLOYMENT:
		for _, deploy := range alamedaScaler.Status.AlamedaController.Deployments {
			if controller.ObjectMeta.Name == deploy.Name {
				isInAlamedaScaler = true
				break
			}
		}
	case datahub_resources.Kind_STATEFULSET:
		for _, statefulSet := range alamedaScaler.Status.AlamedaController.StatefulSets {
			if controller.ObjectMeta.Name == statefulSet.Name {
				isInAlamedaScaler = true
				break
			}
		}
	}
	return isInAlamedaScaler
}

func (r *AlamedaScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&autoscalingv1alpha1.AlamedaScaler{}).
		Complete(r)
}

func (r *AlamedaScalerReconciler) syncCAInfoWithScalerAndMachineGroup(ctx context.Context,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler,
	mgIns autoscalingv1alpha1.AlamedaMachineGroupScaler) error {
	return SyncCAInfoWithScalerAndMachineGroup(ctx, r.ClusterUID, r.Client,
		r.DatahubClient, alamedaScaler, mgIns)
}
