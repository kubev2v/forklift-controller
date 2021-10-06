package provider

import (
	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	model "github.com/konveyor/forklift-controller/pkg/controller/provider/model/ocp"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/vsphere"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type Resolver struct {
	Container *libcontainer.Container
	Log       *logging.Logger
}

func (t *Resolver) Get(id string) (*graphmodel.VsphereProvider, error) {
	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(id),
		},
	}

	// h.Detail = true
	m := &model.Provider{}
	m.With(p)
	r := vsphere.Provider{}
	r.With(m)

	provider := &graphmodel.VsphereProvider{ID: r.UID, Name: r.Name, Kind: r.Type}
	return provider, nil
}

func (t *Resolver) List() ([]*graphmodel.VsphereProvider, error) {
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
