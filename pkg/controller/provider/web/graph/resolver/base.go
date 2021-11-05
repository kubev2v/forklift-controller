package base

import (
	"errors"
	"fmt"

	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	model "github.com/konveyor/forklift-controller/pkg/controller/provider/model/ocp"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	webvsphere "github.com/konveyor/forklift-controller/pkg/controller/provider/web/vsphere"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type Resolver struct {
	Container *libcontainer.Container
	Log       *logging.Logger
}

func (t *Resolver) GetDB(provider string) (libmodel.DB, error) {
	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = t.Container.Get(p); !found {
		msg := fmt.Sprintf("provider '%s' not found", provider)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	db := collector.DB()
	return db, nil
}

func (t *Resolver) GetChildrenIDs(db libmodel.DB, folderId, kind string) (list []string) {
	folder := &vspheremodel.Folder{
		Base: vspheremodel.Base{
			ID: folderId,
		},
	}

	err := db.Get(folder)
	if err != nil {
		return nil
	}

	for _, c := range folder.Children {
		if c.Kind == kind {
			list = append(list, c.ID)
		}
		if c.Kind == "Folder" {
			list = append(list, t.GetChildrenIDs(db, c.ID, kind)...)
		}
	}

	return list
}

//
// Get all providers DBs.
func (t *Resolver) GetDBs(provider *string) (map[string]map[string]libmodel.DB, error) {
	vsphere := map[string]libmodel.DB{}
	ovirt := map[string]libmodel.DB{}
	var allProviders = map[string]map[string]libmodel.DB{
		api.VSphere: vsphere,
		api.OVirt:   ovirt,
	}
	list := t.Container.List()

	for _, collector := range list {
		if p, cast := collector.Owner().(*api.Provider); cast {
			m := &model.Provider{}
			m.With(p)
			r := webvsphere.Provider{}
			r.With(m)

			db, err := t.GetDB(m.UID)
			if err != nil {
				return nil, err
			}

			switch p.Spec.Type {
			case api.VSphere:
				vsphere = make(map[string]libmodel.DB)
				vsphere[m.UID] = db
				allProviders[api.VSphere] = vsphere

			case api.OVirt:
				ovirt = make(map[string]libmodel.DB)
				ovirt[m.UID] = db
				allProviders[api.OVirt] = ovirt
			}
		}
	}

	if provider == nil {
		return allProviders, nil
	}

	providers := t.FindProvider(*provider, allProviders)
	return providers, nil
}

func (t *Resolver) FindProvider(provider string, providers map[string]map[string]libmodel.DB) map[string]map[string]libmodel.DB {
	for providerType, v := range providers {
		for uid := range v {
			if uid == provider {
				found := map[string]map[string]libmodel.DB{
					providerType: v,
				}
				return found
			}
		}
	}
	return nil
}

func (t *Resolver) WithVsphereCluster(m *vspheremodel.Cluster, provider string) (h *graphmodel.VsphereCluster) {
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
		Kind:          api.VSphere + "Cluster",
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

func (t *Resolver) WithVsphereStorage(m *vspheremodel.Datastore, provider string) (h *graphmodel.VsphereDatastore) {
	return &graphmodel.VsphereDatastore{
		ID:          m.ID,
		Kind:        api.VSphere + "Datastore",
		Name:        m.Name,
		Provider:    provider,
		Capacity:    int(m.Capacity),
		Free:        int(m.Free),
		Maintenance: m.MaintenanceMode,
	}
}

func withDisk(m *vspheremodel.Disk) (h *graphmodel.Disk) {
	return &graphmodel.Disk{
		Key:       int(m.Key),
		Datastore: m.Datastore.ID,
		File:      m.File,
		Capacity:  int(m.Capacity),
		Shared:    m.Shared,
		Rdm:       m.RDM,
	}
}

func withConcern(m *vspheremodel.Concern) (c *graphmodel.Concern) {
	return &graphmodel.Concern{
		Label:      m.Label,
		Category:   m.Category,
		Assessment: m.Assessment,
	}
}

func (t *Resolver) WithVsphereNetwork(m *vspheremodel.Network, provider string) (h *graphmodel.VsphereNetwork) {
	return &graphmodel.VsphereNetwork{
		ID:       m.ID,
		Provider: provider,
		Kind:     api.VSphere + "Network",
		Variant:  m.Variant,
		Name:     m.Name,
		Tag:      m.Tag,
	}
}

func (t *Resolver) WithDvPortGroup(m *vspheremodel.Network, provider string) (h *graphmodel.DvPortGroup) {
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

func (t *Resolver) WithDvSwitch(m *vspheremodel.Network, provider string) (h *graphmodel.DvSwitch) {
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

func (t *Resolver) WithVsphereVM(m *vspheremodel.VM, provider string) (h *graphmodel.VsphereVM) {
	var cpuAffinity []int
	for _, c := range m.CpuAffinity {
		cpuAffinity = append(cpuAffinity, int(c))
	}

	var disks []*graphmodel.Disk
	for _, d := range m.Disks {
		disks = append(disks, withDisk(&d))
	}

	var concerns []*graphmodel.Concern
	for _, c := range m.Concerns {
		concerns = append(concerns, withConcern(&c))
	}

	var networks []string
	for _, n := range m.Networks {
		networks = append(networks, n.ID)
	}

	var devices []*graphmodel.Device
	for _, n := range m.Devices {
		d := graphmodel.Device{
			Kind: n.Kind,
		}
		devices = append(devices, &d)
	}

	return &graphmodel.VsphereVM{
		ID:                    m.ID,
		Kind:                  api.VSphere + "VM",
		Provider:              provider,
		Name:                  m.Name,
		Revision:              int(m.Revision),
		RevisionValidated:     int(m.RevisionValidated),
		IPAddress:             m.IpAddress,
		UUID:                  m.UUID,
		Firmware:              m.Firmware,
		PowerState:            m.PowerState,
		CPUHotAddEnabled:      m.CpuHotAddEnabled,
		CPUHotRemoveEnabled:   m.CpuHotRemoveEnabled,
		MemoryHotAddEnabled:   m.MemoryHotAddEnabled,
		FaultToleranceEnabled: m.FaultToleranceEnabled,
		CPUCount:              int(m.CpuCount),
		CoresPerSocket:        int(m.CoresPerSocket),
		MemoryMb:              int(m.MemoryMB),
		GuestName:             m.GuestName,
		BalloonedMemory:       int(m.BalloonedMemory),
		StorageUsed:           int(m.StorageUsed),
		Concerns:              concerns,
		Disks:                 disks,
		NumaNodeAffinity:      m.NumaNodeAffinity,
		CPUAffinity:           cpuAffinity,
		Devices:               devices,
		HostID:                m.Host,
		NetIDs:                networks,
	}
}
