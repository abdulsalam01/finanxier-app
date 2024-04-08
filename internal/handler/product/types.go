package product

import (
	"context"

	"github.com/finanxier-app/cmd/app"
	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/internal/entity/base"
)

type productUsecase interface {
	Create(ctx context.Context, params entity.Product) (entity.Product, error)
	GetByID(ctx context.Context, id string) (entity.Product, error)
	GetByParams(ctx context.Context, params base.PaginationRequest) (entity.ProductBulkResponse, error)
}

type Handler struct {
	productUsecase productUsecase
	baseApp        app.BaseAppInitializer
}
