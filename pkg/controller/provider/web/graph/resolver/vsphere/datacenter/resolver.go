package datacenter

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
// List all datacenters.
func (t *Resolver) List(provider *string) ([]graphmodel.Datacenter, error) {
	var datacenters []graphmodel.Datacenter

	allDBs, err := t.GetDBs(provider)
	if err != nil {
		return nil, nil
	}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	for provider, db := range allDBs[api.VSphere] {
		list := []vspheremodel.Datacenter{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			dc := withVsphere(&m, provider)
			datacenters = append(datacenters, dc)
		}
	}
	for provider, db := range allDBs[api.OVirt] {
		list := []ovirtmodel.DataCenter{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			dc := withOvirt(&m, provider)
			datacenters = append(datacenters, dc)
		}
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

	err = db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("datacenter '%s' not found", id)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	dc := withVsphere(m, provider)

	return dc, nil
}

func withVsphere(m *vspheremodel.Datacenter, provider string) (h *graphmodel.VsphereDatacenter) {
	return &graphmodel.VsphereDatacenter{
		ID:           m.ID,
		Name:         m.Name,
		Kind:         "VsphereDatacenter",
		Provider:     provider,
		ClustersID:   m.Clusters.ID,
		DatastoresID: m.Datastores.ID,
		NetworksID:   m.Networks.ID,
		VmsID:        m.Vms.ID,
	}
}

func withOvirt(m *ovirtmodel.DataCenter, provider string) (h *graphmodel.OvirtDatacenter) {
	return &graphmodel.OvirtDatacenter{
		ID:       m.ID,
		Kind:     "OvirtDatacenter",
		Name:     m.Name,
		Provider: provider,
	}
}
