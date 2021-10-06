package host

import (
	"errors"

	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type Resolver struct {
	Container *libcontainer.Container
	Log       *logging.Logger
}

func (t *Resolver) List(provider string) ([]*graphmodel.VsphereHost, error) {
	var hosts []*graphmodel.VsphereHost
	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = t.Container.Get(p); !found {
		t.Log.Info("Provider not found")
		return nil, nil
	}

	db := collector.DB()
	list := []vspheremodel.Host{}

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

func (t *Resolver) Get(id string, provider string) (*graphmodel.VsphereHost, error) {
	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = t.Container.Get(p); !found {
		t.Log.Info("Provider not found")
		return nil, nil
	}

	m := &vspheremodel.Host{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	db := collector.DB()
	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		t.Log.Info("Host not found")
		return nil, nil
	}

	h := With(m)

	return h, nil
}

func (t *Resolver) GetByCluster(clusterId, provider string) ([]*graphmodel.VsphereHost, error) {
	var hosts []*graphmodel.VsphereHost
	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = t.Container.Get(p); !found {
		t.Log.Info("Provider not found")
		return nil, nil
	}

	db := collector.DB()
	list := []vspheremodel.Host{}
	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail, Predicate: libmodel.Eq("cluster", clusterId)}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		hosts = append(hosts, With(&m))
	}

	return hosts, nil
}

func With(m *vspheremodel.Host) (h *graphmodel.VsphereHost) {
	return &graphmodel.VsphereHost{
		ID:             m.ID,
		Name:           m.Name,
		Kind:           m.Parent.Kind,
		ProductName:    m.ProductName,
		ProductVersion: m.ProductVersion,
		InMaintenance:  m.InMaintenanceMode,
		CPUSockets:     int(m.CpuSockets),
		CPUCores:       int(m.CpuCores),
	}
}
