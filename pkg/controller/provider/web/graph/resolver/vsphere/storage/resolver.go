package storage

import (
	"errors"
	"fmt"

	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	ovirtmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/ovirt"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	base "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver"
)

type Resolver struct {
	base.Resolver
}

//
// List all storage entities.
func (t *Resolver) List(provider *string) ([]graphmodel.Storage, error) {
	var datastores []graphmodel.Storage

	allDBs, err := t.GetDBs(provider)
	if err != nil {
		return nil, nil
	}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	for provider, db := range allDBs[api.VSphere] {
		list := []vspheremodel.Datastore{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			d := withVsphere(&m, provider)
			datastores = append(datastores, d)
		}
	}

	for provider, db := range allDBs[api.OVirt] {
		list := []ovirtmodel.StorageDomain{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			d := withOvirt(&m, provider)
			datastores = append(datastores, d)
		}
	}

	return datastores, nil
}

//
// Get datastores for specific IDs.
func (t *Resolver) GetByIds(ids []string, provider string) ([]*graphmodel.VsphereDatastore, error) {
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
		c := withVsphere(&m, provider)
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

	c := withVsphere(m, provider)

	return c, nil
}

//
// Get all datastores for a specific vsphere datacenter.
func (t *Resolver) GetByVsphereDatacenter(folderID, provider string) ([]*graphmodel.VsphereDatastore, error) {
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
		c := withVsphere(&m, provider)
		datastores = append(datastores, c)
	}
	return datastores, nil
}

//
// Get all storagedomains for a specific ovirt datacenter.
func (t *Resolver) GetByOvirtDatacenter(datacenterId, provider string) ([]*graphmodel.OvirtStorageDomain, error) {
	var clusters []*graphmodel.OvirtStorageDomain

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []ovirtmodel.StorageDomain{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("datacenter", datacenterId)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := withOvirt(&m, provider)
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func withVsphere(m *vspheremodel.Datastore, provider string) (h *graphmodel.VsphereDatastore) {
	return &graphmodel.VsphereDatastore{
		ID:          m.ID,
		Kind:        "VsphereDatastore",
		Name:        m.Name,
		Provider:    provider,
		Capacity:    int(m.Capacity),
		Free:        int(m.Free),
		Maintenance: m.MaintenanceMode,
	}
}

func withOvirt(m *ovirtmodel.StorageDomain, provider string) (h *graphmodel.OvirtStorageDomain) {
	return &graphmodel.OvirtStorageDomain{
		ID:          m.ID,
		Kind:        "OvirtStorageDomain",
		Name:        m.Name,
		Provider:    provider,
		DataCenter:  m.DataCenter,
		Type:        m.Type,
		StorageType: m.Storage.Type,
		Available:   int(m.Available),
		Used:        int(m.Used),
	}
}
