package ocp

import (
	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libocp "github.com/konveyor/controller/pkg/inventory/container/ocp"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	core "k8s.io/api/core/v1"
	"path"
)

//
// New collector.
func New(db libmodel.DB, provider *api.Provider, secret *core.Secret) libcontainer.Collector {
	return &Collector{
		Collector: libocp.New(
			db,
			provider,
			secret,
			&Namespace{
				log: logging.WithName("collection|namespace").WithValues(
					"provider",
					path.Join(
						provider.GetNamespace(),
						provider.GetName())),
			},
			&NetworkAttachmentDefinition{
				log: logging.WithName("collection|network").WithValues(
					"provider",
					path.Join(
						provider.GetNamespace(),
						provider.GetName())),
			},
			&StorageClass{
				log: logging.WithName("collection|storageclass").WithValues(
					"provider",
					path.Join(
						provider.GetNamespace(),
						provider.GetName())),
			},
			&VM{
				log: logging.WithName("collection|vm").WithValues(
					"provider",
					path.Join(
						provider.GetNamespace(),
						provider.GetName())),
			}),
	}
}

//
// OCP collector.
type Collector struct {
	*libocp.Collector
}
