package feature

import (
	"github.com/api-sekejap/internal/entity"
	"github.com/api-sekejap/internal/repository/base"
)

func parserParams(params entity.Feature) Features {
	c := Features{
		Name:           params.Name,
		Description:    params.Description,
		Metadata:       base.Metadata(params.Metadata),
		ExtraAttribute: base.ExtraAttribute(params.ExtraAttribute),
	}

	if params.ID > 0 {
		c.ID = params.ID
	}

	return c
}
