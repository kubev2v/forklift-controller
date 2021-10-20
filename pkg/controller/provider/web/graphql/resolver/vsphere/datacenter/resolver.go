package datacenter

import (
	"errors"
	"fmt"

	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver"
)

type Resolver struct {
	resolver.Resolver
}

//
// List all datacenters.
func (t *Resolver) List(provider string) ([]*graphmodel.VsphereDatacenter, error) {
	var datacenters []*graphmodel.VsphereDatacenter
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	list := []vspheremodel.Datacenter{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err = (*db).List(&list, listOptions)
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
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	m := &vspheremodel.Datacenter{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err = (*db).Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("datacenter '%s' not found", id)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	dc := with(m)
	dc.Provider = provider

	return dc, nil
}

func with(m *vspheremodel.Datacenter) (h *graphmodel.VsphereDatacenter) {
	return &graphmodel.VsphereDatacenter{
		ID:           m.ID,
		Name:         m.Name,
		ClustersID:   m.Clusters.ID,
		DatastoresID: m.Datastores.ID,
		NetworksID:   m.Networks.ID,
		VmsID:        m.Vms.ID,
	}
}
