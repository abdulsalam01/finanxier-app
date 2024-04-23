package grpc

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/internal/entity/base"
	"github.com/finanxier-app/pkg/redis"
	pb "github.com/finanxier-app/proto/gen"
)

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

	pb.UnimplementedProductServiceServer
}
