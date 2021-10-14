package cluster

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
// List all clusters.
func (t *Resolver) List(provider string) ([]*graphmodel.VsphereCluster, error) {
	var clusters []*graphmodel.VsphereCluster

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	list := []vspheremodel.Cluster{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err = (*db).List(&list, listOptions)
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

	err = (*db).Get(m)
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
func (t *Resolver) GetByDatacenter(datacenterId, provider string) ([]*graphmodel.VsphereCluster, error) {
	var clusters []*graphmodel.VsphereCluster
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}
	folderList := []vspheremodel.Folder{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.And(libmodel.Eq("datacenter", datacenterId), libmodel.Eq("name", "host"))}
	err = (*db).List(&folderList, listOptions)
	if err != nil {
		return nil, nil
	}

	var cl []string
	for _, f := range folderList {
		for _, c := range f.Children {
			if c.Kind == "Cluster" {
				cl = append(cl, c.ID)
			}
		}
	}

	list := []vspheremodel.Cluster{}
	listOptions = libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("id", cl)}
	err = (*db).List(&list, listOptions)
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
	return &graphmodel.VsphereCluster{
		ID:   m.ID,
		Name: m.Name,
		// DasVms:      m.DasVms,
		DrsEnabled:  m.DrsEnabled,
		DrsBehavior: m.DrsBehavior,
		// DrsVms:      m.DrsVms,
	}
}
