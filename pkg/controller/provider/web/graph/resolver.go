package graph

import (
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/vsphere/cluster"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/vsphere/host"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graphql/vsphere/provider"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Provider provider.Provider
	Cluster  cluster.Cluster
	Host     host.Host
}
