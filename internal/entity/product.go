package entity

import (
	"github.com/finanxier-app/internal/entity/base"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`

	base.Metadata
	base.ExtraAttribute
}

// Spesific response.
type ProductBulkResponse struct {
	Product []Product `json:"products"`
	Total   int       `json:"total"`
}

// Spesific Handler request.
type ProductRequest struct {
	Name  string  `json:"name" validate:"required,min=4"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

func (p *ProductRequest) NormalizeRequest() Product {
	return Product{
		Name:  p.Name,
		Price: p.Price,
	}
}
