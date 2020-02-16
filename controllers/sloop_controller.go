/*
Copyright 2020 The Kubernetes Authors.

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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	shipsv1beta1 "k8s.io/hello/api/v1beta1"
)

// +kubebuilder:rbac:groups=ships.k8s.io,resources=events,verbs=create;patch
// SloopReconciler reconciles a Sloop object
type SloopReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=ships.k8s.io,resources=sloops,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ships.k8s.io,resources=sloops/status,verbs=get;update;patch

func (r *SloopReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	logger := r.Log.WithValues("sloop", req.NamespacedName)
	logger.Info("Hello world")
	ins := &shipsv1beta1.Sloop{}
	r.Get(context.Background(), req.NamespacedName, ins)

	r.Recorder.Event(ins, "Normal", "NormalType", "works fine")

	// your logic here

	return ctrl.Result{}, nil
}

func (r *SloopReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&shipsv1beta1.Sloop{}).
		Complete(r)
}
