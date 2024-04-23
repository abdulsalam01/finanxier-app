package http

import (
	context "context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/finanxier-app/internal/entity"
	"github.com/go-chi/chi/v5"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockProductUsecase := NewMockproductUsecase(ctrl)
	mockBaseInitializer := NewMockbaseAppInitializerResource(ctrl)

	r := httptest.NewRequest("GET", "/products/55506c8e-f046-4517-8c2e-b50b3c7a8f1a", nil)
	r = r.WithContext(context.WithValue(ctx, chi.RouteCtxKey, &chi.Context{
		URLParams: chi.RouteParams{
			Keys:   []string{"id"},
			Values: []string{"55506c8e-f046-4517-8c2e-b50b3c7a8f1a"},
		},
	}))
	// Create a ResponseRecorder to record the response.
	w := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
		mock    func()
	}{
		{
			name: "success - by id",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				mockProductUsecase.EXPECT().GetByID(r.Context(), gomock.Any()).Return(entity.Product{
					ID:   "55506c8e-f046-4517-8c2e-b50b3c7a8f1a",
					Name: "book",
				}, nil)
			},
			wantErr: false,
			want: entity.Product{
				ID:   "55506c8e-f046-4517-8c2e-b50b3c7a8f1a",
				Name: "book",
			},
		},
		{
			name: "failed - by id",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				mockProductUsecase.EXPECT().GetByID(r.Context(), gomock.Any()).Return(entity.Product{}, errors.New("error"))
			},
			want:    entity.Product{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				productUsecase: mockProductUsecase,
				baseApp:        mockBaseInitializer,
			}

			tt.mock()

			got, err := h.GetProduct(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockProductUsecase := NewMockproductUsecase(ctrl)
	mockBaseInitializer := NewMockbaseAppInitializerResource(ctrl)

	r := httptest.NewRequest("GET", "/products", nil)
	// Create a ResponseRecorder to record the response.
	w := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
		mock    func()
	}{
		{
			name: "success - get by params",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				mockProductUsecase.EXPECT().GetByParams(ctx, gomock.Any()).Return(entity.ProductBulkResponse{
					Product: []entity.Product{
						{
							Name:  "book",
							Price: 100,
						},
						{
							Name:  "pencil",
							Price: 200,
						},
					},
					Total: 2,
				}, nil)
			},
			want: entity.ProductBulkResponse{
				Product: []entity.Product{
					{
						Name:  "book",
						Price: 100,
					},
					{
						Name:  "pencil",
						Price: 200,
					},
				},
				Total: 2,
			},
			wantErr: false,
		},
		{
			name: "failed - get by params",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				mockProductUsecase.EXPECT().GetByParams(ctx, gomock.Any()).Return(entity.ProductBulkResponse{}, errors.New("error"))
			},
			want:    entity.ProductBulkResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				productUsecase: mockProductUsecase,
				baseApp:        mockBaseInitializer,
			}

			tt.mock()

			got, err := h.GetProducts(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GetProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
