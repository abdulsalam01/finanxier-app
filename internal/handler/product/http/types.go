package http

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/internal/entity/base"
	"github.com/finanxier-app/pkg/redis"
)

//go:generate mockgen -destination=product_mock.go -source=types.go -package=http
type productUsecase interface {
	Create(ctx context.Context, params entity.Product) (entity.Product, error)
	GetByID(ctx context.Context, id string) (entity.Product, error)
	GetByParams(ctx context.Context, params base.PaginationRequest) (entity.ProductBulkResponse, error)
}

type baseAppInitializerResource interface {
	Lock(ctx context.Context, key string) (redis.RedisLockResource, error)
	Struct(s interface{}) error
}

type Handler struct {
	productUsecase productUsecase
	baseApp        baseAppInitializerResource
}
