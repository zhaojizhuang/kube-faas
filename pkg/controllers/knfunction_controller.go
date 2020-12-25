/*
Copyright 2020.

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
	appv1 "chumper.cn/kube-faas/api/v1"
	"context"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// KnFunctionReconciler reconciles a KnFunction object
type KnFunctionReconciler struct {
	client.Client
	Log     logr.Logger
	Scheme  *runtime.Scheme
	Context context.Context
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=app.chumper.cn,resources=knfunctions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=app.chumper.cn,resources=knfunctions/status,verbs=get;update;patch

func (r *KnFunctionReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("knfunction", req.NamespacedName)

	// your logic here
	var FunList appv1.KnFunctionList
	r.List(r.Context, &FunList, client.InNamespace(req.Namespace))
	var fun appv1.KnFunction
	r.Get(r.Context,&fun)
	log.V(1).Info("list", "functionlist", FunList)
	r.Recorder.Eventf(req, corev1.EventTypeWarning, "Error", "some error")
	return ctrl.Result{}, nil
}

func (r *KnFunctionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.KnFunction{}).
		//Watches(&source.Kind{Type: Deployment{}}, &handler.EnqueueRequestForObject{}).
		Complete(r)
}
