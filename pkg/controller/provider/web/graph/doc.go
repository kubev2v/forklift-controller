// The graph packages provides graphql services for essentially querying the controller inventory

// Query examples:

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
// 				   ... on VsphereNetwork { id variant }
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
// 				   ... on VsphereNetwork { id variant }
// 				   ... on DvPortGroup { id variant dvSwitch }
// 				   ... on DvSwitch { id variant host { host PNIC } }
// 				}
// 			  }
// 			}
// 		  }
// 		}
// 		... on OvirtDatacenter {
//       id provider name
//       clusters { id kind name }
//     }
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
// 		  ... on VsphereNetwork { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
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
// 		  ... on VsphereNetwork { id variant }
// 		  ... on DvPortGroup { id variant dvSwitch }
// 		  ... on DvSwitch { id variant host { host PNIC } }
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
// }

// query Datacenter($provider: ID!){
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
// }

// query AllClusters{
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
// }

// query Clusters($provider: ID!){
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
// }

// query AllHosts {
// 	hosts {
// 	  ... on VsphereHost {
// 		id name provider cpuCores cpuSockets
// 		datastores { id name }
// 		vspherevms: vms { id name firmware }
// 	  }
// 	  ... on OvirtHost {
//       id provider name
//     ovirtvms: vms {
//       id kind provider description cpuCores cpuSockets cpuAffinity { set cpu } memory
//       }
//     }
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
// }

// query AllStorages {
// 	storages {
//     ... on VsphereDatastore {
//    	  id name provider
// 	    hosts { id name }
// 	    vms { id name firmware }
//     }
//     ... on OvirtStorageDomain {
//       id name provider
//     }
// 	}
// }

// query Datastores($provider: ID!) {
// 	storages (provider: $provider) {
//     ... on VsphereDatastore {
// 	    id name
// 	    hosts { id name }
// 	    vms { id name firmware }
//     }
// 	}
// }

// query Datastore($provider: ID!) {
// 	storage (id: "datastore-45", provider: $provider) {
//     ... on VsphereDatastore {
//       id name
//       hosts { id name }
//       vms { id name firmware }
// 	  }
//   }
// }

// query AllNetworks{
//   networks {
// 	  ... on VsphereNetwork { id provider tag }
// 	  ... on DvPortGroup { id provider dvSwitch }
// 	  ... on DvSwitch { id provider host { host PNIC }}
// 	}
// }

// query Networks($provider: ID!){
// 	networks(provider: $provider) {
// 	  ... on VsphereNetwork { id variant tag }
// 	  ... on DvPortGroup { id variant dvSwitch }
// 	  ... on DvSwitch { id variant host { host PNIC }}
// 	}
// }

// query Network($provider: ID!){
// 	network(id: "dvportgroup-56", provider: $provider) {
// 	  ... on VsphereNetwork { id variant }
// 	  ... on DvPortGroup { id variant dvSwitch }
// 	  ... on DvSwitch { id variant host { host PNIC }}
// 	}
// }

// query AllVMs {
// 	vms {
//  	  ... on VsphereVM {
//   	  id provider name revision ipAddress cpuHotAddEnabled powerState memoryMB
// 	    disks { key datastore }
// 	    concerns { label }
// 	    networks {
// 	    	... on VsphereNetwork { id variant }
//   	  	... on DvPortGroup { id variant dvSwitch }
//   	  	... on DvSwitch { id variant host { host PNIC } }
// 	    }
//     }
//     ... on OvirtVM {
//       id kind provider description cpuCores cpuSockets cpuAffinity { set cpu } memory
//     }
// 	}
// }

// query VMs($provider: ID!){
// 	vms(provider: $provider, filter: {cpuHotAddEnabled: false, ipAddress: "10.19.2.32", powerState: "poweredOn", memoryMB: 4096}) {
// 	  ... on VsphereVM {
//       id name revision ipAddress cpuHotAddEnabled powerState memoryMB
// 	    disks { key datastore }
// 	    concerns { label }
// 	    networks {
// 	    	... on VsphereNetwork { id variant }
// 		    ... on DvPortGroup { id variant dvSwitch }
// 		    ... on DvSwitch { id variant host { host PNIC } }
//       }
// 	  }
// 	}
// }

// query VM($provider: ID!){
// 	vm (id: "vm-3344", provider: $provider) {
//     ... on VsphereVM {
//   	  id name disks { key datastore } concerns { label }
// 	    networks {
// 		    ... on VsphereNetwork { id variant }
// 		    ... on DvPortGroup { id variant dvSwitch }
// 		    ... on DvSwitch { id variant host { host PNIC } }
//       }
// 	  }
// 	}
// }

// query AllFolders {
// 	vspherefolders {
//     id
// 	  name
// 	  provider
//     kind
// 	  children {
//       ...ClusterFields
//       ...DatastoreFields
//       ...NetworkFields
//       ...DvPortGroupFields
//       ...DvSwitchFields
//       ...VMFields
//       ...FolderFields
//       ...FolderRecursive
//    	}
//   }
// }

// query Folder {
// 	vspherefolder(id:"group-n25", provider:"866883b4-4526-4868-807d-8aa8d8d7c72e") {
//     id
// 	  name
// 	  provider
//     kind
// 	  children {
//       # ...ClusterFields
//       # ...DatastoreFields
//       ...NetworkFields
//       ...DvPortGroupFields
//       ...DvSwitchFields
//       # ...VMFields
//       ...FolderFields
//       ...FolderRecursive
//    	}
//   }
// }

// query VMandTemplateTree {
// 	datacenter(id: "datacenter-21", provider: "866883b4-4526-4868-807d-8aa8d8d7c72e") {
// 	  ... on VsphereDatacenter {
// 		  id name
// 		  children: vms {
//         id name kind
//         children {
//           ...VMFields
//           ...FolderFields
//           ...FolderRecursive
//       	}
//       }
//     }
//   }
// }

// fragment ClusterFields on VsphereCluster {
// 	id
// 	name
// 	provider
//   kind
// }

// fragment DatastoreFields on VsphereDatastore {
// 	id
// 	name
// 	provider
//   kind
// }

// fragment NetworkFields on VsphereNetwork {
// 	id
// 	name
// 	provider
//   kind
//   tag
// }

// fragment DvPortGroupFields on DvPortGroup {
// 	id
// 	variant
//   dvSwitch
// }

// fragment DvSwitchFields on DvSwitch{
// 	id
// 	variant
//   host { host PNIC }
// }

// fragment VMFields on VsphereVM {
// 	id
// 	name
// 	provider
//   kind
// }

// fragment FolderFields on VsphereFolder {
// 	id
// 	name
// 	provider
//   kind
//   children {
//     # It seems only Vsphere VM folders are recursive
//     ...VMFields
//   }
// }

// fragment FolderRecursive on VsphereFolder {
// 	children {
// 	  ...FolderFields
// 	  ... on VsphereFolder {
// 		  ...FolderFields
// 	  	children {
// 	    	...FolderFields
// 	    	... on VsphereFolder {
// 	      	...FolderFields
// 	      	children {
// 	        	...FolderFields
// 	        	... on VsphereFolder {
// 	  	      	...FolderFields
//               children {
// 	            	...FolderFields
// 	            	... on VsphereFolder {
// 	  	          	...FolderFields
// 	             	}
// 	          	}
// 	        	}
// 	      	}
// 	  	  }
// 	  	}
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
