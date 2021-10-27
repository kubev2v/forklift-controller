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
// 				 networks {
// 				   ... on Network { id variant }
// 				   ... on DvPortGroup { id variant dvSwitch }
// 				   ... on DvSwitch { id variant host { host PNIC } }
// 				}
// 			  }
// 			}
// 		  }
// 		}
// 		... on OvirtDatacenter {id provider name }
// 	  }
// 	}
// }

// query ProviderTree($provider: ID!) {
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
// 				 networks {
// 				   ... on Network { id variant }
// 				   ... on DvPortGroup { id variant dvSwitch }
// 				   ... on DvSwitch { id variant host { host PNIC } }
// 				}
// 			  }
// 			}
// 		  }
// 		}
// 		... on OvirtDatacenter { id provider name }
// 	  }
// 	}
// }

// query AllDatacenters {
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
// 		  ... on Network { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		  }
// 		}
// 		datastores { id name }
// 		networks {
// 		  ... on Network { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		}
// 		vms { ... on VsphereVM { id name firmware } }
// 	  }
// 	  ... on OvirtDatacenter { id provider name }
// 	}
// }

// query Datacenters($provider: ID!) {
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
// 		  ... on Network { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		  }
// 		}
// 		datastores { id name }
// 		networks {
// 		  ... on Network { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
// 		}
// 		vms { ... on VsphereVM { id name firmware }}
// 	  }
// 	  ... on OvirtDatacenter { id provider name }
// 	}
// }

// query Datacenter($provider: ID!){
// 	datacenter(id: "datacenter-21", provider: $provider) {
// 	  ... on VsphereDatacenter {
// 		id name
// 		clusters { id name }
// 		datastores { id name }
// 		networks {
// 		  ... on Network { id variant name  }
// 		  ... on DvPortGroup { id variant name dvSwitch }
// 		  ... on DvSwitch { id variant name host { host PNIC } }
// 		}
// 		vms { ... on VsphereVM { id name firmware }}
// 	  }
// 	  ... on OvirtDatacenter { id provider name}
// 	}
// }

// query AllClusters{
// 	clusters {
// 	  ... on VsphereCluster {
// 		id kind name provider dasEnabled drsEnabled drsBehavior
// 		hosts { id name cpuCores cpuSockets provider }
// 		networks {
// 		  ... on Network { id variant name }
// 		  ... on DvPortGroup { id variant name dvSwitch }
// 		  ... on DvSwitch { id variant name host { host } }
// 		}
// 		datastores { id name }
// 	  }
// 	  ... on OvirtCluster { id kind provider name }
// 	}
// }

// query Clusters($provider: ID!){
// 	clusters(provider: $provider) {
// 	  ... on VsphereCluster {
// 		 id kind name dasEnabled drsEnabled drsBehavior
// 		 hosts { id name cpuCores cpuSockets provider }
// 		 networks {
// 		   ... on Network { id variant name }
// 		   ... on DvPortGroup { id variant name dvSwitch }
// 		   ... on DvSwitch { id variant name host { host } }
// 		 }
// 		 datastores { id name }
// 	  }
// 	  ... on OvirtCluster { id kind provider name }
// 	}
// }

// query AllHosts {
// 	hosts {
// 	  ... on VsphereHost {
// 		id name provider cpuCores cpuSockets
// 		datastores { id name }
// 		vms { id name firmware }
// 	  }
// 	  ... on OvirtHost { id provider name }
// 	}
// }

// query Hosts($provider: ID!){
// 	hosts(provider: $provider) {
// 	  ... on VsphereHost {
// 		id name cpuCores cpuSockets
// 		datastores { id name }
// 		vms { id name firmware }
// 	  }
// 	  ... on OvirtHost  { id provider name }
// 	}
// }

// query Host($provider: ID!){
// 	host(id: "host-44", provider: $provider) {
// 	  ... on VsphereHost {
// 		id name cpuCores cpuSockets
// 		datastores { id }
// 		networks {
// 		  ... on Network { id variant name }
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
// 	  id name
// 	  hosts { id name }
// 	  vms { id name firmware }
// 	}
// }

// query Datastore($provider: ID!) {
// 	vsphereDatastore(id: "datastore-45", provider: $provider) {
// 	  id name
// 	  hosts { id name }
// 	  vms { id name firmware }
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
// 	vsphereVMs(provider: $provider, filter: {cpuHotAddEnabled: false, ipAddress: "10.19.2.32", powerState: "poweredOn", memoryMB: 4096}) {
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

//   fragment FolderFields on VsphereFolder {
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
// }

// query NonExistingProvider{
// 	hosts(provider: "mystery") {
// 	  ... on VsphereHost {
// 		id name cpuCores cpuSockets
// 	  }
// 	}
// }

// query HostFailure($provider: ID!){
// 	host(id: "host-XYZ", provider: $provider) {
// 	  ... on VsphereHost {
// 		id name cpuCores cpuSockets
// 	  }
// 	}
// }

package graph

import (
	"github.com/konveyor/controller/pkg/inventory/container"
	libweb "github.com/konveyor/controller/pkg/inventory/web"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/base"
)

//
// Build all handlers.
func Handlers(container *container.Container) []libweb.RequestHandler {
	return []libweb.RequestHandler{
		&GraphHandler{
			Handler: base.Handler{
				Container: container,
			},
		},
	}
}
