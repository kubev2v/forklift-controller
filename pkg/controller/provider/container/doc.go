package container

import (
	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/container/ocp"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/container/ovirt"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/container/vsphere"
	core "k8s.io/api/core/v1"
)

//
// Build
func Build(
	db libmodel.DB,
	provider *api.Provider,
	secret *core.Secret) libcontainer.Collector {
	//
	switch provider.Type() {
	case api.OpenShift:
		return ocp.New(db, provider, secret)
	case api.VSphere:
		return vsphere.New(db, provider, secret)
	case api.OVirt:
		return ovirt.New(db, provider, secret)
	}

	return nil
}
