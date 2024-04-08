package product

import (
	"context"

	"github.com/finanxier-app/pkg/database"
)

func (c *ProductsRepo) Count(ctx context.Context) (int, error) {
	var (
		result int
		err    error
	)

	err = c.database.QueryRow(ctx, querySelectCount).Scan(&result)
	if err != nil {
		return result, database.WrapDuplicateKeyValueErr(err)
	}

	return result, nil
}
