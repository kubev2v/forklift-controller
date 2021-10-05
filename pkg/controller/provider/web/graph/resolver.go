package graph

import (
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/cluster"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/datacenter"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/host"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/resolver/vsphere/provider"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Provider   provider.Repository
	Datacenter datacenter.Repository
	Cluster    cluster.Repository
	Host       host.Repository
}
