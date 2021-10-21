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

package web

import (
	"github.com/konveyor/controller/pkg/inventory/container"
	libweb "github.com/konveyor/controller/pkg/inventory/web"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/ocp"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/ovirt"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/vsphere"
	"github.com/konveyor/forklift-controller/pkg/settings"
)

//
// All handlers.
func All(container *container.Container) (all []libweb.RequestHandler) {
	all = []libweb.RequestHandler{
		&libweb.SchemaHandler{},
		&ProviderHandler{
			Handler: base.Handler{
				Container: container,
			},
		},
	}
	all = append(
		all,
		ocp.Handlers(container)...)
	all = append(
		all,
		vsphere.Handlers(container)...)
	all = append(
		all,
		ovirt.Handlers(container)...)
	if settings.Settings.Inventory.GrahpQLEnabled {
		all = append(
			all,
			graph.Handlers(container)...)
	}
	return
}
