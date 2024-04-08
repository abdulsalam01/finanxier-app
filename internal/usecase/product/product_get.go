package product

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/internal/entity/base"
)

func (c *productUsecase) GetByID(ctx context.Context, id string) (entity.Product, error) {
	var (
		result entity.Product
		err    error
	)

	result, err = c.productRepo.GetByID(ctx, id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *productUsecase) GetByParams(ctx context.Context, params base.PaginationRequest) (entity.ProductBulkResponse, error) {
	var (
		result entity.ProductBulkResponse
		total  int
		err    error
	)

	total, err = c.productRepo.Count(ctx)
	if err != nil {
		return result, err
	}

	product, err := c.productRepo.GetByParams(ctx, params)
	if err != nil {
		return result, err
	}

	return entity.ProductBulkResponse{
		Product: product,
		Total:   total,
	}, nil
}
