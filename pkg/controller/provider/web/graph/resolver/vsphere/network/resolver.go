package network

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
// List all Network objects.
func (t *Resolver) List(provider *string) ([]graphmodel.NetworkGroup, error) {
	var networks []graphmodel.NetworkGroup

	providers := t.ListDBs(provider)
	for provider, db := range providers {
		list := []vspheremodel.Network{}
		listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			switch m.Variant {
			case vspheremodel.NetStandard:
				networks = append(networks, withNetwork(&m, provider))
			case vspheremodel.NetDvPortGroup:
				networks = append(networks, withDvPortGroup(&m, provider))
			case vspheremodel.NetDvSwitch:
				networks = append(networks, withDvSwitch(&m, provider))
			}
		}
	}

	return networks, nil
}

//
// Get a specific Network object.
func (t *Resolver) Get(id string, provider string) (graphmodel.NetworkGroup, error) {
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

	var network graphmodel.NetworkGroup
	switch m.Variant {
	case vspheremodel.NetStandard:
		network = withNetwork(m, provider)
	case vspheremodel.NetDvPortGroup:
		network = withDvPortGroup(m, provider)
	case vspheremodel.NetDvSwitch:
		network = withDvSwitch(m, provider)
	}

	return network, nil
}

//
// Get Network from specific IDs
func (t *Resolver) GetByIDs(ids []string, provider string) ([]graphmodel.NetworkGroup, error) {
	var networks []graphmodel.NetworkGroup

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
			networks = append(networks, withNetwork(&m, provider))
		case vspheremodel.NetDvPortGroup:
			networks = append(networks, withDvPortGroup(&m, provider))
		case vspheremodel.NetDvSwitch:
			networks = append(networks, withDvSwitch(&m, provider))
		}
	}

	return networks, nil
}

//
// Get all networks for a specific datacenter.
func (t *Resolver) GetByDatacenter(folderID, provider string) ([]graphmodel.NetworkGroup, error) {
	var networks []graphmodel.NetworkGroup
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	nl := t.GetChildrenIDs(db, folderID, "Network")

	list := []vspheremodel.Network{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("id", nl)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		switch m.Variant {
		case vspheremodel.NetStandard:
			networks = append(networks, withNetwork(&m, provider))
		case vspheremodel.NetDvPortGroup:
			networks = append(networks, withDvPortGroup(&m, provider))
		case vspheremodel.NetDvSwitch:
			networks = append(networks, withDvSwitch(&m, provider))
		}
	}
	return networks, nil
}

func withNetwork(m *vspheremodel.Network, provider string) (h *graphmodel.Network) {
	return &graphmodel.Network{
		ID:       m.ID,
		Variant:  m.Variant,
		Name:     m.Name,
		Provider: provider,
		Tag:      m.Tag,
	}
}

func withDvPortGroup(m *vspheremodel.Network, provider string) (h *graphmodel.DvPortGroup) {
	return &graphmodel.DvPortGroup{
		ID:       m.ID,
		Variant:  m.Variant,
		Name:     m.Name,
		Provider: provider,
		DvSwitch: m.DVSwitch.ID,
		// Host:  m.Host,
		// Ports: m.Ports,
	}
}

func withDvSwitch(m *vspheremodel.Network, provider string) (h *graphmodel.DvSwitch) {
	var host []*graphmodel.DvSHost

	for _, h := range m.Host {
		nh := graphmodel.DvSHost{
			Host: h.Host.ID,
			Pnic: h.PNIC,
		}
		host = append(host, &nh)
	}

	return &graphmodel.DvSwitch{
		ID:       m.ID,
		Variant:  m.Variant,
		Name:     m.Name,
		Provider: provider,
		Host:     host,
	}
}
