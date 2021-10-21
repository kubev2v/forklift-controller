package cluster

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
// List all clusters.
func (t *Resolver) List(provider *string) ([]*graphmodel.VsphereCluster, error) {
	var clusters []*graphmodel.VsphereCluster

	providers := t.ListDBs(provider)
	for provider, db := range providers {
		list := []vspheremodel.Cluster{}
		listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			c := with(&m)
			c.Provider = provider
			clusters = append(clusters, c)
		}
	}

	return clusters, nil
}

//
// Get a specific cluster.
func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereCluster, error) {
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	m := &vspheremodel.Cluster{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err = db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("cluster '%s' not found", id)
		t.Log.Info(msg)
	}

	c := with(m)
	c.Provider = provider

	return c, nil
}

//
// Get all clusters for a specific datacenter.
func (t *Resolver) GetByDatacenter(folderID, provider string) ([]*graphmodel.VsphereCluster, error) {
	var clusters []*graphmodel.VsphereCluster
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	cl := t.GetChildrenIDs(db, folderID, "Cluster")

	list := []vspheremodel.Cluster{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("id", cl)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := with(&m)
		c.Provider = provider
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func with(m *vspheremodel.Cluster) (h *graphmodel.VsphereCluster) {
	var dasVmList []string
	for _, dasVm := range m.DasVms {
		dasVmList = append(dasVmList, dasVm.ID)
	}

	var drsVmList []string
	for _, dasVm := range m.DasVms {
		drsVmList = append(drsVmList, dasVm.ID)
	}

	var datastoresIDs []string
	for _, ds := range m.Datastores {
		datastoresIDs = append(datastoresIDs, ds.ID)
	}

	var networksIDs []string
	for _, n := range m.Networks {
		networksIDs = append(networksIDs, n.ID)
	}

	return &graphmodel.VsphereCluster{
		ID:            m.ID,
		Name:          m.Name,
		DasEnabled:    m.DasEnabled,
		DasVmsIDs:     dasVmList,
		DrsEnabled:    m.DrsEnabled,
		DrsBehavior:   m.DrsBehavior,
		DrsVmsIDs:     drsVmList,
		DatastoresIDs: datastoresIDs,
		NetworksIDs:   networksIDs,
	}
}
