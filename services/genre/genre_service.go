package genre

import (
	"context"
	"log"

	"github.com/hulhay/ganesha/graph/gqlmodel"
	"github.com/hulhay/ganesha/lib/utils"
	"github.com/hulhay/ganesha/models"
)

//go:generate mockgen -source=genre_service.go -destination=genre_service_mock.go -package=genre

type GenreService struct {
	genreRepo genreRepo
}

type GenreOptions struct {
	GenreRepo genreRepo
}

type genreRepo interface {
	GetGenresFromDB(ctx context.Context) ([]models.Genre, error)
}

func NewGenreService(option GenreOptions) *GenreService {
	return &GenreService{
		genreRepo: option.GenreRepo,
	}
}

func (gs *GenreService) GetGenres(ctx context.Context) ([]*gqlmodel.Genre, error) {
	genres, err := gs.genreRepo.GetGenresFromDB(ctx)
	if err != nil {
		log.Printf("Error get genres from db : %v\n", err.Error())
		return nil, utils.ErrFetchData
	}
	return buildGetGenresResponse(genres), nil
}
