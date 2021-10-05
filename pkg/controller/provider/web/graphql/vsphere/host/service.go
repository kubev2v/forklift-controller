package host

import graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"

type Host interface {
	Get(id string, provider string) (*graphmodel.VsphereHost, error)
	GetByCluster(clusterId, provider string) ([]*graphmodel.VsphereHost, error)
	List(provider string) ([]*graphmodel.VsphereHost, error)
}
