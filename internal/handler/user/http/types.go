package http

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/pkg/redis"
)

//go:generate mockgen -destination=user_mock.go -source=types.go -package=http

type userUsecase interface {
	GetCurrentUser(ctx context.Context, token any) (entity.User, error)
	GenerateValidJWT(ctx context.Context) (string, error)
}
type baseAppInitializerResource interface {
	Lock(ctx context.Context, key string) (redis.RedisLockResource, error)
	Struct(s interface{}) error
}

type Handler struct {
	userUsecase userUsecase
	baseApp     baseAppInitializerResource
}
