package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/bsm/redislock"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_GenerateJWT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockUserUsecase := NewMockuserUsecase(ctrl)
	mockBaseInitializer := NewMockbaseAppInitializerResource(ctrl)

	r := httptest.NewRequest("GET", "/token-generator", nil)
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
				var r *redislock.Lock
				// Set up the expected calls to the mockProductUsecase.
				mockBaseInitializer.EXPECT().Lock(ctx, gomock.Any()).Return(r, nil)
				mockUserUsecase.EXPECT().GenerateValidJWT(ctx).Return("token", nil)
			},
			want:    "token",
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
		{
			name: "failed - get by params",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				var r *redislock.Lock
				// Set up the expected calls to the mockProductUsecase.
				mockBaseInitializer.EXPECT().Lock(ctx, gomock.Any()).Return(r, nil)
				mockUserUsecase.EXPECT().GenerateValidJWT(ctx).Return("", errors.New("error"))
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				userUsecase: mockUserUsecase,
				baseApp:     mockBaseInitializer,
			}

			tt.mock()

			got, err := h.GenerateJWT(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GenerateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GenerateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
