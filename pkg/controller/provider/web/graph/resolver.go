package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
)

type Resolver struct {
	vsphereProvider []*model.VsphereProvider
}
