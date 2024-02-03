package feature

import (
	"context"

	"github.com/api-sekejap/internal/entity"
	"github.com/api-sekejap/pkg/database"
	"github.com/wcamarao/pmx"
)

func (c *FeaturesRepo) Create(ctx context.Context, params entity.Feature) (int, error) {
	var (
		id  int
		err error
	)

	paramsParser := parserParams(params)
	_, err = pmx.Insert(ctx, c.database, &paramsParser)
	if err != nil {
		return id, database.WrapDuplicateKeyValueErr(err)
	}

	return id, nil
}
