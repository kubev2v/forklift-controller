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

func (r *queryResolver) VsphereDatastore(ctx context.Context, id string, provider string) (*graphmodel.VsphereDatastore, error) {
	return r.Resolver.Datastore.Get(id, provider)
}

func (r *queryResolver) VsphereDatastores(ctx context.Context, provider string) ([]*graphmodel.VsphereDatastore, error) {
	return r.Resolver.Datastore.List(provider)
}

func (r *queryResolver) VsphereNetwork(ctx context.Context, id string, provider string) (graphmodel.NetworkGroup, error) {
	return r.Resolver.Network.Get(id, provider)
}

func (r *queryResolver) VsphereNetworks(ctx context.Context, provider string) ([]graphmodel.NetworkGroup, error) {
	return r.Resolver.Network.List(provider)
}

func (r *queryResolver) VsphereVMs(ctx context.Context, provider string) ([]*graphmodel.VsphereVM, error) {
	return r.Resolver.VM.List(provider)
}

func (r *queryResolver) VsphereVM(ctx context.Context, id string, provider string) (*graphmodel.VsphereVM, error) {
	return r.Resolver.VM.Get(id, provider)
}

func (r *vsphereClusterResolver) Hosts(ctx context.Context, obj *graphmodel.VsphereCluster) ([]*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.GetByCluster(obj.ID, obj.Provider)
}

func (r *vsphereDatacenterResolver) Clusters(ctx context.Context, obj *graphmodel.VsphereDatacenter) ([]*graphmodel.VsphereCluster, error) {
	return r.Resolver.Cluster.GetByDatacenter(obj.ID, obj.Provider)
}

func (r *vsphereDatastoreResolver) Hosts(ctx context.Context, obj *graphmodel.VsphereDatastore) ([]*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.GetbyDatastore(obj.ID, obj.Provider)
}

func (r *vsphereDatastoreResolver) Vms(ctx context.Context, obj *graphmodel.VsphereDatastore) ([]*graphmodel.VsphereVM, error) {
	return r.Resolver.VM.GetbyDatastore(obj.ID, obj.Provider)
}

func (r *vsphereHostResolver) Vms(ctx context.Context, obj *graphmodel.VsphereHost) ([]*graphmodel.VsphereVM, error) {
	return r.Resolver.VM.GetByHost(obj.ID, obj.Provider)
}

func (r *vsphereProviderResolver) Datacenters(ctx context.Context, obj *graphmodel.VsphereProvider) ([]*graphmodel.VsphereDatacenter, error) {
	return r.Resolver.Datacenter.List(obj.ID)
}

func (r *vsphereVMResolver) Networks(ctx context.Context, obj *graphmodel.VsphereVM) ([]graphmodel.NetworkGroup, error) {
	return r.Resolver.Network.GetByIDs(obj.NetRefs, obj.Provider)
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

// VsphereDatastore returns generated.VsphereDatastoreResolver implementation.
func (r *Resolver) VsphereDatastore() generated.VsphereDatastoreResolver {
	return &vsphereDatastoreResolver{r}
}

// VsphereHost returns generated.VsphereHostResolver implementation.
func (r *Resolver) VsphereHost() generated.VsphereHostResolver { return &vsphereHostResolver{r} }

// VsphereProvider returns generated.VsphereProviderResolver implementation.
func (r *Resolver) VsphereProvider() generated.VsphereProviderResolver {
	return &vsphereProviderResolver{r}
}

// VsphereVM returns generated.VsphereVMResolver implementation.
func (r *Resolver) VsphereVM() generated.VsphereVMResolver { return &vsphereVMResolver{r} }

type queryResolver struct{ *Resolver }
type vsphereClusterResolver struct{ *Resolver }
type vsphereDatacenterResolver struct{ *Resolver }
type vsphereDatastoreResolver struct{ *Resolver }
type vsphereHostResolver struct{ *Resolver }
type vsphereProviderResolver struct{ *Resolver }
type vsphereVMResolver struct{ *Resolver }
