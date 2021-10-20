package datastore

import (
	"errors"
	"fmt"

	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	base "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver"
)

type Resolver struct {
	base.Resolver
}

//
// List all datastores.
func (t *Resolver) List(provider string) ([]*graphmodel.VsphereDatastore, error) {
	var datastores []*graphmodel.VsphereDatastore

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	list := []vspheremodel.Datastore{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := with(&m)
		c.Provider = provider
		datastores = append(datastores, c)
	}

	return datastores, nil
}

//
// List all datastores for specific IDs.
func (t *Resolver) ListByIds(ids []string, provider string) ([]*graphmodel.VsphereDatastore, error) {
	var datastores []*graphmodel.VsphereDatastore

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	list := []vspheremodel.Datastore{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("id", ids)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := with(&m)
		c.Provider = provider
		datastores = append(datastores, c)
	}

	return datastores, nil
}

//
// Get a specific datastore.
func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereDatastore, error) {
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	m := &vspheremodel.Datastore{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err = db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("datastore '%s' not found", id)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	c := with(m)
	c.Provider = provider

	return c, nil
}

//
// Get all datastores for a specific datacenter.
func (t *Resolver) GetByDatacenter(folderID, provider string) ([]*graphmodel.VsphereDatastore, error) {
	var datastores []*graphmodel.VsphereDatastore
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	dl := t.GetChildrenIDs(db, folderID, "Datastore")

	list := []vspheremodel.Datastore{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("id", dl)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := with(&m)
		c.Provider = provider
		datastores = append(datastores, c)
	}
	return datastores, nil
}

func with(m *vspheremodel.Datastore) (h *graphmodel.VsphereDatastore) {
	return &graphmodel.VsphereDatastore{
		ID:          m.ID,
		Name:        m.Name,
		Capacity:    int(m.Capacity),
		Free:        int(m.Free),
		Maintenance: m.MaintenanceMode,
	}
}