package vsphere

import (
	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	model "github.com/konveyor/forklift-controller/pkg/controller/provider/model/ocp"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
)

type ProviderRepository struct {
	Container *libcontainer.Container
	Log       *logging.Logger
}

func (t *ProviderRepository) Get(id string) (*graphmodel.VsphereProvider, error) {
	return nil, nil
}

func (t *ProviderRepository) List() ([]*graphmodel.VsphereProvider, error) {
	var providers []*graphmodel.VsphereProvider

	list := t.Container.List()

	for _, collector := range list {
		if p, cast := collector.Owner().(*api.Provider); cast {
			if p.Type() != api.VSphere {
				continue
			}
			m := &model.Provider{}
			m.With(p)

			provider := &graphmodel.VsphereProvider{ID: m.UID, Name: m.Name, Kind: m.Type}
			providers = append(providers, provider)
		}
	}

	return providers, nil
}
