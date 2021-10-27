// The graph packages provides graphql services for essentially querying the controller inventory

// Example of queries:

// query Providers {
// 	vsphereProviders {
// 	  id name type
// 	  datacenters {
// 		id name
// 		clusters {
// 		  id name
// 		  hosts {
// 			id name cpuCores cpuSockets productName
// 			datastores { id name capacity maintenance free }
// 			vms {
// 			  id name firmware
// 			   networks {
// 				 ... on Network { id variant }
// 				 ... on DvPortGroup { id variant dvSwitch }
// 				 ... on DvSwitch { id variant host { host PNIC } }
// 			  }
// 			}
// 		  }
// 		}
// 	  }
// 	}
//}

// query ProviderTree($provider: ID!) {
// 	vsphereProvider(id: $provider) {
// 	  id name type
// 	  datacenters {
// 		id name
// 		clusters {
// 		  id name
// 		  hosts {
// 			id name cpuCores cpuSockets productName
// 			datastores { id name capacity maintenance free }
// 			vms {
// 			  id name firmware
// 			   networks {
// 				 ... on Network { id variant }
// 				 ... on DvPortGroup { id variant dvSwitch }
// 				 ... on DvSwitch { id variant host { host PNIC } }
// 			  }
// 			}
// 		  }
// 		}
// 	  }
// 	}
// }

// query AllDatacenters {
// 	vsphereDatacenters {
// 	  id
// 	  name
// 	  provider
// 	  clusters {
// 		id
// 		name
// 		hosts {
// 		  id name cpuCores cpuSockets productName
// 		}
// 		datastores {
// 		  id
// 		  name
// 		}
// 		networks {
// 		... on Network { id variant }
// 		... on DvPortGroup { id variant dvSwitch }
// 		... on DvSwitch { id variant host { host PNIC } }
// 		}
// 	  }
// 	  datastores { id name }
// 	  networks {
// 		... on Network { id variant }
// 		... on DvPortGroup { id variant dvSwitch }
// 		... on DvSwitch { id variant host { host PNIC } }
// 	  }
// 	  vms { ... on VsphereVM { id name firmware }}
// 	}
// }

// query Datacenters($provider: ID!) {
// 	vsphereDatacenters(provider: $provider) {
// 	  id
// 	  name
// 	  provider
// 	  clusters {
// 		id
// 		name
// 		hosts {
// 		  id name cpuCores cpuSockets productName
// 		}
// 		datastores {
// 		  id
// 		  name
// 		}
// 		networks {
// 		... on Network { id variant }
// 		... on DvPortGroup { id variant dvSwitch }
// 		... on DvSwitch { id variant host { host PNIC } }
// 		}
// 	  }
// 	  datastores { id name }
// 	  networks {
// 		... on Network { id variant }
// 		... on DvPortGroup { id variant dvSwitch }
// 		... on DvSwitch { id variant host { host PNIC } }
// 	  }
// 	  vms { ... on VsphereVM { id name firmware }}
// 	}
// }

// query Datacenter($provider: ID!){
// 	vsphereDatacenter(id: "datacenter-21", provider: $provider) {
// 	  id name
// 	  clusters { id name }
// 	  datastores { id name }
// 	  networks {
// 		... on Network { id variant name  }
// 		... on DvPortGroup { id variant name dvSwitch }
// 		... on DvSwitch { id variant name host { host PNIC } }
// 	  }
// 	  vms { ... on VsphereVM { id name firmware }}
// 	}
// }

// query AllClusters{
// 	vsphereClusters {
// 	  id name provider dasEnabled drsEnabled drsBehavior
// 	  hosts { id name cpuCores cpuSockets provider }
// 	  networks {
// 		... on Network { id variant name }
// 		... on DvPortGroup { id variant name dvSwitch }
// 		... on DvSwitch { id variant name host { host } }
// 	  }
// 	  datastores { id name }
// 	}
// }

//  query Clusters($provider: ID!){
// 	vsphereClusters(provider: $provider) {
// 	  id name dasEnabled drsEnabled drsBehavior
// 	  hosts { id name cpuCores cpuSockets provider }
// 	  networks {
// 		... on Network { id variant name }
// 		... on DvPortGroup { id variant name dvSwitch }
// 		... on DvSwitch { id variant name host { host } }
// 	  }
// 	  datastores { id name }
// 	}
// }

// query AllHosts {
// 	vsphereHosts {
// 	  id name provider cpuCores cpuSockets
// 	  datastores { id name }
// 	  vms { id name firmware }
// 	}
// }

// query Hosts($provider: ID!){
// 	vsphereHosts(provider: $provider) {
// 	  id name cpuCores cpuSockets
// 	  datastores { id name }
// 	  vms { id name firmware }
// 	}
// }

