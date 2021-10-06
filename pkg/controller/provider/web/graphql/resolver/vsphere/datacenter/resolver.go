package datacenter

import (
	"errors"

	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type Resolver struct {
	Container *libcontainer.Container
	Log       *logging.Logger
}

func (t *Resolver) List(provider string) ([]*graphmodel.VsphereDatacenter, error) {
	var datacenters []*graphmodel.VsphereDatacenter
	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = t.Container.Get(p); !found {
		t.Log.Info("Provider not found")
		return nil, nil
	}

	db := collector.DB()
	list := []vspheremodel.Datacenter{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		dc := With(&m)
		dc.Provider = provider
		datacenters = append(datacenters, dc)
	}

	return datacenters, nil
}

func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereDatacenter, error) {
	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = t.Container.Get(p); !found {
		t.Log.Info("Provider not found")
		return nil, nil
	}

	m := &vspheremodel.Datacenter{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	db := collector.DB()
	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		t.Log.Info("Datacenter not found")
		return nil, nil
	}

	dc := With(m)
	dc.Provider = provider

	return dc, nil
}

func With(m *vspheremodel.Datacenter) (h *graphmodel.VsphereDatacenter) {
	return &graphmodel.VsphereDatacenter{
		ID:   m.ID,
		Name: m.Name,
		Kind: m.Parent.Kind,
	}
}
