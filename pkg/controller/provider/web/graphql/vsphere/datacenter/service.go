package datacenter

import graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"

type Datacenter interface {
	Get(id, provider string) (*graphmodel.VsphereDatacenter, error)
	List(provider string) ([]*graphmodel.VsphereDatacenter, error)
}
