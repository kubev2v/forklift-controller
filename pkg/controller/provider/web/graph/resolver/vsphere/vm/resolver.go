package vm

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
// List all vms.
func (t *Resolver) List(provider *string, filter *graphmodel.VMFilter) ([]graphmodel.VM, error) {
	var vms []graphmodel.VM
	var listOptions = libmodel.ListOptions{Detail: libmodel.MaxDetail}

	if filter != nil {
		var predicates = libmodel.And()
		if filter.CPUHotAddEnabled != nil {
			predicates.Predicates = append(predicates.Predicates, libmodel.Eq("CpuHotAddEnabled", filter.CPUHotAddEnabled))
		}
		if filter.IPAddress != nil {
			predicates.Predicates = append(predicates.Predicates, libmodel.Eq("IPAddress", filter.IPAddress))
		}
		if filter.PowerState != nil {
			predicates.Predicates = append(predicates.Predicates, libmodel.Eq("PowerState", filter.PowerState))
		}
		if filter.MemoryMb != nil {
			predicates.Predicates = append(predicates.Predicates, libmodel.Eq("MemoryMb", filter.MemoryMb))
		}
		listOptions.Predicate = predicates
	}

	allDBs, err := t.GetDBs(provider)
	if err != nil {
		return nil, nil
	}

	for provider, db := range allDBs[api.VSphere] {
		list := []vspheremodel.VM{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			vm := withVsphere(&m, provider)
			vms = append(vms, vm)
		}
	}

	for provider, db := range allDBs[api.OVirt] {
		list := []ovirtmodel.VM{}
		err := db.List(&list, listOptions)
		if err != nil {
			return nil, nil
		}

		for _, m := range list {
			vm := withOvirt(&m, provider)
			vms = append(vms, vm)
		}
	}
	return vms, nil
}

//
// Get a specific vm.
func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereVM, error) {
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	m := &vspheremodel.VM{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err = db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("VM '%s' not found", id)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	vm := withVsphere(m, provider)
	return vm, nil
}

//
// Get a specific vm.
func (t *Resolver) GetByHost(hostId, provider string) ([]*graphmodel.VsphereVM, error) {
	var vms []*graphmodel.VsphereVM

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []vspheremodel.VM{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("host", hostId)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}
	for _, m := range list {
		vms = append(vms, withVsphere(&m, provider))
	}

	return vms, nil
}

//
// Get all VMs for a specific storage/
func (t *Resolver) GetByDatastore(datastoreId, provider string) ([]*graphmodel.VsphereVM, error) {
	var vms []*graphmodel.VsphereVM

	list, err := t.listVsphere(provider)
	if err != nil {
		return nil, nil
	}

	for _, vm := range list {
		if contains(vm.Disks, datastoreId) {
			vms = append(vms, vm)
		}
	}

	return vms, nil
}

//
// Get all vms for a specific datacenter.
func (t *Resolver) GetByDatacenter(folderID, provider string) ([]graphmodel.VsphereVMGroup, error) {
	var vms []graphmodel.VsphereVMGroup
	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err

	}

	cl := t.GetChildrenIDs(db, folderID, "VM")

	list := []vspheremodel.VM{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("id", cl)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := withVsphere(&m, provider)
		vms = append(vms, c)
	}

	return vms, nil
}

func (t *Resolver) listVsphere(provider string) ([]*graphmodel.VsphereVM, error) {
	var vms []*graphmodel.VsphereVM

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, nil
	}

	list := []vspheremodel.VM{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		vm := withVsphere(&m, provider)
		vms = append(vms, vm)
	}
	return vms, nil
}

func contains(l []*graphmodel.Disk, s string) bool {
	for _, d := range l {
		if d.Datastore == s {
			return true
		}
	}
	return false
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

func withVsphere(m *vspheremodel.VM, provider string) (h *graphmodel.VsphereVM) {
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
		Kind:                  api.VSphere,
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

func withOvirt(m *ovirtmodel.VM, provider string) (h *graphmodel.OvirtVM) {
	return &graphmodel.OvirtVM{
		ID:                          m.ID,
		Kind:                        api.OVirt,
		Provider:                    provider,
		Name:                        m.Name,
		Description:                 m.Description,
		Cluster:                     m.Cluster,
		Host:                        m.Host,
		RevisionValidated:           int(m.RevisionValidated),
		PolicyVersion:               m.PolicyVersion,
		GuestName:                   m.GuestName,
		CPUSockets:                  int(m.CpuSockets),
		CPUCores:                    int(m.CpuCores),
		CPUThreads:                  int(m.CpuThreads),
		CPUShares:                   int(m.CpuShares),
		Memory:                      int(m.Memory),
		BalloonedMemory:             m.BootMenuEnabled,
		Bios:                        m.BIOS,
		Display:                     m.Display,
		IOThreads:                   int(m.IOThreads),
		StorageErrorResumeBehaviour: m.StorageErrorResumeBehaviour,
		HaEnabled:                   m.HaEnabled,
		UsbEnabled:                  m.UsbEnabled,
		BootMenuEnabled:             m.BootMenuEnabled,
		// CPUAffinity:                 m.CpuAffinity,
	}
}
