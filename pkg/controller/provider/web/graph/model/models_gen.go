// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NetworkGroup interface {
	IsNetworkGroup()
}

type VsphereFolderGroup interface {
	IsVsphereFolderGroup()
}

type VsphereVMGroup interface {
	IsVsphereVMGroup()
}

type Concern struct {
	Label      string `json:"label"`
	Category   string `json:"category"`
	Assessment string `json:"assessment"`
}

type ConfigNetwork struct {
	VNICs      []*Vnic      `json:"vNICs"`
	PNICs      []*Pnic      `json:"pNICs"`
	PortGroups []*PortGroup `json:"portGroups"`
	VSwitches  []*VSwitch   `json:"vSwitches"`
}

type Device struct {
	Kind string `json:"Kind"`
}

type Disk struct {
	Key       int    `json:"key"`
	File      string `json:"file"`
	Datastore string `json:"datastore"`
	Capacity  int    `json:"capacity"`
	Shared    bool   `json:"shared"`
	Rdm       bool   `json:"rdm"`
}

type DvPortGroup struct {
	ID       string         `json:"id"`
	Variant  string         `json:"variant"`
	Name     string         `json:"name"`
	Provider string         `json:"provider"`
	Parent   *VsphereFolder `json:"parent"`
	DvSwitch string         `json:"dvSwitch"`
	Ports    []string       `json:"ports"`
	Vms      []*VsphereVM   `json:"vms"`
}

func (DvPortGroup) IsNetworkGroup()       {}
func (DvPortGroup) IsVsphereFolderGroup() {}

type DvSHost struct {
	Host string   `json:"host"`
	Pnic []string `json:"PNIC"`
}

type DvSwitch struct {
	ID         string         `json:"id"`
	Variant    string         `json:"variant"`
	Name       string         `json:"name"`
	Provider   string         `json:"provider"`
	Parent     *VsphereFolder `json:"parent"`
	Portgroups []*DvPortGroup `json:"portgroups"`
	Host       []*DvSHost     `json:"host"`
}

func (DvSwitch) IsNetworkGroup()       {}
func (DvSwitch) IsVsphereFolderGroup() {}

type Network struct {
	ID       string         `json:"id"`
	Variant  string         `json:"variant"`
	Name     string         `json:"name"`
	Provider string         `json:"provider"`
	Parent   *VsphereFolder `json:"parent"`
	Tag      string         `json:"tag"`
	Vms      []*VsphereVM   `json:"vms"`
}

func (Network) IsNetworkGroup()       {}
func (Network) IsVsphereFolderGroup() {}

type NetworkAdapter struct {
	Name      string `json:"name"`
	IPAddress string `json:"ipAddress"`
	LinkSpeed int    `json:"linkSpeed"`
	Mtu       int    `json:"mtu"`
}

type Pnic struct {
	Key       string `json:"key"`
	LinkSpeed int    `json:"linkSpeed"`
}

type PortGroup struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Vswitch string `json:"vswitch"`
}

type VMFilter struct {
	CPUHotAddEnabled *bool   `json:"cpuHotAddEnabled"`
	IPAddress        *string `json:"ipAddress"`
	PowerState       *string `json:"powerState"`
	MemoryMb         *int    `json:"memoryMB"`
}

type Vnic struct {
	Key        string `json:"key"`
	PortGroup  string `json:"portGroup"`
	DPortGroup string `json:"dPortGroup"`
	IPAddress  string `json:"ipAddress"`
	Mtu        int    `json:"mtu"`
}

type VSwitch struct {
	Key        string   `json:"key"`
	Name       string   `json:"name"`
	PortGroups []string `json:"portGroups"`
	PNICs      []string `json:"pNICs"`
}

type VsphereCluster struct {
	ID            string              `json:"id"`
	Provider      string              `json:"provider"`
	Name          string              `json:"name"`
	DatastoresIDs []string            `json:"datastoresIDs"`
	Datastores    []*VsphereDatastore `json:"datastores"`
	NetworksIDs   []string            `json:"networksIDs"`
	Networks      []NetworkGroup      `json:"networks"`
	Hosts         []*VsphereHost      `json:"hosts"`
	DasEnabled    bool                `json:"dasEnabled"`
	DasVmsIDs     []string            `json:"dasVmsIDs"`
	DasVms        []*VsphereVM        `json:"dasVms"`
	DrsEnabled    bool                `json:"drsEnabled"`
	DrsBehavior   string              `json:"drsBehavior"`
	DrsVmsIDs     []string            `json:"drsVmsIDs"`
	DrsVms        []*VsphereVM        `json:"drsVms"`
}

