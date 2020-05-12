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
	"time"
	machinesetrepository "github.com/containers-ai/alameda/operator/datahub/client/machineset"
	machinegrouprepository "github.com/containers-ai/alameda/operator/datahub/client/machinegroup"
	"github.com/containers-ai/alameda/operator/pkg/machinegroup"
	datahubschemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/go-logr/logr"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
)

// AlamedaMachineGroupScalerReconciler reconciles a AlamedaMachineGroupScaler object
type AlamedaMachineGroupScalerReconciler struct {
	client.Client
	ClusterUID                       string
	Log                              logr.Logger
	Scheme                           *runtime.Scheme
	DatahubMachineSetRepo   *machinesetrepository.MachineSetRepository
	DatahubMachineGroupRepo          machinegrouprepository.MachineGroupRepository
	DatahubCAMachineGroupSchema      datahubschemas.Schema
	DatahubCAMachineGroupMeasurement datahubschemas.Measurement
	ReconcileTimeout                 time.Duration
}

// +kubebuilder:rbac:groups=autoscaling.containers.ai,resources=alamedamachinegroupscalers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=autoscaling.containers.ai,resources=alamedamachinegroupscalers/status,verbs=get;update;patch

func (r *AlamedaMachineGroupScalerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ReconcileTimeout)
	defer cancel()

	mgs := []machinegroup.MachineGroup{
		{
			ClusterName: r.ClusterUID,
			ResourceMeta: machinegroup.ResourceMeta{
				KubernetesMeta: machinegroup.KubernetesMeta{
					Namespace: req.Namespace,
					Name:      req.Name,
				},
			},
		},
	}

	instance := autoscalingv1alpha1.AlamedaMachineGroupScaler{}
	err := r.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.Name}, &instance)
	if err != nil && k8sErrors.IsNotFound(err) {
		err = r.DatahubMachineGroupRepo.DeleteMachineGroups(ctx, mgs)
		if err != nil {
			scope.Errorf("Delete AlamedaMachineGroupScaler (%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
		return ctrl.Result{Requeue: false}, nil
	} else if err != nil {
		scope.Errorf("Get AlamedaMachineGroupScaler(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}

	alamedaScalerList := &autoscalingv1alpha1.AlamedaScalerList{}
	if err = r.List(ctx, alamedaScalerList); err != nil {
		scope.Errorf("List AlamedaScaler failed: %s", err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
	}
	for _, alamedaScaler := range alamedaScalerList.Items {
		if alamedaScaler.GetNamespace() == req.Namespace &&
			alamedaScaler.GetType() == autoscalingv1alpha1.AlamedaScalerTypeDefault &&
			alamedaScaler.Spec.ScalingTool.Type == autoscalingv1alpha1.ScalingToolTypeCA &&
			alamedaScaler.Spec.ScalingTool.MachineGroupScaler == req.Name {
			scope.Infof("AlamedaScaler (%s/%s) is matched for MachineGroupScaler (%s/%s)",
				alamedaScaler.GetNamespace(), alamedaScaler.GetName(),
				req.Namespace, req.Name)
			if err = r.syncCAInfoWithScalerAndMachineGroup(ctx, alamedaScaler, instance); err != nil {
				scope.Errorf("sync CA info with alamedascaler (%s/%s) and machinegroupscaler (%s/%s) failed",
					alamedaScaler.GetNamespace(), alamedaScaler.GetName(), instance.GetNamespace(), instance.GetName())
				return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
			}
			break
		}
	}

	return ctrl.Result{}, nil
}

func (r *AlamedaMachineGroupScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&autoscalingv1alpha1.AlamedaMachineGroupScaler{}).
		Complete(r)
}

func (r *AlamedaMachineGroupScalerReconciler) syncCAInfoWithScalerAndMachineGroup(ctx context.Context,
	alamedaScaler autoscalingv1alpha1.AlamedaScaler,
	mgIns autoscalingv1alpha1.AlamedaMachineGroupScaler) error {
	return SyncCAInfoWithScalerAndMachineGroup(ctx, r.ClusterUID, r.Client,
		r.DatahubMachineSetRepo, &r.DatahubMachineGroupRepo, alamedaScaler, mgIns)
}
