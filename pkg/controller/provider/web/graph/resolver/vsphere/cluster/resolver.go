package cluster

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
// List all clusters.
func (t *Resolver) List(provider *string) ([]graphmodel.Cluster, error) {
	var clusters []graphmodel.Cluster
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}

	allDBs, err := t.GetDBs(provider)
	if err != nil {
		return nil, nil
	}

	for provider, db := range allDBs[api.VSphere] {
		list := []vspheremodel.Cluster{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			c := t.WithVsphereCluster(&m, provider)
			clusters = append(clusters, c)
		}
	}

	for provider, db := range allDBs[api.OVirt] {
		list := []ovirtmodel.Cluster{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			c := withOvirt(&m, provider)
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

	c := t.WithVsphereCluster(m, provider)
	return c, nil
}

//
// Get all clusters for a specific ovirt datacenter.
func (t *Resolver) GetByOvirtDatacenter(datacenterId, provider string) ([]*graphmodel.OvirtCluster, error) {
	var clusters []*graphmodel.OvirtCluster

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []ovirtmodel.Cluster{}
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

func withOvirt(m *ovirtmodel.Cluster, provider string) (h *graphmodel.OvirtCluster) {
	return &graphmodel.OvirtCluster{
		ID:            m.ID,
		Name:          m.Name,
		Provider:      provider,
		Kind:          "OvirtCluster",
		DataCenter:    m.DataCenter,
		HaReservation: m.HaReservation,
		KsmEnabled:    m.KsmEnabled,
		BiosType:      m.BiosType,
	}
}
