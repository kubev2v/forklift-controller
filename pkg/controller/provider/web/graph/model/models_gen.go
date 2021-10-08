// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Concern struct {
	Label      string `json:"label"`
	Category   string `json:"category"`
	Assessment string `json:"assessment"`
}

type Device struct {
	Kind string `json:"Kind"`
}

type Disk struct {
	ID        string `json:"id"`
	Kind      string `json:"kind"`
	File      string `json:"file"`
	Datastore string `json:"datastore"`
	Capacity  string `json:"capacity"`
	Shared    bool   `json:"shared"`
	Rdm       bool   `json:"rdm"`
}

type VsphereCluster struct {
	ID          string         `json:"id"`
	Provider    string         `json:"provider"`
	Kind        string         `json:"kind"`
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
	Kind     string            `json:"kind"`
	Name     string            `json:"name"`
	Provider string            `json:"provider"`
	Clusters []*VsphereCluster `json:"clusters"`
}

type VsphereHost struct {
	ID             string       `json:"id"`
	Provider       string       `json:"provider"`
	Kind           string       `json:"kind"`
	Name           string       `json:"name"`
	ProductName    string       `json:"productName"`
	ProductVersion string       `json:"productVersion"`
	InMaintenance  bool         `json:"inMaintenance"`
	CPUSockets     int          `json:"cpuSockets"`
	CPUCores       int          `json:"cpuCores"`
	Vms            []*VsphereVM `json:"vms"`
}

type VsphereProvider struct {
	ID          string               `json:"id"`
	Kind        string               `json:"kind"`
	Name        string               `json:"name"`
	Datacenters []*VsphereDatacenter `json:"datacenters"`
}

type VsphereVM struct {
	ID                    string       `json:"id"`
	Kind                  string       `json:"kind"`
	Name                  string       `json:"name"`
	Path                  string       `json:"path"`
	Revision              int          `json:"revision"`
	SelfLink              string       `json:"selfLink"`
	UUID                  string       `json:"uuid"`
	Firmware              string       `json:"firmware"`
	PowerState            string       `json:"powerState"`
	CPUHotAddEnabled      bool         `json:"cpuHotAddEnabled"`
	CPUHotRemoveEnabled   bool         `json:"cpuHotRemoveEnabled"`
	MemoryHotAddEnabled   bool         `json:"memoryHotAddEnabled"`
	FaultToleranceEnabled bool         `json:"faultToleranceEnabled"`
	CPUCount              int          `json:"cpuCount"`
	CoresPerSocket        int          `json:"coresPerSocket"`
	MemoryMb              int          `json:"memoryMB"`
	GuestName             string       `json:"guestName"`
	BalloonedMemory       int          `json:"balloonedMemory"`
	IPAddress             string       `json:"ipAddress"`
	StorageUsed           int          `json:"storageUsed"`
	NumaNodeAffinity      []string     `json:"numaNodeAffinity"`
	Devices               []*Device    `json:"devices"`
	CPUAffinity           []int        `json:"cpuAffinity"`
	Host                  *VsphereHost `json:"host"`
	RevisionAnalyzed      int          `json:"revisionAnalyzed"`
	Disks                 []*Disk      `json:"disks"`
	Concerns              []*Concern   `json:"concerns"`
}
