package host

import graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"

// type HostItem struct {
// 	Id   string
// 	Name string
// }

type Host interface {
	// Initialise() error
	Get(id string, provider string) (*graphmodel.VsphereHost, error)
	// List() ([]HostItem, error)
}
