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
			c := withVsphere(&m, provider)
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

	c := withVsphere(m, provider)
	return c, nil
}

//
// Get all clusters for a specific vsphere datacenter.
func (t *Resolver) GetByVsphereDatacenter(folderID, provider string) ([]*graphmodel.VsphereCluster, error) {
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
		c := withVsphere(&m, provider)
		clusters = append(clusters, c)
	}

	return clusters, nil
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

func withVsphere(m *vspheremodel.Cluster, provider string) (h *graphmodel.VsphereCluster) {
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
		Kind:          "VsphereCluster",
		Provider:      provider,
		DasEnabled:    m.DasEnabled,
		DasVmsIDs:     dasVmList,
		DrsEnabled:    m.DrsEnabled,
		DrsBehavior:   m.DrsBehavior,
		DrsVmsIDs:     drsVmList,
		DatastoresIDs: datastoresIDs,
		NetworksIDs:   networksIDs,
	}
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
