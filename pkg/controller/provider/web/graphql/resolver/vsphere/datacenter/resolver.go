package datacenter

import (
	"errors"

	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/base"
)

type Resolver struct {
	base.Resolver
}

//
// List all datacenters.
func (t *Resolver) List(provider string) ([]*graphmodel.VsphereDatacenter, error) {
	var datacenters []*graphmodel.VsphereDatacenter
	db := *t.GetDB(provider)
	list := []vspheremodel.Datacenter{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		dc := with(&m)
		dc.Provider = provider
		datacenters = append(datacenters, dc)
	}

	return datacenters, nil
}

//
// Get a specific datacenter.
func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereDatacenter, error) {
	db := *t.GetDB(provider)
	m := &vspheremodel.Datacenter{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		t.Log.Info("Datacenter not found")
		return nil, nil
	}

	dc := with(m)
	dc.Provider = provider

	return dc, nil
}

func with(m *vspheremodel.Datacenter) (h *graphmodel.VsphereDatacenter) {
	return &graphmodel.VsphereDatacenter{
		ID:   m.ID,
		Name: m.Name,
	}
}
