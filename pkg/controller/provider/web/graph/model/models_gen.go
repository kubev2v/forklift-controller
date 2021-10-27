// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cluster interface {
	IsCluster()
}

type Datacenter interface {
	IsDatacenter()
}

type Host interface {
	IsHost()
}

type Network interface {
	IsNetwork()
}

type Storage interface {
	IsStorage()
}

type VM interface {
	IsVM()
}

type VsphereFolderGroup interface {
	IsVsphereFolderGroup()
}

type VsphereNetworkGroup interface {
	IsVsphereNetworkGroup()
}

type VsphereVMGroup interface {
	IsVsphereVMGroup()
}

type Cdrom struct {
	ID   string `json:"id"`
	File string `json:"file"`
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

type CPUPinning struct {
	Set int `json:"set"`
	CPU int `json:"cpu"`
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

type DiskAttachment struct {
	ID              string `json:"id"`
	Interface       string `json:"interface"`
	ScsiReservation bool   `json:"scsiReservation"`
	Disk            string `json:"disk"`
}

type DiskProfile struct {
	ID            string              `json:"id"`
	Provider      string              `json:"provider"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	StorageDomain *OvirtStorageDomain `json:"storageDomain"`
	Qos           string              `json:"qos"`
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

func (DvPortGroup) IsNetwork()             {}
func (DvPortGroup) IsVsphereNetworkGroup() {}
func (DvPortGroup) IsVsphereFolderGroup()  {}

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

func (DvSwitch) IsNetwork()             {}
func (DvSwitch) IsVsphereNetworkGroup() {}
func (DvSwitch) IsVsphereFolderGroup()  {}

type Guest struct {
	Distribution string `json:"distribution"`
	FullVersion  string `json:"fullVersion"`
}

type HostDevice struct {
	Capability string `json:"capability"`
	Product    string `json:"product"`
	Vendor     string `json:"vendor"`
}

type HostNic struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LinkSpeed int    `json:"linkSpeed"`
	Mtu       int    `json:"mtu"`
	VLan      string `json:"vLan"`
}

type NetworkAdapter struct {
	Name      string `json:"name"`
	IPAddress string `json:"ipAddress"`
	LinkSpeed int    `json:"linkSpeed"`
	Mtu       int    `json:"mtu"`
}

type NetworkAttachment struct {
	ID      string `json:"id"`
	Network string `json:"network"`
}

type OvirtCluster struct {
	ID            string       `json:"id"`
	Provider      string       `json:"provider"`
	Kind          string       `json:"kind"`
	Name          string       `json:"name"`
	DataCenter    string       `json:"dataCenter"`
	HaReservation bool         `json:"haReservation"`
	KsmEnabled    bool         `json:"ksmEnabled"`
	BiosType      string       `json:"biosType"`
	Hosts         []*OvirtHost `json:"hosts"`
	Vms           []*OvirtVM   `json:"vms"`
}

func (OvirtCluster) IsCluster() {}

type OvirtDatacenter struct {
	ID       string                `json:"id"`
	Provider string                `json:"provider"`
	Kind     string                `json:"kind"`
	Name     string                `json:"name"`
	Clusters []*OvirtCluster       `json:"clusters"`
	Storages []*OvirtStorageDomain `json:"storages"`
	Networks []*OvirtNetwork       `json:"networks"`
}

func (OvirtDatacenter) IsDatacenter() {}

type OvirtHost struct {
	ID                 string               `json:"id"`
	Provider           string               `json:"provider"`
	Kind               string               `json:"kind"`
	Name               string               `json:"name"`
	Cluster            string               `json:"cluster"`
	Status             string               `json:"status"`
	ProductName        string               `json:"productName"`
	ProductVersion     string               `json:"productVersion"`
	InMaintenance      bool                 `json:"inMaintenance"`
	CPUSockets         int                  `json:"cpuSockets"`
	CPUCores           int                  `json:"cpuCores"`
	NetworkAttachments []*NetworkAttachment `json:"networkAttachments"`
	Nics               []*HostNic           `json:"nics"`
	Vms                []*OvirtVM           `json:"vms"`
}

func (OvirtHost) IsHost() {}

type OvirtNICProfile struct {
	ID            string      `json:"id"`
	Provider      string      `json:"provider"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Network       string      `json:"network"`
	PortMirroring bool        `json:"portMirroring"`
	NetworkFilter string      `json:"networkFilter"`
	Qos           string      `json:"qos"`
	Properties    []*Property `json:"properties"`
	PassThrough   bool        `json:"passThrough"`
}

type OvirtNetwork struct {
	ID          string   `json:"id"`
	Provider    string   `json:"provider"`
	Kind        string   `json:"kind"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	DataCenter  string   `json:"dataCenter"`
	VLan        string   `json:"vLan"`
	Usages      []string `json:"usages"`
	Profiles    []string `json:"profiles"`
}

func (OvirtNetwork) IsNetwork() {}

type OvirtStorageDomain struct {
	ID          string `json:"id"`
	Provider    string `json:"provider"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DataCenter  string `json:"dataCenter"`
	Type        string `json:"type"`
	StorageType string `json:"storageType"`
	Available   int    `json:"available"`
	Used        int    `json:"used"`
}

func (OvirtStorageDomain) IsStorage() {}

type OvirtVM struct {
	ID                          string            `json:"id"`
	Kind                        string            `json:"kind"`
	Provider                    string            `json:"provider"`
	Name                        string            `json:"name"`
	Description                 string            `json:"description"`
	Cluster                     string            `json:"cluster"`
	Host                        string            `json:"host"`
	RevisionValidated           int               `json:"revisionValidated"`
	PolicyVersion               int               `json:"policyVersion"`
	GuestName                   string            `json:"guestName"`
	CPUSockets                  int               `json:"cpuSockets"`
	CPUCores                    int               `json:"cpuCores"`
	CPUThreads                  int               `json:"cpuThreads"`
	CPUAffinity                 []*CPUPinning     `json:"cpuAffinity"`
	CPUShares                   int               `json:"cpuShares"`
	Memory                      int               `json:"memory"`
	BalloonedMemory             bool              `json:"balloonedMemory"`
	Bios                        string            `json:"bios"`
	Display                     string            `json:"display"`
	IOThreads                   int               `json:"iOThreads"`
	StorageErrorResumeBehaviour string            `json:"storageErrorResumeBehaviour"`
	HaEnabled                   bool              `json:"haEnabled"`
	UsbEnabled                  bool              `json:"usbEnabled"`
	BootMenuEnabled             bool              `json:"bootMenuEnabled"`
	PlacementPolicyAffinity     string            `json:"placementPolicyAffinity"`
	Timezone                    string            `json:"timezone"`
	Status                      string            `json:"status"`
	Stateless                   string            `json:"stateless"`
	SerialNumber                string            `json:"serialNumber"`
	HasIllegalImages            bool              `json:"hasIllegalImages"`
	NumaNodeAffinity            []string          `json:"numaNodeAffinity"`
	LeaseStorageDomain          string            `json:"leaseStorageDomain"`
	DiskAttachments             []*DiskAttachment `json:"diskAttachments"`
	Nics                        []string          `json:"nics"`
	HostDevices                 []*HostDevice     `json:"hostDevices"`
	Cdroms                      []*Cdrom          `json:"cdroms"`
	WatchDogs                   []*WatchDog       `json:"watchDogs"`
	Properties                  []*Property       `json:"properties"`
	Snapshots                   []*Snapshot       `json:"snapshots"`
	Concerns                    []*Concern        `json:"concerns"`
	Guest                       *Guest            `json:"guest"`
	OsType                      string            `json:"osType"`
}

func (OvirtVM) IsVM() {}

type Pnic struct {
	Key       string `json:"key"`
	LinkSpeed int    `json:"linkSpeed"`
}

type PortGroup struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Vswitch string `json:"vswitch"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Provider struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Product     string       `json:"product"`
	Datacenters []Datacenter `json:"datacenters"`
}

type Snapshot struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	Type          string `json:"type"`
	PersistMemory bool   `json:"persistMemory"`
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
	ID            string                `json:"id"`
	Provider      string                `json:"provider"`
	Kind          string                `json:"kind"`
	Name          string                `json:"name"`
	DatastoresIDs []string              `json:"datastoresIDs"`
	Datastores    []*VsphereDatastore   `json:"datastores"`
	NetworksIDs   []string              `json:"networksIDs"`
	Networks      []VsphereNetworkGroup `json:"networks"`
	Hosts         []*VsphereHost        `json:"hosts"`
	DasEnabled    bool                  `json:"dasEnabled"`
	DasVmsIDs     []string              `json:"dasVmsIDs"`
	DasVms        []*VsphereVM          `json:"dasVms"`
	DrsEnabled    bool                  `json:"drsEnabled"`
	DrsBehavior   string                `json:"drsBehavior"`
	DrsVmsIDs     []string              `json:"drsVmsIDs"`
	DrsVms        []*VsphereVM          `json:"drsVms"`
}

func (VsphereCluster) IsCluster()            {}
func (VsphereCluster) IsVsphereFolderGroup() {}

type VsphereDatacenter struct {
	ID           string                `json:"id"`
	Provider     string                `json:"provider"`
	Kind         string                `json:"kind"`
	Name         string                `json:"name"`
	ClustersID   string                `json:"clustersID"`
	Clusters     []*VsphereCluster     `json:"clusters"`
	DatastoresID string                `json:"datastoresID"`
	Datastores   []*VsphereDatastore   `json:"datastores"`
	NetworksID   string                `json:"networksID"`
	Networks     []VsphereNetworkGroup `json:"networks"`
	VmsID        string                `json:"vmsID"`
	Vms          []VsphereVMGroup      `json:"vms"`
}

func (VsphereDatacenter) IsDatacenter()         {}
func (VsphereDatacenter) IsVsphereFolderGroup() {}

type VsphereDatastore struct {
	ID          string         `json:"id"`
	Provider    string         `json:"provider"`
	Kind        string         `json:"kind"`
	Name        string         `json:"name"`
	Capacity    int            `json:"capacity"`
	Free        int            `json:"free"`
	Maintenance string         `json:"maintenance"`
	Hosts       []*VsphereHost `json:"hosts"`
	Vms         []*VsphereVM   `json:"vms"`
}

func (VsphereDatastore) IsStorage()            {}
func (VsphereDatastore) IsVsphereFolderGroup() {}

type VsphereFolder struct {
	ID          string               `json:"id"`
	Provider    string               `json:"provider"`
	Name        string               `json:"name"`
	Parent      string               `json:"parent"`
	ChildrenIDs []string             `json:"childrenIDs"`
	Children    []VsphereFolderGroup `json:"children"`
}

func (VsphereFolder) IsVM()                 {}
func (VsphereFolder) IsVsphereVMGroup()     {}
func (VsphereFolder) IsVsphereFolderGroup() {}

type VsphereHost struct {
	ID              string                `json:"id"`
	Provider        string                `json:"provider"`
	Kind            string                `json:"kind"`
	Name            string                `json:"name"`
	Cluster         string                `json:"cluster"`
	ProductName     string                `json:"productName"`
	ProductVersion  string                `json:"productVersion"`
	InMaintenance   bool                  `json:"inMaintenance"`
	CPUSockets      int                   `json:"cpuSockets"`
	CPUCores        int                   `json:"cpuCores"`
	Vms             []*VsphereVM          `json:"vms"`
	DatastoreIDs    []string              `json:"datastoreIDs"`
	Datastores      []*VsphereDatastore   `json:"datastores"`
	Networking      *ConfigNetwork        `json:"networking"`
	NetworksIDs     []string              `json:"networksIDs"`
	Networks        []VsphereNetworkGroup `json:"networks"`
	NetworkAdapters []*NetworkAdapter     `json:"networkAdapters"`
}

func (VsphereHost) IsHost() {}

type VsphereNetwork struct {
	ID       string         `json:"id"`
	Provider string         `json:"provider"`
	Kind     string         `json:"kind"`
	Variant  string         `json:"variant"`
	Name     string         `json:"name"`
	Parent   *VsphereFolder `json:"parent"`
	Tag      string         `json:"tag"`
	Vms      []*VsphereVM   `json:"vms"`
}

func (VsphereNetwork) IsNetwork()             {}
func (VsphereNetwork) IsVsphereNetworkGroup() {}
func (VsphereNetwork) IsVsphereFolderGroup()  {}

type VsphereVM struct {
	ID                    string                `json:"id"`
	Kind                  string                `json:"kind"`
	Provider              string                `json:"provider"`
	Name                  string                `json:"name"`
	Path                  string                `json:"path"`
	Revision              int                   `json:"revision"`
	RevisionValidated     int                   `json:"revisionValidated"`
	UUID                  string                `json:"uuid"`
	Firmware              string                `json:"firmware"`
	IPAddress             string                `json:"ipAddress"`
	PowerState            string                `json:"powerState"`
	CPUAffinity           []int                 `json:"cpuAffinity"`
	CPUHotAddEnabled      bool                  `json:"cpuHotAddEnabled"`
	CPUHotRemoveEnabled   bool                  `json:"cpuHotRemoveEnabled"`
	MemoryHotAddEnabled   bool                  `json:"memoryHotAddEnabled"`
	FaultToleranceEnabled bool                  `json:"faultToleranceEnabled"`
	CPUCount              int                   `json:"cpuCount"`
	CoresPerSocket        int                   `json:"coresPerSocket"`
	MemoryMb              int                   `json:"memoryMB"`
	GuestName             string                `json:"guestName"`
	BalloonedMemory       int                   `json:"balloonedMemory"`
	NumaNodeAffinity      []string              `json:"numaNodeAffinity"`
	StorageUsed           int                   `json:"storageUsed"`
	Snapshot              int                   `json:"snapshot"`
	IsTemplate            bool                  `json:"isTemplate"`
	HostID                string                `json:"hostID"`
	Host                  *VsphereHost          `json:"host"`
	Devices               []*Device             `json:"devices"`
	Disks                 []*Disk               `json:"disks"`
	NetIDs                []string              `json:"netIDs"`
	Networks              []VsphereNetworkGroup `json:"networks"`
	Concerns              []*Concern            `json:"concerns"`
}

func (VsphereVM) IsVM()                 {}
func (VsphereVM) IsVsphereVMGroup()     {}
func (VsphereVM) IsVsphereFolderGroup() {}

type WatchDog struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Model  string `json:"model"`
}
