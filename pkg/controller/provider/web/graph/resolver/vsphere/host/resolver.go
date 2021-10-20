package host

import (
	"errors"
	"fmt"

	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver"
)

type Resolver struct {
	resolver.Resolver
}

//
// List all hosts.
func (t *Resolver) List(provider string) ([]*graphmodel.VsphereHost, error) {
	var hosts []*graphmodel.VsphereHost

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	list := []vspheremodel.Host{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}
	for _, m := range list {
		h := with(&m)
		h.Provider = provider
		hosts = append(hosts, h)
	}

	return hosts, nil
}

//
// Get a specific host.
func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereHost, error) {
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	m := &vspheremodel.Host{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err = db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("host '%s' not found", id)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	h := with(m)
	h.Provider = provider

	return h, nil
}

//
// Get all host for a specific cluster.
func (t *Resolver) GetByCluster(clusterId, provider string) ([]*graphmodel.VsphereHost, error) {
	var hosts []*graphmodel.VsphereHost

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []vspheremodel.Host{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("cluster", clusterId)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		h := with(&m)
		h.Provider = provider
		hosts = append(hosts, h)
	}

	return hosts, nil
}

func contains(l []string, s string) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}
	return false
}

//
// Get all hosts for a specific datastore.
func (t *Resolver) GetbyDatastore(datastoreId, provider string) ([]*graphmodel.VsphereHost, error) {
	list, err := t.List(provider)
	if err != nil {
		return nil, nil
	}

	var hosts []*graphmodel.VsphereHost
	for _, h := range list {
		if contains(h.DatastoreIDs, datastoreId) {
			hosts = append(hosts, h)
		}
	}

	return hosts, nil
}

func with(m *vspheremodel.Host) (h *graphmodel.VsphereHost) {
	var datastores []string
	for _, d := range m.Datastores {
		datastores = append(datastores, d.ID)
	}
	return &graphmodel.VsphereHost{
		ID:             m.ID,
		Name:           m.Name,
		ProductName:    m.ProductName,
		ProductVersion: m.ProductVersion,
		InMaintenance:  m.InMaintenanceMode,
		CPUSockets:     int(m.CpuSockets),
		CPUCores:       int(m.CpuCores),
		DatastoreIDs:   datastores,
	}
}