func (VsphereCluster) IsVsphereFolderGroup() {}

type VsphereDatacenter struct {
	ID           string              `json:"id"`
	Provider     string              `json:"provider"`
	Name         string              `json:"name"`
	ClustersID   string              `json:"clustersID"`
	Clusters     []*VsphereCluster   `json:"clusters"`
	DatastoresID string              `json:"datastoresID"`
	Datastores   []*VsphereDatastore `json:"datastores"`
	NetworksID   string              `json:"networksID"`
	Networks     []NetworkGroup      `json:"networks"`
	VmsID        string              `json:"vmsID"`
	Vms          []VsphereVMGroup    `json:"vms"`
}

func (VsphereDatacenter) IsVsphereFolderGroup() {}

type VsphereDatastore struct {
	ID          string         `json:"id"`
	Provider    string         `json:"provider"`
	Name        string         `json:"name"`
	Capacity    int            `json:"capacity"`
	Free        int            `json:"free"`
	Maintenance string         `json:"maintenance"`
	Hosts       []*VsphereHost `json:"hosts"`
	Vms         []*VsphereVM   `json:"vms"`
}

func (VsphereDatastore) IsVsphereFolderGroup() {}

type VsphereFolder struct {
	ID          string               `json:"id"`
	Provider    string               `json:"provider"`
	Name        string               `json:"name"`
	Parent      string               `json:"parent"`
	ChildrenIDs []string             `json:"childrenIDs"`
	Children    []VsphereFolderGroup `json:"children"`
}

func (VsphereFolder) IsVsphereVMGroup()     {}
func (VsphereFolder) IsVsphereFolderGroup() {}

type VsphereHost struct {
	ID              string              `json:"id"`
	Provider        string              `json:"provider"`
	Name            string              `json:"name"`
	Cluster         string              `json:"cluster"`
	ProductName     string              `json:"productName"`
	ProductVersion  string              `json:"productVersion"`
	InMaintenance   bool                `json:"inMaintenance"`
	CPUSockets      int                 `json:"cpuSockets"`
	CPUCores        int                 `json:"cpuCores"`
	Vms             []*VsphereVM        `json:"vms"`
	DatastoreIDs    []string            `json:"datastoreIDs"`
	Datastores      []*VsphereDatastore `json:"datastores"`
	Networking      *ConfigNetwork      `json:"networking"`
	NetworksIDs     []string            `json:"networksIDs"`
	Networks        []NetworkGroup      `json:"networks"`
	NetworkAdapters []*NetworkAdapter   `json:"networkAdapters"`
}

type VsphereProvider struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Type        string               `json:"type"`
	Product     string               `json:"product"`
	Datacenters []*VsphereDatacenter `json:"datacenters"`
}

type VsphereVM struct {
	ID                    string         `json:"id"`
	Provider              string         `json:"provider"`
	Name                  string         `json:"name"`
	Path                  string         `json:"path"`
	Revision              int            `json:"revision"`
	RevisionValidated     int            `json:"revisionValidated"`
	UUID                  string         `json:"uuid"`
	Firmware              string         `json:"firmware"`
	IPAddress             string         `json:"ipAddress"`
	PowerState            string         `json:"powerState"`
	CPUAffinity           []int          `json:"cpuAffinity"`
	CPUHotAddEnabled      bool           `json:"cpuHotAddEnabled"`
	CPUHotRemoveEnabled   bool           `json:"cpuHotRemoveEnabled"`
	MemoryHotAddEnabled   bool           `json:"memoryHotAddEnabled"`
	FaultToleranceEnabled bool           `json:"faultToleranceEnabled"`
	CPUCount              int            `json:"cpuCount"`
	CoresPerSocket        int            `json:"coresPerSocket"`
	MemoryMb              int            `json:"memoryMB"`
	GuestName             string         `json:"guestName"`
	BalloonedMemory       int            `json:"balloonedMemory"`
	NumaNodeAffinity      []string       `json:"numaNodeAffinity"`
	StorageUsed           int            `json:"storageUsed"`
	Snapshot              int            `json:"snapshot"`
	IsTemplate            bool           `json:"isTemplate"`
	HostID                string         `json:"hostID"`
	Host                  *VsphereHost   `json:"host"`
	Devices               []*Device      `json:"devices"`
	Disks                 []*Disk        `json:"disks"`
	NetIDs                []string       `json:"netIDs"`
	Networks              []NetworkGroup `json:"networks"`
	Concerns              []*Concern     `json:"concerns"`
}

func (VsphereVM) IsVsphereVMGroup()     {}
func (VsphereVM) IsVsphereFolderGroup() {}
