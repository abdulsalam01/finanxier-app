package user

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity"
	"github.com/golang-jwt/jwt/v4"
	gomock "go.uber.org/mock/gomock"
)

func Test_userUsecase_GenerateValidJWT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockUserRepo := NewMockuserResource(ctrl)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mock    func()
	}{
		{
			name: "success - generate jwt",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				user := entity.User{
					Username: "admin",
				}
				mockUserRepo.EXPECT().GetFirstOne(ctx).Return(user, nil)
				mockUserRepo.EXPECT().JWTGenerator(ctx, user).Return("token", nil)
			},
			want:    "token",
			wantErr: false,
		},
		{
			name: "failed - get first one",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				mockUserRepo.EXPECT().GetFirstOne(ctx).Return(entity.User{}, errors.New("error happened"))
			},
			wantErr: true,
		},
		{
			name: "failed - generate token",
			args: args{
				ctx: ctx,
			},
			mock: func() {
				user := entity.User{
					Username: "admin",
				}
				mockUserRepo.EXPECT().GetFirstOne(ctx).Return(user, nil)
				mockUserRepo.EXPECT().JWTGenerator(ctx, user).Return("", errors.New("error happened"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userUsecase{
				userRepo: mockUserRepo,
			}

			tt.mock()

			got, err := c.GenerateValidJWT(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GenerateValidJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userUsecase.GenerateValidJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_GetCurrentUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockUserRepo := NewMockuserResource(ctrl)
	claims := jwt.MapClaims{
		"sub":   "1234567890",
		"user":  "admin",           // A custom claim to include the user's name
		"admin": true,              // A custom claim to specify if the user is an admin
		"iat":   time.Now().Unix(), // Issued At: When the token was issued
		"exp":   time.Now().Add(constant.SessionTTL).Unix(),
	}

	type args struct {
		ctx   context.Context
		token any
	}
	tests := []struct {
		name    string
		args    args
		want    entity.User
		wantErr bool
		mock    func()
	}{
		{
			name: "success - get current user",
			args: args{
				ctx:   ctx,
				token: claims,
			},
			mock: func() {
				mockUserRepo.EXPECT().GetCurrentUser(ctx, gomock.Any()).Return(entity.User{
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
			name: "failed - token empty",
			args: args{
				ctx:   ctx,
				token: jwt.MapClaims{},
			},
			mock: func() {
			},
			wantErr: true,
		},
		{
			name: "failed - type data not string",
			args: args{
				ctx: ctx,
				token: jwt.MapClaims{
					"user": 1,
				},
			},
			mock: func() {
			},
			wantErr: true,
		},
		{
			name: "failed - get current user",
			args: args{
				ctx:   ctx,
				token: claims,
			},
			mock: func() {
				mockUserRepo.EXPECT().GetCurrentUser(ctx, gomock.Any()).Return(entity.User{}, errors.New("error happened"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &userUsecase{
				userRepo: mockUserRepo,
			}

			tt.mock()

			got, err := c.GetCurrentUser(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.GetCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.GetCurrentUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
