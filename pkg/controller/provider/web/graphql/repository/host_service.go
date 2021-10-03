package repository

import (
	"errors"
	"fmt"

	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	"github.com/konveyor/controller/pkg/logging"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type HostImpl struct {
	Container *libcontainer.Container
	Log       *logging.Logger
}

func (t *HostImpl) Get(id string, provider string) (*graphmodel.VsphereHost, error) {
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

	myhost := &graphmodel.VsphereHost{
		ID:             m.ID,
		Name:           m.Name,
		Kind:           m.Parent.Kind,
		ProductName:    m.ProductName,
		ProductVersion: m.ProductVersion,
		InMaintenance:  m.InMaintenanceMode,
		CPUSockets:     int(m.CpuSockets),
		CPUCores:       int(m.CpuCores),
	}

	fmt.Println(myhost)
	// return &host, nil
	return myhost, nil

	// hostItem := host.HostItem{Id: "01234", Name: "blah"}
	// return &hostItem, nil
}
