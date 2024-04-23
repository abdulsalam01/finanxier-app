package entity

import (
	"github.com/finanxier-app/internal/entity/base"
	pb "github.com/finanxier-app/proto/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (p *Product) Normalize() *pb.Product {
	return &pb.Product{
		Id:    p.ID,
		Name:  p.Name,
		Price: float32(p.Price),
		Meta: &pb.Extra{
			CreatedBy: int32(p.CreatedBy),
			UpdatedBy: int64(p.UpdatedBy),
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt),
		},
	}
}
