package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
)

func (r *queryResolver) VsphereProviders(ctx context.Context) ([]*graphmodel.VsphereProvider, error) {
	return r.Resolver.Provider.List()
}

func (r *queryResolver) VsphereProvider(ctx context.Context, id string) (*graphmodel.VsphereProvider, error) {
	return r.Resolver.Provider.Get(id)
}

func (r *queryResolver) VsphereDatacenters(ctx context.Context, provider string) ([]*graphmodel.VsphereDatacenter, error) {
	return r.Resolver.Datacenter.List(provider)
}

func (r *queryResolver) VsphereDatacenter(ctx context.Context, id string, provider string) (*graphmodel.VsphereDatacenter, error) {
	return r.Resolver.Datacenter.Get(id, provider)
}

func (r *queryResolver) VsphereClusters(ctx context.Context, provider string) ([]*graphmodel.VsphereCluster, error) {
	return r.Resolver.Cluster.List(provider)
}

func (r *queryResolver) VsphereCluster(ctx context.Context, id string, provider string) (*graphmodel.VsphereCluster, error) {
	return r.Resolver.Cluster.Get(id, provider)
}

func (r *queryResolver) VsphereHosts(ctx context.Context, provider string) ([]*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.List(provider)
}

func (r *queryResolver) VsphereHost(ctx context.Context, id string, provider string) (*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.Get(id, provider)
}

func (r *vsphereClusterResolver) Hosts(ctx context.Context, obj *graphmodel.VsphereCluster) ([]*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.GetByCluster(obj.ID, obj.Provider)
}

func (r *vsphereDatacenterResolver) Clusters(ctx context.Context, obj *graphmodel.VsphereDatacenter) ([]*graphmodel.VsphereCluster, error) {
	return r.Resolver.Cluster.GetByDatacenter(obj.ID, obj.Provider)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// VsphereCluster returns generated.VsphereClusterResolver implementation.
func (r *Resolver) VsphereCluster() generated.VsphereClusterResolver {
	return &vsphereClusterResolver{r}
}

// VsphereDatacenter returns generated.VsphereDatacenterResolver implementation.
func (r *Resolver) VsphereDatacenter() generated.VsphereDatacenterResolver {
	return &vsphereDatacenterResolver{r}
}

type queryResolver struct{ *Resolver }
type vsphereClusterResolver struct{ *Resolver }
type vsphereDatacenterResolver struct{ *Resolver }
