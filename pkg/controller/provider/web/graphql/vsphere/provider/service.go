package provider

import graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"

type Provider interface {
	Get(id string) (*graphmodel.VsphereProvider, error)
	List() ([]*graphmodel.VsphereProvider, error)
}
