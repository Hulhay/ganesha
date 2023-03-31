package genre

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/hulhay/ganesha/graph/gqlmodel"
	"github.com/hulhay/ganesha/lib/utils"
	"github.com/hulhay/ganesha/models"
	"github.com/stretchr/testify/assert"
)

func TestGenreService_GetGenres(t *testing.T) {
	ctx := context.Background()
	genres := []models.Genre{
		{
			ID:   1,
			Uuid: utils.GetUUID(),
			Name: "test",
		},
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		doMock  func(mock *MockgenreRepo)
		want    []*gqlmodel.Genre
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			doMock: func(mock *MockgenreRepo) {
				mock.EXPECT().GetGenresFromDB(ctx).Return(genres, nil)
			},
			want: []*gqlmodel.Genre{
				{
					GenreUUID: genres[0].Uuid,
					GenreName: genres[0].Name,
				},
			},
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
			},
			doMock: func(mock *MockgenreRepo) {
				mock.EXPECT().GetGenresFromDB(ctx).Return(nil, assert.AnError)
			},
			wantErr: utils.ErrFetchData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GenreService{}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mock := NewMockgenreRepo(mockCtrl)
			gs.genreRepo = mock
			tt.doMock(mock)

			got, gotErr := gs.GetGenres(tt.args.ctx)

			assert.ObjectsAreEqualValues(got, tt.want)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}
