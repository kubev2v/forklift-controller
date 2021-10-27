// The graph packages provides graphql services for essentially querying the controller inventory

// Example of queries:

// query Providers {
// 	providers {
// 	  id name type
// 	  datacenters {
// 		... on VsphereDatacenter {
// 		  id name
// 		  clusters {
// 			id name
// 			hosts {
// 			  id name cpuCores cpuSockets productName
// 			  datastores { id name capacity maintenance free }
// 			  vms {
// 				id name firmware
// 				networks {
// 				  ... on VsphereNetwork { id variant }
// 				  ... on DvPortGroup { id variant dvSwitch }
// 				  ... on DvSwitch { id variant host { host PNIC } }
// 				}
// 			  }
// 			}
// 		  }
// 		}
// 		... on OvirtDatacenter {id provider name }
// 	  }
// 	}
//   }

//   query ProviderTree($provider: ID!) {
// 	provider(id: $provider) {
// 	  id name type
// 	  datacenters {
// 		... on VsphereDatacenter {
// 		  id name
// 		  clusters {
// 			id name
// 			hosts {
// 			  id name cpuCores cpuSockets productName
// 			  datastores { id name capacity maintenance free }
// 			  vms {
// 				id name firmware
// 				networks {
// 				  ... on VsphereNetwork { id variant }
// 				  ... on DvPortGroup { id variant dvSwitch }
// 				  ... on DvSwitch { id variant host { host PNIC } }
// 				}
// 			  }
// 			}
// 		  }
// 		}
// 		... on OvirtDatacenter {
// 		  id provider name
// 		  storagedomain
// 		}
// 	  }
// 	}
//   }

//   query AllDatacenters {
// 	datacenters {
// 	   ... on VsphereDatacenter {
// 		id
// 		name
// 		provider
// 		clusters {
// 		  id
// 		  name
// 		  hosts {
// 			id name cpuCores cpuSockets productName
// 		  }
// 		  datastores {
// 			id
// 			name
// 		  }
// 		  networks {
// 			... on VsphereNetwork { id variant }
// 			... on DvPortGroup { id variant dvSwitch }
// 			... on DvSwitch { id variant host { host PNIC } }
// 		  }
// 		}
// 		datastores { id name }
// 		networks {
// 		  ... on VsphereNetwork { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		}
// 		vms { ... on VsphereVM { id name firmware } }
// 	  }
// 	  ... on OvirtDatacenter { id provider name }
// 	}
//   }

//   query Datacenters($provider: ID!) {
// 	datacenters(provider: $provider) {
// 	  ... on VsphereDatacenter {
// 		id
// 		name
// 		provider
// 		clusters {
// 		  id
// 		  name
// 		  hosts {
// 			id name cpuCores cpuSockets productName
// 		  }
// 		  datastores {
// 			id
// 			name
// 		  }
// 		  networks {
// 			... on VsphereNetwork { id variant }
// 			... on DvPortGroup { id variant dvSwitch }
// 			... on DvSwitch { id variant host { host PNIC } }
// 		  }
// 		}
// 		datastores { id name }
// 		networks {
// 		  ... on VsphereNetwork { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		}
// 		vms { ... on VsphereVM { id name firmware }}
// 	  }
// 	  ... on OvirtDatacenter { id provider name }
// 	}
//   }

//   query Datacenter($provider: ID!){
// 	datacenter(id: "datacenter-21", provider: $provider) {
// 	  ... on VsphereDatacenter {
// 		id name
// 		clusters { id name }
// 		datastores { id name }
// 		networks {
// 		  ... on VsphereNetwork { id variant name  }
// 		  ... on DvPortGroup { id variant name dvSwitch }
// 		  ... on DvSwitch { id variant name host { host PNIC } }
// 		}
// 		vms { ... on VsphereVM { id name firmware }}
// 	  }
// 	  ... on OvirtDatacenter { id provider name}
// 	}
//   }

//   query AllClusters{
// 	clusters {
// 	  ... on VsphereCluster {
// 		id kind name provider dasEnabled drsEnabled drsBehavior
// 		hosts { id name cpuCores cpuSockets provider }
// 		networks {
// 		  ... on VsphereNetwork { id variant name }
// 		  ... on DvPortGroup { id variant name dvSwitch }
// 		  ... on DvSwitch { id variant name host { host } }
// 		}
// 		datastores { id name }
// 	  }
// 	  ... on OvirtCluster { id kind provider name }
// 	}
//   }