//  query Host($provider: ID!){
// 	vsphereHost(id: "host-44", provider: $provider) {
// 	  id name cpuCores cpuSockets
// 	  datastores { id }
// 	  networks {
// 		... on Network { id variant name }
// 		... on DvPortGroup { id variant name dvSwitch }
// 		... on DvSwitch { id variant name host { host } }
// 	  }
// 	  networking {
// 		portGroups {
// 		  name key
// 		}
// 		vSwitches {
// 		  name
// 		  portGroups
// 		  pNICs
// 		}
// 		vNICs {
// 		  key
// 		  ipAddress
// 		}
// 		pNICs {
// 		  key
// 		  linkSpeed
// 		}
// 	  }
// 	  vms { id name firmware }
// 	}
// }

// query AllDatastores {
// 	vsphereDatastores {
// 	  id name provider
// 	  hosts { id name }
// 	  vms { id name firmware }
// 	}
// }

// query Datastores($provider: ID!) {
// 	vsphereDatastores(provider: $provider) {
// 	  id name hosts { id name } vms { id name firmware }
// 	}
// }

// query Datastore($provider: ID!) {
// 	vsphereDatastore(id: "datastore-45", provider: $provider) {
// 	  id name hosts { id name } vms { id name firmware }
// 	}
// }

// query AllNetworks{
// 	vsphereNetworks {
// 	  ... on Network { id provider tag }
// 	  ... on DvPortGroup { id provider dvSwitch }
// 	  ... on DvSwitch { id provider host { host PNIC }}
// 	}
// }

// query Networks($provider: ID!){
// 	vsphereNetworks(provider: $provider) {
// 	  ... on Network { id variant tag }
// 	  ... on DvPortGroup { id variant dvSwitch }
// 	  ... on DvSwitch { id variant host { host PNIC }}
// 	}
// }

// query Network($provider: ID!){
// 	vsphereNetwork(id: "dvportgroup-56", provider: $provider) {
// 	  ... on Network { id variant }
// 	  ... on DvPortGroup { id variant dvSwitch }
// 	  ... on DvSwitch { id variant host { host PNIC }}
// 	}
// }

// query AllVMs {
// 	vsphereVMs {
// 	  id provider name revision ipAddress cpuHotAddEnabled powerState memoryMB
// 	  disks { key datastore }
// 	  concerns { label }
// 	  networks {
// 		... on Network { id variant }
// 		... on DvPortGroup { id variant dvSwitch }
// 		... on DvSwitch { id variant host { host PNIC } }
// 	  }
// 	}
// }
// query VMs($provider: ID!){
// 	vsphereVMs(filter: {provider: $provider, cpuHotAddEnabled: false, ipAddress: "10.19.2.32", powerState: "poweredOn", memoryMB: 4096}) {
// 	  id name revision ipAddress cpuHotAddEnabled powerState memoryMB
// 	  disks { key datastore }
// 	  concerns { label }
// 	  networks {
// 		... on Network { id variant }
// 		... on DvPortGroup { id variant dvSwitch }
// 		... on DvSwitch { id variant host { host PNIC } }
// 	  }
// 	}
// }

// query VM($provider: ID!){
// 	vsphereVM(id: "vm-3344", provider: $provider) {
// 	  id name disks { key datastore } concerns { label }
// 	  networks {
// 		... on Network { id variant }
// 		... on DvPortGroup { id variant dvSwitch }
// 		... on DvSwitch { id variant host { host PNIC } }
// 	  }
// 	}
// }

// query AllFolders {
// 	vspherefolders {
// 	  ...FolderFields
// 	  ...ChildrenRecursive
// 	}
// }

// query Folders($provider: ID!) {
// 	vspherefolders(provider: $provider) {
// 	  ...FolderFields
// 	  ...ChildrenRecursive
// 	}
// }

// query Folder($provider: ID!) {
// 	vspherefolder(id: "group-s24", provider: $provider) {
// 	  ...FolderFields
// 	  ...ChildrenRecursive
// 	}
// }

// fragment FolderFields on VsphereFolder {
// 	id
// 	name
// 	provider
// 	parent
// 	children {
// 	  ... on VsphereDatacenter {id name }
// 	  ... on VsphereCluster {id name }
// 	  ... on VsphereDatastore {id name }
// 	  ... on Network {id name }
// 	  ... on DvPortGroup {id name }
// 	  ... on DvSwitch {id name }
// 	  ... on VsphereVM {id name }
// 	}
// }

// fragment ChildrenRecursive on VsphereFolder {
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
// }

// query VMandTemplateTree {
// 	vsphereProviders {
// 	  id name product
// 	  datacenters {
// 		id name
// 		children: vms {
// 		  ...FolderFields
// 		  ...ChildrenRecursive
// 		  ... on VsphereVM {id name }
// 		}
// 	  }
// 	}
// }

// query NonExistingProvider{
// 	vsphereHosts(provider: "mystery") {
// 	  id name cpuCores cpuSockets
// 	}
// }

// query HostFailure($provider: ID!){
// 	vsphereHost(id: "host-XYZ", provider: $provider) {
// 	  id name cpuCores cpuSockets
// 	}
// }

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
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/datastore"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/folder"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/host"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/network"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/resolver/vsphere/provider"
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
			Datastore: datastore.Resolver{
				Resolver: newBaseResolver(h.Container, "datastore"),
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
