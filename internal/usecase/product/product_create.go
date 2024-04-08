package product

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/google/uuid"
)

func (c *productUsecase) Create(ctx context.Context, params entity.Product) (entity.Product, error) {
	var (
		product entity.Product
		err     error
	)

	// Set new ID.
	params.ID = uuid.NewString()
	product, err = c.productRepo.Create(ctx, params)
	if err != nil {
		return product, err
	}

	return product, nil
}
