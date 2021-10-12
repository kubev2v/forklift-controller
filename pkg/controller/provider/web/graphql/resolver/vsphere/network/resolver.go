package network

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

//
// List all Network objects.
func (t *Resolver) List(provider string) ([]graphmodel.NetworkGroup, error) {
	var networks []graphmodel.NetworkGroup

	db := *t.GetDB(provider)
	networkList := []vspheremodel.Network{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&networkList, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range networkList {
		switch m.Variant {
		case vspheremodel.NetStandard:
			networks = append(networks, withNetwork(&m))
		case vspheremodel.NetDvPortGroup:
			networks = append(networks, withDvPortGroup(&m))
		case vspheremodel.NetDvSwitch:
			networks = append(networks, withDvSwitch(&m))
		}
	}

	return networks, nil
}

//
// Get a specific Network object.
func (t *Resolver) Get(id string, provider string) (graphmodel.NetworkGroup, error) {
	db := *t.GetDB(provider)

	m := &vspheremodel.Network{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		t.Log.Info("Network not found")
		return nil, nil
	}

	var network graphmodel.NetworkGroup
	switch m.Variant {
	case vspheremodel.NetStandard:
		network = withNetwork(m)
	case vspheremodel.NetDvPortGroup:
		network = withDvPortGroup(m)
	case vspheremodel.NetDvSwitch:
		network = withDvSwitch(m)
	}

	return network, nil
}

func withNetwork(m *vspheremodel.Network) (h *graphmodel.Network) {
	return &graphmodel.Network{
		ID:   m.ID,
		Kind: m.Variant,
		Name: m.Name,
		Tag:  m.Tag,
	}
}

func withDvPortGroup(m *vspheremodel.Network) (h *graphmodel.DvPortGroup) {
	return &graphmodel.DvPortGroup{
		ID:       m.ID,
		Kind:     m.Variant,
		Name:     m.Name,
		DvSwitch: m.DVSwitch.ID,
		// Host:  m.Host,
		// Ports: m.Ports,
	}
}

func withDvSwitch(m *vspheremodel.Network) (h *graphmodel.DvSwitch) {
	var host []*graphmodel.DvSHost

	for _, h := range m.Host {
		nh := graphmodel.DvSHost{
			Host: h.Host.ID,
			Pnic: h.PNIC,
		}
		host = append(host, &nh)
	}

	return &graphmodel.DvSwitch{
		ID:   m.ID,
		Kind: m.Variant,
		Name: m.Name,
		Host: host,
	}
}
