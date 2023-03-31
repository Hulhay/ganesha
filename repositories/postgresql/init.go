package postgresql

import (
	"context"

	"github.com/hulhay/ganesha/lib/config"
	"github.com/hulhay/ganesha/lib/database"
	"github.com/hulhay/ganesha/models"
)

type genreRepository struct {
	qry database.SqlxDatabase
}

type GenreRepository interface {
	GetGenresFromDB(ctx context.Context) ([]models.Genre, error)
}

func NewGenreRepository(cfg *config.Config) GenreRepository {
	return &genreRepository{
		qry: cfg.DB(),
	}
}
