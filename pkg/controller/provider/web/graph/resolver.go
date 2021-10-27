package graph

import (
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/cluster"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/datacenter"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/folder"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/host"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/network"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/provider"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/storage"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/vm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Folder     folder.Resolver
	Providers  provider.Resolver
	Datacenter datacenter.Resolver
	Cluster    cluster.Resolver
	Network    network.Resolver
	Host       host.Resolver
	Storage    storage.Resolver
	VM         vm.Resolver
}
