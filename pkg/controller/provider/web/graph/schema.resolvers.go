package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
)

func (r *ovirtClusterResolver) Hosts(ctx context.Context, obj *graphmodel.OvirtCluster) ([]*graphmodel.OvirtHost, error) {
	return r.Resolver.Host.GetByOvirtCluster(obj.ID, obj.Provider)
}

func (r *ovirtClusterResolver) Vms(ctx context.Context, obj *graphmodel.OvirtCluster) ([]*graphmodel.OvirtVM, error) {
	return r.Resolver.VM.GetByOvirtCluster(obj.ID, obj.Provider)
}

func (r *ovirtDatacenterResolver) Clusters(ctx context.Context, obj *graphmodel.OvirtDatacenter) ([]*graphmodel.OvirtCluster, error) {
	return r.Resolver.Cluster.GetByOvirtDatacenter(obj.ID, obj.Provider)
}

func (r *ovirtDatacenterResolver) Storages(ctx context.Context, obj *graphmodel.OvirtDatacenter) ([]*graphmodel.OvirtStorageDomain, error) {
	return r.Resolver.Storage.GetByOvirtDatacenter(obj.ID, obj.Provider)
}

func (r *ovirtDatacenterResolver) Networks(ctx context.Context, obj *graphmodel.OvirtDatacenter) ([]*graphmodel.OvirtNetwork, error) {
	return r.Resolver.Network.GetByOvirtDatacenter(obj.ID, obj.Provider)
}

func (r *ovirtHostResolver) Vms(ctx context.Context, obj *graphmodel.OvirtHost) ([]*graphmodel.OvirtVM, error) {
	return r.Resolver.VM.GetByOvirtHost(obj.ID, obj.Provider)
}

func (r *providerResolver) Datacenters(ctx context.Context, obj *graphmodel.Provider) ([]graphmodel.Datacenter, error) {
	return r.Resolver.Datacenter.List(&obj.ID)
}

func (r *queryResolver) Providers(ctx context.Context) ([]*graphmodel.Provider, error) {
	return r.Resolver.Providers.List()
}

func (r *queryResolver) Provider(ctx context.Context, id string) (*graphmodel.Provider, error) {
	return r.Resolver.Providers.Get(id)
}

func (r *queryResolver) Datacenters(ctx context.Context, provider *string) ([]graphmodel.Datacenter, error) {
	return r.Resolver.Datacenter.List(provider)
}

func (r *queryResolver) Datacenter(ctx context.Context, id string, provider string) (graphmodel.Datacenter, error) {
	return r.Resolver.Datacenter.Get(id, provider)
}

func (r *queryResolver) Clusters(ctx context.Context, provider *string) ([]graphmodel.Cluster, error) {
	return r.Resolver.Cluster.List(provider)
}

func (r *queryResolver) Cluster(ctx context.Context, id string, provider string) (graphmodel.Cluster, error) {
	return r.Resolver.Cluster.Get(id, provider)
}

func (r *queryResolver) Hosts(ctx context.Context, provider *string) ([]graphmodel.Host, error) {
	return r.Resolver.Host.List(provider)
}

func (r *queryResolver) Host(ctx context.Context, id string, provider string) (graphmodel.Host, error) {
	return r.Resolver.Host.Get(id, provider)
}

func (r *queryResolver) Storages(ctx context.Context, provider *string) ([]graphmodel.Storage, error) {
	return r.Resolver.Storage.List(provider)
}

func (r *queryResolver) Storage(ctx context.Context, id string, provider string) (graphmodel.Storage, error) {
	return r.Resolver.Storage.Get(id, provider)
}

func (r *queryResolver) Networks(ctx context.Context, provider *string) ([]graphmodel.Network, error) {
	return r.Resolver.Network.List(provider)
}

func (r *queryResolver) Network(ctx context.Context, id string, provider string) (graphmodel.Network, error) {
	return r.Resolver.Network.Get(id, provider)
}

func (r *queryResolver) Vms(ctx context.Context, provider *string, filter *graphmodel.VMFilter) ([]graphmodel.VM, error) {
	return r.Resolver.VM.List(provider, filter)
}

func (r *queryResolver) VM(ctx context.Context, id string, provider string) (graphmodel.VM, error) {
	return r.Resolver.VM.Get(id, provider)
}

func (r *queryResolver) Vspherefolders(ctx context.Context, provider *string) ([]*graphmodel.VsphereFolder, error) {
	return r.Resolver.Folder.List(provider)
}

func (r *queryResolver) Vspherefolder(ctx context.Context, id string, provider string) (*graphmodel.VsphereFolder, error) {
	return r.Resolver.Folder.Get(id, provider)
}

func (r *vsphereClusterResolver) Datastores(ctx context.Context, obj *graphmodel.VsphereCluster) ([]*graphmodel.VsphereDatastore, error) {
	return r.Resolver.Storage.GetByIds(obj.DatastoresIDs, obj.Provider)
}

func (r *vsphereClusterResolver) Networks(ctx context.Context, obj *graphmodel.VsphereCluster) ([]graphmodel.VsphereNetworkGroup, error) {
	return r.Resolver.Network.GetByIDs(obj.NetworksIDs, obj.Provider)
}

