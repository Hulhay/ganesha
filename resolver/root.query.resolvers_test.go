package resolver

import (
	"context"
	"testing"

	"github.com/hulhay/ganesha/graph/model"
	"github.com/magiconair/properties/assert"
)

func Test_queryResolver_Health(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *model.GeneralResponse
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: &model.GeneralResponse{
				Message:   "Ganesha is running",
				IsSuccess: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{}
			got, gotErr := r.Query().Health(tt.args.ctx)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}
