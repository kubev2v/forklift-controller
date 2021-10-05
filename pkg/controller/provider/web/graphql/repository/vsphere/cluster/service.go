package cluster

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

type Repository struct {
	Container *libcontainer.Container
	Log       *logging.Logger
}

func (t *Repository) List(provider string) ([]*graphmodel.VsphereCluster, error) {
	var clusters []*graphmodel.VsphereCluster
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
	list := []vspheremodel.Cluster{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		c := With(&m)
		c.Provider = provider
		clusters = append(clusters, c)
	}

	return clusters, nil
}

func (t *Repository) Get(id string, provider string) (*graphmodel.VsphereCluster, error) {
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

	m := &vspheremodel.Cluster{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	db := collector.DB()
	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		t.Log.Info("Cluster not found")
		return nil, nil
	}

	c := With(m)
	c.Provider = provider

	return c, nil
}

func With(m *vspheremodel.Cluster) (h *graphmodel.VsphereCluster) {
	return &graphmodel.VsphereCluster{
		ID:   m.ID,
		Name: m.Name,
		Kind: m.Parent.Kind,
		// DasVms:      m.DasVms,
		DrsEnabled:  m.DrsEnabled,
		DrsBehavior: m.DrsBehavior,
		// DrsVms:      m.DrsVms,
	}
}
