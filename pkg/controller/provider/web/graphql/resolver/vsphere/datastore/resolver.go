package datastore

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

func (t *Resolver) List(provider string) ([]*graphmodel.VsphereDatastore, error) {
	var datastores []*graphmodel.VsphereDatastore

	db := *t.GetDB(provider)
	list := []vspheremodel.Datastore{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := With(&m)
		c.Provider = provider
		datastores = append(datastores, c)
	}

	return datastores, nil
}

func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereDatastore, error) {
	db := *t.GetDB(provider)
	m := &vspheremodel.Datastore{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		t.Log.Info("Datastore not found")
		return nil, nil
	}

	c := With(m)
	c.Provider = provider

	return c, nil
}

func With(m *vspheremodel.Datastore) (h *graphmodel.VsphereDatastore) {
	return &graphmodel.VsphereDatastore{
		ID:          m.ID,
		Name:        m.Name,
		Kind:        m.Parent.Kind,
		Capacity:    int(m.Capacity),
		Free:        int(m.Free),
		Maintenance: m.MaintenanceMode,
	}
}