func (r *vsphereClusterResolver) Hosts(ctx context.Context, obj *graphmodel.VsphereCluster) ([]*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.GetByCluster(obj.ID, obj.Provider)
}

func (r *vsphereClusterResolver) DasVms(ctx context.Context, obj *graphmodel.VsphereCluster) ([]*graphmodel.VsphereVM, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *vsphereClusterResolver) DrsVms(ctx context.Context, obj *graphmodel.VsphereCluster) ([]*graphmodel.VsphereVM, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *vsphereDatacenterResolver) Clusters(ctx context.Context, obj *graphmodel.VsphereDatacenter) (*graphmodel.VsphereFolder, error) {
	return r.Resolver.Folder.Get(obj.ClustersID, obj.Provider)
}

func (r *vsphereDatacenterResolver) Datastores(ctx context.Context, obj *graphmodel.VsphereDatacenter) (*graphmodel.VsphereFolder, error) {
	return r.Resolver.Folder.Get(obj.DatastoresID, obj.Provider)
}

func (r *vsphereDatacenterResolver) Networks(ctx context.Context, obj *graphmodel.VsphereDatacenter) (*graphmodel.VsphereFolder, error) {
	return r.Resolver.Folder.Get(obj.NetworksID, obj.Provider)
}

func (r *vsphereDatacenterResolver) Vms(ctx context.Context, obj *graphmodel.VsphereDatacenter) (*graphmodel.VsphereFolder, error) {
	return r.Resolver.Folder.Get(obj.VmsID, obj.Provider)
}

func (r *vsphereDatastoreResolver) Hosts(ctx context.Context, obj *graphmodel.VsphereDatastore) ([]*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.GetByDatastore(obj.ID, obj.Provider)
}

func (r *vsphereDatastoreResolver) Vms(ctx context.Context, obj *graphmodel.VsphereDatastore) ([]*graphmodel.VsphereVM, error) {
	return r.Resolver.VM.GetByDatastore(obj.ID, obj.Provider)
}

func (r *vsphereFolderResolver) Children(ctx context.Context, obj *graphmodel.VsphereFolder) ([]graphmodel.VsphereFolderGroup, error) {
	return r.Resolver.Folder.GetChildren(obj.ID, obj.Provider), nil
}

func (r *vsphereHostResolver) Vms(ctx context.Context, obj *graphmodel.VsphereHost) ([]*graphmodel.VsphereVM, error) {
	return r.Resolver.VM.GetByHost(obj.ID, obj.Provider)
}

func (r *vsphereHostResolver) Datastores(ctx context.Context, obj *graphmodel.VsphereHost) ([]*graphmodel.VsphereDatastore, error) {
	return r.Resolver.Storage.GetByIds(obj.DatastoreIDs, obj.Provider)
}

func (r *vsphereHostResolver) Networks(ctx context.Context, obj *graphmodel.VsphereHost) ([]graphmodel.VsphereNetworkGroup, error) {
	return r.Resolver.Network.GetByIDs(obj.NetworksIDs, obj.Provider)
}

func (r *vsphereVMResolver) Host(ctx context.Context, obj *graphmodel.VsphereVM) (*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.Get(obj.HostID, obj.Provider)
}

func (r *vsphereVMResolver) Networks(ctx context.Context, obj *graphmodel.VsphereVM) ([]graphmodel.VsphereNetworkGroup, error) {
	return r.Resolver.Network.GetByIDs(obj.NetIDs, obj.Provider)
}

// OvirtCluster returns generated.OvirtClusterResolver implementation.
func (r *Resolver) OvirtCluster() generated.OvirtClusterResolver { return &ovirtClusterResolver{r} }

// OvirtDatacenter returns generated.OvirtDatacenterResolver implementation.
func (r *Resolver) OvirtDatacenter() generated.OvirtDatacenterResolver {
	return &ovirtDatacenterResolver{r}
}

// OvirtHost returns generated.OvirtHostResolver implementation.
func (r *Resolver) OvirtHost() generated.OvirtHostResolver { return &ovirtHostResolver{r} }

// Provider returns generated.ProviderResolver implementation.
func (r *Resolver) Provider() generated.ProviderResolver { return &providerResolver{r} }

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

// VsphereFolder returns generated.VsphereFolderResolver implementation.
func (r *Resolver) VsphereFolder() generated.VsphereFolderResolver { return &vsphereFolderResolver{r} }

// VsphereHost returns generated.VsphereHostResolver implementation.
func (r *Resolver) VsphereHost() generated.VsphereHostResolver { return &vsphereHostResolver{r} }

// VsphereVM returns generated.VsphereVMResolver implementation.
func (r *Resolver) VsphereVM() generated.VsphereVMResolver { return &vsphereVMResolver{r} }

type ovirtClusterResolver struct{ *Resolver }
type ovirtDatacenterResolver struct{ *Resolver }
type ovirtHostResolver struct{ *Resolver }
type providerResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type vsphereClusterResolver struct{ *Resolver }
type vsphereDatacenterResolver struct{ *Resolver }
type vsphereDatastoreResolver struct{ *Resolver }
type vsphereFolderResolver struct{ *Resolver }
type vsphereHostResolver struct{ *Resolver }
type vsphereVMResolver struct{ *Resolver }
