// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

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
	ID             string `json:"id"`
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	ProductName    string `json:"productName"`
	ProductVersion string `json:"productVersion"`
	InMaintenance  bool   `json:"inMaintenance"`
	CPUSockets     int    `json:"cpuSockets"`
	CPUCores       int    `json:"cpuCores"`
}

type VsphereProvider struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}
