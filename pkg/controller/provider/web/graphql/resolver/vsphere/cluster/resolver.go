package cluster

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

func (t *Resolver) List(provider string) ([]*graphmodel.VsphereCluster, error) {
	var clusters []*graphmodel.VsphereCluster

	db := *t.GetDB(provider)
	list := []vspheremodel.Cluster{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := With(&m)
		c.Provider = provider
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereCluster, error) {
	db := *t.GetDB(provider)
	m := &vspheremodel.Cluster{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		t.Log.Info("Cluster not found")
		return nil, nil
	}

	c := With(m)
	c.Provider = provider

	return c, nil
}

func (t *Resolver) GetByDatacenter(datacenterId, provider string) ([]*graphmodel.VsphereCluster, error) {
	var clusters []*graphmodel.VsphereCluster
	db := *t.GetDB(provider)
	folderList := []vspheremodel.Folder{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.And(libmodel.Eq("datacenter", datacenterId), libmodel.Eq("name", "host"))}
	err := db.List(&folderList, listOptions)
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
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := With(&m)
		c.Provider = provider
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func With(m *vspheremodel.Cluster) (h *graphmodel.VsphereCluster) {
	return &graphmodel.VsphereCluster{
		ID:   m.ID,
		Name: m.Name,
		Kind: m.Parent.Kind,
		// DasVms:      m.DasVms,
		DrsEnabled:  m.DrsEnabled,
		DrsBehavior: m.DrsBehavior,
		// DrsVms:      m.DrsVms,
	}
}
