package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
)

func (r *queryResolver) VsphereProviders(ctx context.Context) ([]*model.VsphereProvider, error) {
	var providers []*model.VsphereProvider
	dummyProvider := model.VsphereProvider{
		Name:    "our dummy provider",
		Kind:    "Provider",
		Product: "VMware",
	}
	providers = append(providers, &dummyProvider)
	return providers, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
