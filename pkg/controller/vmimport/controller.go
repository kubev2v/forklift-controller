/*
Copyright 2021 Red Hat Inc.

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

package vmimport

import (
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1alpha1"
	"github.com/konveyor/forklift-controller/pkg/settings"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	// Name of a Controller.
	Name = "vmimport"
)

//
// Package logger.
var log = logging.WithName(Name)

//
// Settings of an application.
var Settings = &settings.Settings

//
// Add method creates a new VMImport Controller and adds it to the Manager.
func Add(mgr manager.Manager) error {
	reconciler := &Reconciler{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}
	cnt, err := controller.New(
		Name,
		mgr,
		controller.Options{
			Reconciler: reconciler,
		})
	if err != nil {
		log.Trace(err)
		return err
	}
	// Primary CR.
	err = cnt.Watch(
		&source.Kind{
			Type: &api.VMImport{},
		},
		&handler.EnqueueRequestForObject{})
	if err != nil {
		log.Trace(err)
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &Reconciler{}

//
// Reconciler reconciles a Migration object.
type Reconciler struct {
	client.Client
	scheme *runtime.Scheme
}

//
// Reconcile a VMImport CR.
func (r *Reconciler) Reconcile(request reconcile.Request) (result reconcile.Result, err error) {
	noReQ := reconcile.Result{}
	result = noReQ

	// Reset the logger.
	log.Reset()
	log.SetValues("vmimport-reconciler", request.Name)
	log.Info("Reconcile draft")

	defer func() {
		if err != nil {
			log.Trace(err)
			err = nil
		}
	}()

	// Done
	return
}
