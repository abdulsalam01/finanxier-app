package product

import (
	"github.com/finanxier-app/internal/entity"
	baseEnt "github.com/finanxier-app/internal/entity/base"
	"github.com/finanxier-app/internal/repository/base"
)

func parserParams(params entity.Product) Product {
	c := Product{
		ID:             params.ID,
		Name:           params.Name,
		Price:          params.Price,
		Metadata:       base.Metadata(params.Metadata),
		ExtraAttribute: base.ExtraAttribute(params.ExtraAttribute),
	}

	return c
}

func (p *Product) Normalize() entity.Product {
	return entity.Product{
		ID:             p.ID,
		Name:           p.Name,
		Price:          p.Price,
		Metadata:       baseEnt.Metadata(p.Metadata),
		ExtraAttribute: baseEnt.ExtraAttribute(p.ExtraAttribute),
	}
}
