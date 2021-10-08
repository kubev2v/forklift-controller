package vm

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

func (t *Resolver) List(provider string) ([]*graphmodel.VsphereVM, error) {
	var hosts []*graphmodel.VsphereVM

	db := *t.GetDB(provider)
	list := []vspheremodel.VM{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		hosts = append(hosts, With(&m))
	}

	return hosts, nil
}

func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereVM, error) {
	db := *t.GetDB(provider)

	m := &vspheremodel.VM{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		t.Log.Info("VM not found")
		return nil, nil
	}

	h := With(m)

	return h, nil
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
		vms = append(vms, With(&m))
	}

	return vms, nil
}

func With(m *vspheremodel.VM) (h *graphmodel.VsphereVM) {
	return &graphmodel.VsphereVM{
		ID:                    m.ID,
		Name:                  m.Name,
		Kind:                  "VM",
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
		// Host: m.Host,
		// Concerns: m.Concerns,
		// Networks: m.networks,
		// Disks: m.Disks,
		NumaNodeAffinity: m.NumaNodeAffinity,
		// Devices: m.Devices,
		// CPUAffinity: m.CpuAffinity,
	}
}
