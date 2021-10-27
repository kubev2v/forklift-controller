package web

import (
	"github.com/konveyor/controller/pkg/inventory/container"
	libweb "github.com/konveyor/controller/pkg/inventory/web"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/ocp"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/ovirt"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/vsphere"
	"github.com/konveyor/forklift-controller/pkg/settings"
)

//
// All handlers.
func All(container *container.Container) (all []libweb.RequestHandler) {
	all = []libweb.RequestHandler{
		&libweb.SchemaHandler{},
		&ProviderHandler{
			Handler: base.Handler{
				Container: container,
			},
		},
	}
	all = append(
		all,
		ocp.Handlers(container)...)
	all = append(
		all,
		vsphere.Handlers(container)...)
	all = append(
		all,
		ovirt.Handlers(container)...)
	if settings.Settings.Inventory.GrahpQLEnabled {
		all = append(
			all,
			graph.Handlers(container)...)
	}
	return
}
