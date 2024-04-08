package product

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/finanxier-app/internal/entity"
	"go.uber.org/mock/gomock"
)

func Test_productUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockProductRepo := NewMockproductResource(ctrl)

	type args struct {
		ctx    context.Context
		params entity.Product
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Product
		wantErr bool
		mock    func()
	}{
		{
			name: "success - create",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockProductRepo.EXPECT().Create(ctx, gomock.Any()).Return(entity.Product{
					ID:    "cbdd9ad8-1125-47c1-98ce-6cc518e2070d",
					Name:  "Book",
					Price: 1000,
				}, nil)
			},
			wantErr: false,
			want: entity.Product{
				ID:    "cbdd9ad8-1125-47c1-98ce-6cc518e2070d",
				Name:  "Book",
				Price: 1000,
			},
		},
		{
			name: "failed - create",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockProductRepo.EXPECT().Create(ctx, gomock.Any()).Return(entity.Product{}, errors.New("error happened"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &productUsecase{
				productRepo: mockProductRepo,
			}

			tt.mock()

			got, err := c.Create(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("productUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
