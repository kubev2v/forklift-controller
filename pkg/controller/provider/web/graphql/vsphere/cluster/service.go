package cluster

import graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"

type Cluster interface {
	Get(id string, provider string) (*graphmodel.VsphereCluster, error)
	GetByDatacenter(datacenterId, provider string) ([]*graphmodel.VsphereCluster, error)
	List(provider string) ([]*graphmodel.VsphereCluster, error)
}
