package network

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
// List all Network objects.
func (t *Resolver) List(provider *string) ([]graphmodel.Network, error) {
	var networks []graphmodel.Network

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	providers, _ := t.GetDBs(provider)
	for provider, db := range providers[api.VSphere] {
		list := []vspheremodel.Network{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			switch m.Variant {
			case vspheremodel.NetStandard:
				networks = append(networks, t.WithVsphereNetwork(&m, provider))
			case vspheremodel.NetDvPortGroup:
				networks = append(networks, t.WithDvPortGroup(&m, provider))
			case vspheremodel.NetDvSwitch:
				networks = append(networks, t.WithDvSwitch(&m, provider))
			}
		}
	}

	for provider, db := range providers[api.OVirt] {
		list := []ovirtmodel.Network{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			networks = append(networks, withOvirtNetwork(&m, provider))
		}
	}

	return networks, nil
}

//
// Get a specific Network object.
func (t *Resolver) Get(id string, provider string) (graphmodel.Network, error) {
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	m := &vspheremodel.Network{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err = db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("network '%s' not found", id)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	var network graphmodel.Network
	switch m.Variant {
	case vspheremodel.NetStandard:
		network = t.WithVsphereNetwork(m, provider)
	case vspheremodel.NetDvPortGroup:
		network = t.WithDvPortGroup(m, provider)
	case vspheremodel.NetDvSwitch:
		network = t.WithDvSwitch(m, provider)
	}

	return network, nil
}

//
// Get Network from specific IDs
func (t *Resolver) GetByIDs(ids []string, provider string) ([]graphmodel.VsphereNetworkGroup, error) {
	var networks []graphmodel.VsphereNetworkGroup

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	networkList := []vspheremodel.Network{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("id", ids)}
	err = db.List(&networkList, listOptions)
	if err != nil {
		return nil, nil
	}
	for _, m := range networkList {
		switch m.Variant {
		case vspheremodel.NetStandard:
			networks = append(networks, t.WithVsphereNetwork(&m, provider))
		case vspheremodel.NetDvPortGroup:
			networks = append(networks, t.WithDvPortGroup(&m, provider))
		case vspheremodel.NetDvSwitch:
			networks = append(networks, t.WithDvSwitch(&m, provider))
		}
	}

	return networks, nil
}

// Get all networks for a specific ovirt datacenter.
func (t *Resolver) GetByOvirtDatacenter(datacenterId, provider string) ([]*graphmodel.OvirtNetwork, error) {
	var clusters []*graphmodel.OvirtNetwork

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []ovirtmodel.Network{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("datacenter", datacenterId)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := withOvirtNetwork(&m, provider)
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func withOvirtNetwork(m *ovirtmodel.Network, provider string) (h *graphmodel.OvirtNetwork) {
	return &graphmodel.OvirtNetwork{
		ID:          m.ID,
		Provider:    provider,
		Kind:        "OvirtNetwork",
		Name:        m.Name,
		Description: m.Description,
		DataCenter:  m.DataCenter,
		Vlan:        m.VLan,
		Usages:      m.Usages,
		Profiles:    m.Profiles,
	}
}
