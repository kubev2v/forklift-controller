// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NetworkGroup interface {
	IsNetworkGroup()
}

type Concern struct {
	Label      string `json:"label"`
	Category   string `json:"category"`
	Assessment string `json:"assessment"`
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
	ID       string       `json:"id"`
	Variant  string       `json:"variant"`
	Name     string       `json:"name"`
	Parent   *Folder      `json:"parent"`
	DvSwitch string       `json:"dvSwitch"`
	Ports    []string     `json:"ports"`
	Vms      []*VsphereVM `json:"vms"`
}

func (DvPortGroup) IsNetworkGroup() {}

type DvSHost struct {
	Host string   `json:"host"`
	Pnic []string `json:"PNIC"`
}

type DvSwitch struct {
	ID         string         `json:"id"`
	Variant    string         `json:"variant"`
	Name       string         `json:"name"`
	Parent     *Folder        `json:"parent"`
	Portgroups []*DvPortGroup `json:"portgroups"`
	Host       []*DvSHost     `json:"host"`
}

func (DvSwitch) IsNetworkGroup() {}

type Folder struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Parent   string    `json:"parent"`
	Children []*Folder `json:"children"`
}

type Network struct {
	ID      string       `json:"id"`
	Variant string       `json:"variant"`
	Name    string       `json:"name"`
	Parent  *Folder      `json:"parent"`
	Tag     string       `json:"tag"`
	Vms     []*VsphereVM `json:"vms"`
}

func (Network) IsNetworkGroup() {}

type VsphereCluster struct {
	ID          string         `json:"id"`
	Provider    string         `json:"provider"`
	Name        string         `json:"name"`
	Hosts       []*VsphereHost `json:"hosts"`
	DasEnabled  bool           `json:"dasEnabled"`
	DasVms      []string       `json:"dasVms"`
	DrsEnabled  bool           `json:"drsEnabled"`
	DrsBehavior string         `json:"drsBehavior"`
	DrsVms      []*string      `json:"drsVms"`
}

type VsphereDatacenter struct {
	ID       string            `json:"id"`
	Provider string            `json:"provider"`
	Name     string            `json:"name"`
	Clusters []*VsphereCluster `json:"clusters"`
}

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

type VsphereHost struct {
	ID             string              `json:"id"`
	Provider       string              `json:"provider"`
	Name           string              `json:"name"`
	ProductName    string              `json:"productName"`
	ProductVersion string              `json:"productVersion"`
	InMaintenance  bool                `json:"inMaintenance"`
	CPUSockets     int                 `json:"cpuSockets"`
	CPUCores       int                 `json:"cpuCores"`
	Vms            []*VsphereVM        `json:"vms"`
	DatastoreIDs   []string            `json:"datastoreIDs"`
	Datastores     []*VsphereDatastore `json:"datastores"`
}

type VsphereProvider struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Type        string               `json:"type"`
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
	IPAddress             string         `json:"ipAddress"`
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
