package postgresql

import (
	"context"

	"github.com/hulhay/ganesha/models"
)

func (p *genreRepository) GetGenresFromDB(ctx context.Context) ([]models.Genre, error) {
	res := []models.Genre{}
	err := p.qry.Read().SelectContext(ctx, &res, getGenresQuery)
	if err != nil {
		return res, err
	}

	return res, nil
}
