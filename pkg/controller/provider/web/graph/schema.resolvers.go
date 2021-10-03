package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/generated"
	graphmodel "github.com/konveyor/forklift-controller/pkg/controller/provider/web/graph/model"
)

func (r *queryResolver) VsphereProviders(ctx context.Context) ([]*graphmodel.VsphereProvider, error) {
	return r.Resolver.Provider.List()
}

func (r *queryResolver) VsphereHosts(ctx context.Context, provider string) ([]*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.List(provider)
}

func (r *queryResolver) VsphereHost(ctx context.Context, id string, provider string) (*graphmodel.VsphereHost, error) {
	return r.Resolver.Host.Get(id, provider)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
