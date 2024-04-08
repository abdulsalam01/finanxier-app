package product

import (
	"bytes"
	context "context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/bsm/redislock"
	"github.com/finanxier-app/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockProductUsecase := NewMockproductUsecase(ctrl)
	mockBaseInitializer := NewMockbaseAppInitializerResource(ctrl)

	// Create a new HTTP request with the method POST.
	r := httptest.NewRequest("POST", "/products", bytes.NewBufferString(`{"name":"test product","price":100}`))
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
			name: "success - create",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				var r *redislock.Lock
				// Set up the expected calls to the mockProductUsecase.
				mockBaseInitializer.EXPECT().Lock(ctx, gomock.Any()).Return(r, nil)
				mockBaseInitializer.EXPECT().Struct(gomock.Any()).Return(nil)
				mockProductUsecase.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(entity.Product{
						Name:  "test",
						Price: 100,
					}, nil)

			},
			want: entity.Product{
				Name:  "test",
				Price: 100,
			},
			wantErr: false,
		},
		{
			name: "failed - lock",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				// Set up the expected calls to the mockProductUsecase.
				mockBaseInitializer.EXPECT().Lock(ctx, gomock.Any()).Return(nil, errors.New("error"))

			},
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

			got, err := h.CreateProduct(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
