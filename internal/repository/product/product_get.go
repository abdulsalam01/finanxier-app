package product

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/internal/entity/base"
	"github.com/finanxier-app/pkg/database"
)

func (c *ProductsRepo) GetByID(ctx context.Context, id string) (entity.Product, error) {
	var (
		result  entity.Product
		product Product
		err     error
	)

	err = c.database.QueryRow(ctx, querySelectByID, id).
		Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.IsActive,
			&product.CreatedBy,
			&product.UpdatedBy,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
	if err != nil {
		return result, database.WrapDuplicateKeyValueErr(err)
	}

	return product.Normalize(), nil
}

func (c *ProductsRepo) GetByParams(ctx context.Context, params base.PaginationRequest) ([]entity.Product, error) {
	var (
		result []entity.Product
		err    error
	)

	rows, err := c.database.Query(ctx, querySelectByParams, params.Limit, params.Offset)
	if err != nil {
		return result, database.WrapDuplicateKeyValueErr(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.IsActive,
			&product.CreatedBy,
			&product.UpdatedBy,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return result, err
		}

		result = append(result, product.Normalize())
	}

	return result, nil
}
