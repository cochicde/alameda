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

package autoscaling

import (
	"context"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"

	mahcinev1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MachineSetReconciler reconciles a MachineSet object
type MachineSetReconciler struct {
	client.Client
	Scheme           *runtime.Scheme
	ClusterUID       string
	ReconcileTimeout time.Duration
	DatahubClient    *datahubpkg.Client
}

func (r *MachineSetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	instance := mahcinev1beta1.MachineSet{}
	err := r.Get(context.Background(), req.NamespacedName, &instance)
	if err != nil && k8serrors.IsNotFound(err) {
		err = r.DatahubClient.DeleteByOpts(&entities.ResourceClusterAutoscalerMachineset{}, datahubpkg.Option{
			Entity: entities.ResourceClusterAutoscalerMachineset{
				ClusterName: r.ClusterUID,
				Namespace:   req.Namespace,
				Name:        req.Name,
			},
			Fields: []string{"ClusterName", "Namespace", "Name"},
		})
		if err != nil {
			scope.Errorf("Delete MachineSet (%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
			return ctrl.Result{Requeue: true, RequeueAfter: 1 * time.Second}, nil
		}
		return ctrl.Result{}, nil
	} else if err != nil {
		scope.Warnf("Get MachineSet(%s/%s) failed: %s", req.Namespace, req.Name, err.Error())
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	return ctrl.Result{}, nil
}

func (r *MachineSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mahcinev1beta1.MachineSet{}).
		Complete(r)
}