//   query Clusters($provider: ID!){
// 	clusters(provider: $provider) {
// 	  ... on VsphereCluster {
// 		 id kind name dasEnabled drsEnabled drsBehavior
// 		 hosts { id name cpuCores cpuSockets provider }
// 		 networks {
// 		   ... on VsphereNetwork { id variant name }
// 		   ... on DvPortGroup { id variant name dvSwitch }
// 		   ... on DvSwitch { id variant name host { host } }
// 		 }
// 		 datastores { id name }
// 	  }
// 	  ... on OvirtCluster { id kind provider name }
// 	}
//   }

//   query AllHosts {
// 	hosts {
// 	  ... on VsphereHost {
// 		id name provider cpuCores cpuSockets
// 		datastores { id name }
// 		vms { id name firmware }
// 	  }
// 	  ... on OvirtHost { id provider name }
// 	}
//   }

//   query Hosts($provider: ID!){
// 	hosts(provider: $provider) {
// 	  ... on VsphereHost {
// 		id name cpuCores cpuSockets
// 		datastores { id name }
// 		vms { id name firmware }
// 	  }
// 	  ... on OvirtHost  { id provider name }
// 	}
//   }

//   query Host($provider: ID!){
// 	host(id: "host-44", provider: $provider) {
// 	  ... on VsphereHost {
// 		id name cpuCores cpuSockets
// 		datastores { id }
// 		networks {
// 		  ... on VsphereNetwork { id variant name }
// 		  ... on DvPortGroup { id variant name dvSwitch }
// 		  ... on DvSwitch { id variant name host { host } }
// 		}
// 		networking {
// 		  portGroups {
// 			name key
// 		  }
// 		  vSwitches {
// 			name
// 			portGroups
// 			pNICs
// 		  }
// 		  vNICs {
// 			key
// 			ipAddress
// 		  }
// 		  pNICs {
// 			key
// 			linkSpeed
// 		  }
// 		}
// 		vms { id name firmware }
// 	  }
// 	  ... on OvirtHost  { id provider name }
// 	}
//   }

//   query AllStorages {
// 	storages {
// 	  ... on VsphereDatastore {
// 		id name provider
// 		hosts { id name }
// 		vms { id name firmware }
// 	  }
// 	  ... on OvirtStorageDomain {
// 		id name provider
// 	  }
// 	}
//   }

//   query Storages($provider: ID!) {
// 	storages(provider: $provider) {
// 	  ... on VsphereDatastore {
// 		id name
// 		hosts { id name }
// 		vms { id name firmware }
// 	  }
// 	  ... on OvirtStorageDomain {
// 		id name provider
// 	  }
// 	}
//   }

//   query Storage($provider: ID!) {
// 	storage(id: "datastore-45", provider: $provider) {
// 	  ... on VsphereDatastore {
// 		id name
// 		hosts { id name }
// 		vms { id name firmware }
// 	  }
// 	  ... on OvirtStorageDomain {
// 		id name provider
// 	  }
// 	}
//   }

//   query AllNetworks{
// 	networks {
// 	  ... on VsphereNetwork { id provider tag }
// 	  ... on DvPortGroup { id provider dvSwitch }
// 	  ... on DvSwitch { id provider host { host PNIC }}
// 	  ... on OvirtNetwork { id provider name }
// 	}
//   }

//   query Networks($provider: ID!){
// 	networks(provider: $provider) {
// 	  ... on VsphereNetwork { id variant tag }
// 	  ... on DvPortGroup { id variant dvSwitch }
// 	  ... on DvSwitch { id variant host { host PNIC }}
// 	  ... on OvirtNetwork { id provider name }
// 	}
//   }

//   query Network($provider: ID!){
// 	network(id: "dvportgroup-56", provider: $provider) {
// 	  ... on VsphereNetwork { id variant }
// 	  ... on DvPortGroup { id variant dvSwitch }
// 	  ... on DvSwitch { id variant host { host PNIC }}
// 	  ... on OvirtNetwork { id provider name }
// 	}
//   }

//   query AllVMs {
// 	vms {
// 	  ... on VsphereVM {
// 		id kind provider name revision ipAddress cpuHotAddEnabled powerState memoryMB
// 		disks { key datastore }
// 		concerns { label }
// 		networks {
// 		  ... on VsphereNetwork { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		}
// 	  }
// 	  ... on OvirtVM {
// 		id kind provider name description cluster host guestName
// 	  }
// 	}
//   }

//   query VMs($provider: ID!){
// 	vms(provider: $provider, filter: {cpuHotAddEnabled: false, ipAddress: "10.19.2.32", powerState: "poweredOn", memoryMB: 4096}) {
// 	  ... on VsphereVM {
// 		id kind name revision ipAddress cpuHotAddEnabled powerState memoryMB
// 		disks { key datastore }
// 		concerns { label }
// 		networks {
// 		  ... on VsphereNetwork { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		}
// 	  }
// 	  ... on OvirtVM {id kind provider name }
// 	}
//   }

