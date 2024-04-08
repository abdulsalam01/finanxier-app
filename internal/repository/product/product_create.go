package product

import (
	"context"

	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/pkg/database"
)

func (c *ProductsRepo) Create(ctx context.Context, params entity.Product) (entity.Product, error) {
	var (
		result entity.Product
		id     string
		err    error
	)

	paramsParser := parserParams(params)
	err = c.database.
		QueryRow(
			ctx,
			queryInsertWithReturning,
			&paramsParser.ID, &paramsParser.Name, &paramsParser.Price).
		Scan(&id)
	if err != nil {
		return result, database.WrapDuplicateKeyValueErr(err)
	}

	// Re-set the success data.
	result = entity.Product{
		ID:    id,
		Name:  params.Name,
		Price: params.Price,
	}
	return result, nil
}
