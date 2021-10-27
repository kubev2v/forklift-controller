package host

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
// List all hosts.
func (t *Resolver) List(provider *string) ([]graphmodel.Host, error) {
	var hosts []graphmodel.Host

	allDBs, err := t.GetDBs(provider)
	if err != nil {
		return nil, nil
	}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	for provider, db := range allDBs[api.VSphere] {
		list := []vspheremodel.Host{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			h := withVsphere(&m, provider)
			hosts = append(hosts, h)
		}
	}

	for provider, db := range allDBs[api.OVirt] {
		list := []ovirtmodel.Host{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			h := withOvirt(&m, provider)
			hosts = append(hosts, h)
		}
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

	h := withVsphere(m, provider)

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
		h := withVsphere(&m, provider)
		hosts = append(hosts, h)
	}

	return hosts, nil
}

//
// Get all host for a specific cluster.
func (t *Resolver) GetByOvirtCluster(clusterId, provider string) ([]*graphmodel.OvirtHost, error) {
	var hosts []*graphmodel.OvirtHost

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []ovirtmodel.Host{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("cluster", clusterId)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		h := withOvirt(&m, provider)
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
func (t *Resolver) GetByDatastore(datastoreId, providerId string) ([]*graphmodel.VsphereHost, error) {
	var hosts []*graphmodel.VsphereHost

	list, err := t.listVsphere(providerId)
	if err != nil {
		return nil, nil
	}

	for _, vh := range list {
		if contains(vh.DatastoreIDs, datastoreId) {
			hosts = append(hosts, vh)
		}
	}
	return hosts, nil
}

func (t *Resolver) listVsphere(provider string) ([]*graphmodel.VsphereHost, error) {
	var hosts []*graphmodel.VsphereHost

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, nil
	}

	list := []vspheremodel.Host{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		h := withVsphere(&m, provider)
		hosts = append(hosts, h)
	}

	return hosts, nil
}

func withVsphere(m *vspheremodel.Host, provider string) (h *graphmodel.VsphereHost) {
	var datastores []string
	for _, d := range m.Datastores {
		datastores = append(datastores, d.ID)
	}

	var networksIDs []string
	for _, n := range m.Networks {
		networksIDs = append(networksIDs, n.ID)
	}

	var pnics []*graphmodel.Pnic
	for _, p := range m.Network.PNICs {
		pnics = append(pnics, &graphmodel.Pnic{Key: p.Key, LinkSpeed: int(p.LinkSpeed)})
	}

	var vnics []*graphmodel.Vnic
	for _, v := range m.Network.VNICs {
		vnics = append(vnics, &graphmodel.Vnic{
			Key:        v.Key,
			PortGroup:  v.PortGroup,
			DPortGroup: v.DPortGroup,
			IPAddress:  v.IpAddress,
			// SubnetMask ?
			Mtu: int(v.MTU),
		})
	}

	var portgroups []*graphmodel.PortGroup
	for _, p := range m.Network.PortGroups {
		portgroups = append(portgroups, &graphmodel.PortGroup{
			Key:     p.Key,
			Name:    p.Name,
			Vswitch: p.Switch,
		})
	}

	var vswitches []*graphmodel.VSwitch
	for _, v := range m.Network.Switches {
		vswitches = append(vswitches, &graphmodel.VSwitch{
			Key:        v.Key,
			Name:       v.Name,
			PortGroups: v.PortGroups,
			PNICs:      v.PNICs,
		})
	}

	// TODO: needed for backward compatibility
	// meanwhile can graphql replace this encapsulation ?
	var networkAdapters []*graphmodel.NetworkAdapter

	return &graphmodel.VsphereHost{
		ID:             m.ID,
		Name:           m.Name,
		Kind:           "VsphereHost",
		Provider:       provider,
		Cluster:        m.Cluster,
		ProductName:    m.ProductName,
		ProductVersion: m.ProductVersion,
		InMaintenance:  m.InMaintenanceMode,
		CPUSockets:     int(m.CpuSockets),
		CPUCores:       int(m.CpuCores),
		DatastoreIDs:   datastores,
		NetworksIDs:    networksIDs,
		Networking: &graphmodel.ConfigNetwork{
			PNICs:      pnics,
			VNICs:      vnics,
			PortGroups: portgroups,
			VSwitches:  vswitches,
		},
		NetworkAdapters: networkAdapters,
	}
}

func withOvirt(m *ovirtmodel.Host, provider string) (h *graphmodel.OvirtHost) {
	return &graphmodel.OvirtHost{
		ID:             m.ID,
		Name:           m.Name,
		Kind:           "OvirtHost",
		Provider:       provider,
		Status:         m.Status,
		ProductName:    m.ProductName,
		ProductVersion: m.ProductVersion,
		InMaintenance:  m.InMaintenance,
		CPUSockets:     int(m.CpuSockets),
		CPUCores:       int(m.CpuCores),
		// networkAttachments: [NetworkAttachment!]!,
		// nics: [HostNIC!]!,
	}
}
