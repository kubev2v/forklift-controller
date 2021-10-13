package vm

import (
	"errors"
	"fmt"

	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/base"
)

type Resolver struct {
	base.Resolver
}

//
// List all vms.
func (t *Resolver) List(provider string) ([]*graphmodel.VsphereVM, error) {
	var vms []*graphmodel.VsphereVM

	db := *t.GetDB(provider)
	list := []vspheremodel.VM{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		vm := with(&m)
		vm.Provider = provider
		vms = append(vms, vm)
	}

	return vms, nil
}

//
// Get a specific vm.
func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereVM, error) {
	db := *t.GetDB(provider)

	m := &vspheremodel.VM{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		msg := fmt.Sprintf("VM '%s' not found", id)
		t.Log.Info(msg)
		return nil, errors.New(msg)
	}

	vm := with(m)
	vm.Provider = provider

	return vm, nil
}

func (t *Resolver) GetByHost(hostId, provider string) ([]*graphmodel.VsphereVM, error) {
	var vms []*graphmodel.VsphereVM

	db := *t.GetDB(provider)
	list := []vspheremodel.VM{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("host", hostId)}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}
	for _, m := range list {
		vms = append(vms, with(&m))
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

func (t *Resolver) GetbyDatastore(datastoreId, provider string) ([]*graphmodel.VsphereVM, error) {
	list, err := t.List(provider)
	if err != nil {
		return nil, nil
	}

	var vms []*graphmodel.VsphereVM
	for _, vm := range list {
		if contains(vm.Disks, datastoreId) {
			vms = append(vms, vm)
		}
	}

	return vms, nil
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

func with(m *vspheremodel.VM) (h *graphmodel.VsphereVM) {
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

	return &graphmodel.VsphereVM{
		ID:                    m.ID,
		Name:                  m.Name,
		Revision:              int(m.Revision),
		RevisionValidated:     int(m.RevisionValidated),
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
		IPAddress:             m.IpAddress,
		StorageUsed:           int(m.StorageUsed),
		Concerns:              concerns,
		Disks:                 disks,
		NumaNodeAffinity:      m.NumaNodeAffinity,
		CPUAffinity:           cpuAffinity,
		HostID:                m.Host,
		NetIDs:                networks,
		// Devices:               m.Devices,
	}
}