//   query VM($provider: ID!){
// 	vm(id: "vm-3344", provider: $provider) {
// 	  ... on VsphereVM {
// 		id name disks { key datastore } concerns { label }
// 		networks {
// 		  ... on VsphereNetwork { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		}
// 	  }
// 	  ... on OvirtVM {id provider name }
// 	}
//   }

//   query AllFolders {
// 	vspherefolders {
// 	  ...FolderFields
// 	  ...ChildrenRecursive
// 	}
//   }

//   query Folders($provider: ID!) {
// 	vspherefolders(provider: $provider) {
// 	  ...FolderFields
// 	  ...ChildrenRecursive
// 	}
//   }

//   query Folder($provider: ID!) {
// 	vspherefolder(id: "group-s24", provider: $provider) {
// 	  ...FolderFields
// 	  ...ChildrenRecursive
// 	}
//   }

//   fragment FolderFields on VsphereFolder {
// 	id
// 	name
// 	provider
// 	parent
// 	children {
// 	  ... on VsphereDatacenter {id name }
// 	  ... on VsphereCluster {id name }
// 	  ... on VsphereDatastore {id name }
// 	  ... on VsphereNetwork {id name }
// 	  ... on DvPortGroup {id name }
// 	  ... on DvSwitch {id name }
// 	  ... on VsphereVM {id name }
// 	}
//   }

//   fragment ChildrenRecursive on VsphereFolder {
// 	children {
// 	  ...FolderFields
// 	  ... on VsphereFolder {
// 		...FolderFields
// 		children {
// 		  ...FolderFields
// 		  ... on VsphereFolder {
// 			...FolderFields
// 			children {
// 			  ...FolderFields
// 			  ... on VsphereFolder {
// 				...FolderFields
// 			  }
// 			}
// 		  }
// 		}
// 	  }
// 	}
//   }

//   query NonExistingProvider{
// 	hosts(provider: "mystery") {
// 	  ... on VsphereHost {
// 		id name cpuCores cpuSockets
// 	  }
// 	}
//   }

//   query HostFailure($provider: ID!){
// 	host(id: "host-XYZ", provider: $provider) {
// 	  ... on VsphereHost {
// 		id name cpuCores cpuSockets
// 	  }
// 	}
//   }

package graph

import (
	"github.com/gin-gonic/gin"
	"github.com/konveyor/controller/pkg/inventory/container"
	"github.com/konveyor/controller/pkg/logging"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	resolverbase "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/cluster"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/datacenter"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/folder"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/host"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/network"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/provider"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/storage"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/vm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

//
// Routes.
const (
	GraphqlRoot = "graphql"
)

//
// GraphQL handler.
type GraphHandler struct {
	base.Handler
}

//
// Add routes to the `gin` router.
func (h *GraphHandler) AddRoutes(e *gin.Engine) {
	e.POST(GraphqlRoot, h.Post)
	e.GET(GraphqlRoot+"/playground", h.Get)
}

func newBaseResolver(c *container.Container, name string) resolverbase.Resolver {
	return resolverbase.Resolver{
		Container: c,
		Log:       logging.WithName("graphql|" + name),
	}
}

//
// GraphQL Queries handler.
func (h GraphHandler) Post(ctx *gin.Context) {
	config := generated.Config{
		Resolvers: &Resolver{
			Folder: folder.Resolver{
				Resolver: newBaseResolver(h.Container, "folder"),
			},
			Providers: provider.Resolver{
				Resolver: newBaseResolver(h.Container, "provider"),
			},
			Datacenter: datacenter.Resolver{
				Resolver: newBaseResolver(h.Container, "datacenter"),
			},
			Cluster: cluster.Resolver{
				Resolver: newBaseResolver(h.Container, "cluster"),
			},
			Host: host.Resolver{
				Resolver: newBaseResolver(h.Container, "host"),
			},
			Storage: storage.Resolver{
				Resolver: newBaseResolver(h.Container, "storage"),
			},
			Network: network.Resolver{
				Resolver: newBaseResolver(h.Container, "network"),
			},
			VM: vm.Resolver{
				Resolver: newBaseResolver(h.Container, "vm"),
			},
		},
	}

	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	handler.ServeHTTP(ctx.Writer, ctx.Request)
}

//
// GraphQL Playground plugin handler.
func (h GraphHandler) Get(ctx *gin.Context) {
	handler := playground.Handler("GraphQL Playground", "/query")
	handler.ServeHTTP(ctx.Writer, ctx.Request)
}
