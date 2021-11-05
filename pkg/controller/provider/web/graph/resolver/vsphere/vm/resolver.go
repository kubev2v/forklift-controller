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
			vm := t.WithVsphereVM(&m, provider)
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

	vm := t.WithVsphereVM(m, provider)
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
		vms = append(vms, t.WithVsphereVM(&m, provider))
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
// Get all VMs for a specific ovirt cluster.
func (t *Resolver) GetByOvirtCluster(clusterId, provider string) ([]*graphmodel.OvirtVM, error) {
	var vms []*graphmodel.OvirtVM

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []ovirtmodel.VM{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("cluster", clusterId)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		h := withOvirt(&m, provider)
		vms = append(vms, h)
	}

	return vms, nil
}

//
// Get all VMs for a specific ovirt host.
func (t *Resolver) GetByOvirtHost(clusterId, provider string) ([]*graphmodel.OvirtVM, error) {
	var vms []*graphmodel.OvirtVM

	db, err := t.GetDB(provider)
	if err != nil {
		return nil, err
	}

	list := []ovirtmodel.VM{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("host", clusterId)}
	err = db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		h := withOvirt(&m, provider)
		vms = append(vms, h)
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
		vm := t.WithVsphereVM(&m, provider)
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

func withOvirt(m *ovirtmodel.VM, provider string) (h *graphmodel.OvirtVM) {
	return &graphmodel.OvirtVM{
		ID:                          m.ID,
		Kind:                        api.OVirt + "VM",
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
