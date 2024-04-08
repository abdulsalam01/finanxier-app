package product

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/internal/entity/base"
	"go.uber.org/mock/gomock"
)

func Test_productUsecase_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockProductRepo := NewMockproductResource(ctrl)

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Product
		wantErr bool
		mock    func()
	}{
		{
			name: "success - get by id",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockProductRepo.EXPECT().GetByID(ctx, gomock.Any()).Return(entity.Product{
					ID:    "cbdd9ad8-1125-47c1-98ce-6cc518e2070d",
					Name:  "Book",
					Price: 1000,
				}, nil)
			},
			want: entity.Product{
				ID:    "cbdd9ad8-1125-47c1-98ce-6cc518e2070d",
				Name:  "Book",
				Price: 1000,
			},
			wantErr: false,
		},
		{
			name: "failed - get by id",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockProductRepo.EXPECT().GetByID(ctx, gomock.Any()).Return(entity.Product{}, errors.New("error happened"))
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

			got, err := c.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("productUsecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productUsecase.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_productUsecase_GetByParams(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockProductRepo := NewMockproductResource(ctrl)

	type args struct {
		ctx    context.Context
		params base.PaginationRequest
	}
	tests := []struct {
		name    string
		args    args
		want    entity.ProductBulkResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "success - get by params",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockProductRepo.EXPECT().Count(ctx).Return(2, nil)
				mockProductRepo.EXPECT().GetByParams(ctx, gomock.Any()).Return([]entity.Product{
					{
						ID:    "cbdd9ad8-1125-47c1-98ce-6cc518e2070d",
						Name:  "Book",
						Price: 1000,
					},
					{
						ID:    "6cc59ad8-1125-47c1-98ce-adc518e2070d",
						Name:  "Pencil",
						Price: 500,
					},
				}, nil)
			},
			want: entity.ProductBulkResponse{
				Product: []entity.Product{
					{
						ID:    "cbdd9ad8-1125-47c1-98ce-6cc518e2070d",
						Name:  "Book",
						Price: 1000,
					},
					{
						ID:    "6cc59ad8-1125-47c1-98ce-adc518e2070d",
						Name:  "Pencil",
						Price: 500,
					},
				},
				Total: 2,
			},
			wantErr: false,
		},
		{
			name: "failed - count",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockProductRepo.EXPECT().Count(ctx).Return(0, errors.New("error happened"))
			},
			wantErr: true,
		},
		{
			name: "failed - get by params",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockProductRepo.EXPECT().Count(ctx).Return(2, nil)
				mockProductRepo.EXPECT().GetByParams(ctx, gomock.Any()).Return([]entity.Product{}, errors.New("error happened"))
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

			got, err := c.GetByParams(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("productUsecase.GetByParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productUsecase.GetByParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
