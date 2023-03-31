package resolver

import (
	"context"

	"github.com/hulhay/ganesha/graph/gqlmodel"
)

type genreServiceProvider interface {
	GetGenres(ctx context.Context) ([]*gqlmodel.Genre, error)
}

// Resolver main struct contain all queries related
type Resolver struct {
	genreService genreServiceProvider
}

// Options parameter used to create resolver
type Options struct {
	GenreService genreServiceProvider
}

// New creating new resolver.
func New(args Options) *Resolver {
	return &Resolver{
		genreService: args.GenreService,
	}
}
