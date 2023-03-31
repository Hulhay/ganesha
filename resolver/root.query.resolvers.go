package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"

	"github.com/hulhay/ganesha/graph/generated"
	"github.com/hulhay/ganesha/graph/gqlmodel"
)

// Health is the resolver for the health field.
func (r *queryResolver) Health(ctx context.Context) (*gqlmodel.GeneralResponse, error) {
	return &gqlmodel.GeneralResponse{
		Message:   "Ganesha is running",
		IsSuccess: true,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
