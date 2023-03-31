package genre

import (
	"github.com/hulhay/ganesha/graph/gqlmodel"
	"github.com/hulhay/ganesha/models"
)

func buildGetGenresResponse(genres []models.Genre) []*gqlmodel.Genre {
	result := []*gqlmodel.Genre{}
	for _, genre := range genres {
		result = append(result, buildGetGenreResponse(genre))
	}
	return result
}

func buildGetGenreResponse(genre models.Genre) *gqlmodel.Genre {
	return &gqlmodel.Genre{
		GenreUUID: genre.Uuid,
		GenreName: genre.Name,
	}
}
