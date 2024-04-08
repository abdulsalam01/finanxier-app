package product

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/internal/entity/base"
)

//go:generate mockgen -destination=product_mock.go -source=types.go -package=product
type productResource interface {
	// Get area.
	GetByID(ctx context.Context, id string) (entity.Product, error)
	GetByParams(ctx context.Context, params base.PaginationRequest) ([]entity.Product, error)

	// Calcute area.
	Count(ctx context.Context) (int, error)

	// Management area.
	Create(ctx context.Context, params entity.Product) (entity.Product, error)
}

type productUsecase struct {
	productRepo productResource
}
