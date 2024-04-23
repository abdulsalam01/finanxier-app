package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity"
	"github.com/golang-jwt/jwt/v4"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_GetCurrentUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	claims := jwt.MapClaims{
		"sub":   "1234567890",
		"user":  "admin",           // A custom claim to include the user's name
		"admin": true,              // A custom claim to specify if the user is an admin
		"iat":   time.Now().Unix(), // Issued At: When the token was issued
		"exp":   time.Now().Add(constant.SessionTTL).Unix(),
	}
	ctx = context.WithValue(ctx, constant.UserContext, claims)
	mockUserUsecase := NewMockuserUsecase(ctrl)
	mockBaseInitializer := NewMockbaseAppInitializerResource(ctrl)

	r := httptest.NewRequest("GET", "/users/current", nil)
	r = r.WithContext(ctx)
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
			name: "success - current user",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				mockUserUsecase.EXPECT().GetCurrentUser(ctx, gomock.Any()).Return(entity.User{
					ID:       "cbdd9ad8-1125-47c1-98ce-6cc518e2070d",
					Username: "admin",
				}, nil)
			},
			want: entity.User{
				ID:       "cbdd9ad8-1125-47c1-98ce-6cc518e2070d",
				Username: "admin",
			},
			wantErr: false,
		},
		{
			name: "failed - current user",
			args: args{
				w: w,
				r: r,
			},
			mock: func() {
				mockUserUsecase.EXPECT().GetCurrentUser(ctx, gomock.Any()).Return(entity.User{}, errors.New("error"))
			},
			want:    entity.User{},
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

			got, err := h.GetCurrentUser(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GetCurrentUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
