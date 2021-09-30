package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	libcontainer "github.com/konveyor/controller/pkg/inventory/container"
	libmodel "github.com/konveyor/controller/pkg/inventory/model"
	api "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1"
	vspheremodel "github.com/konveyor/forklift-controller/pkg/controller/provider/model/vsphere"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (r *queryResolver) VsphereProviders(ctx context.Context) ([]*graphmodel.VsphereProvider, error) {
	var providers []*graphmodel.VsphereProvider

	container := ctx.Value("GraphqlContainer").(*libcontainer.Container)
	if container == nil {
		log.Info("could not retrieve Container")
		return providers, nil
	}

	list := container.List()

	for _, collector := range list {
		fmt.Println(collector)
		provider := graphmodel.VsphereProvider{Name: "Name", Kind: "Kind"}
		providers = append(providers, &provider)
	}

	return providers, nil
}

func (r *queryResolver) VsphereHosts(ctx context.Context, provider string) ([]*graphmodel.VsphereHost, error) {
	var hosts []*graphmodel.VsphereHost

	c := ctx.Value("GraphqlContainer").(*libcontainer.Container)
	if c == nil {
		log.Info("could not retrieve Container")
		return hosts, nil
	}

	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = c.Get(p); !found {
		log.Info("Provider not found")
		return nil, nil
	}

	p = collector.Owner().(*api.Provider)
	// TODO: EnsureParity

	db := collector.DB()
	list := []vspheremodel.Host{}

	listOptions := libmodel.ListOptions{Detail: libmodel.MaxDetail}
	err := db.List(&list, listOptions)
	if err != nil {
		return nil, nil
	}

	for _, m := range list {
		host := &graphmodel.VsphereHost{
			Name:           m.Name,
			Kind:           m.Parent.Kind,
			ProductName:    m.ProductName,
			ProductVersion: m.ProductVersion,
			InMaintenance:  m.InMaintenanceMode,
			CPUSockets:     int(m.CpuSockets),
			CPUCores:       int(m.CpuCores),
		}
		hosts = append(hosts, host)
	}

	return hosts, nil
}

func (r *queryResolver) VsphereHost(ctx context.Context, id string, provider string) (*graphmodel.VsphereHost, error) {
	var host *graphmodel.VsphereHost

	c := ctx.Value("GraphqlContainer").(*libcontainer.Container)
	if c == nil {
		log.Info("could not retrieve Container")
		return host, nil
	}

	p := &api.Provider{
		ObjectMeta: meta.ObjectMeta{
			UID: types.UID(provider),
		},
	}

	var found bool
	var collector libcontainer.Collector
	if collector, found = c.Get(p); !found {
		log.Info("Provider not found")
		return nil, nil
	}

	p = collector.Owner().(*api.Provider)
	// TODO: EnsureParity

	m := &vspheremodel.Host{
		Base: vspheremodel.Base{
			ID: id,
		},
	}

	db := collector.DB()
	err := db.Get(m)
	if errors.Is(err, vspheremodel.NotFound) {
		log.Info("Host not found")
		return nil, nil
	}

	host = &graphmodel.VsphereHost{
		Name:           m.Name,
		Kind:           m.Parent.Kind,
		ProductName:    m.ProductName,
		ProductVersion: m.ProductVersion,
		InMaintenance:  m.InMaintenanceMode,
		CPUSockets:     int(m.CpuSockets),
		CPUCores:       int(m.CpuCores),
	}

	return host, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
