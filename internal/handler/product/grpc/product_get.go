package grpc

import (
	"context"

	"github.com/finanxier-app/internal/entity/base"
	pb "github.com/finanxier-app/proto/gen"
)

func (h *Handler) GetByParams(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	var (
		res *pb.ProductResponse = &pb.ProductResponse{}
		err error
	)

	// Normalize input.
	args := base.PaginationRequest{
		Limit:  int(*req.Limit),
		Offset: int(*req.Offset),
	}

	products, err := h.productUsecase.GetByParams(ctx, args)
	if err != nil {
		return res, err
	}

	// Normalize response.
	for _, v := range products.Product {
		res.Products = append(res.Products, &pb.Product{
			Id:    v.ID,
			Name:  v.Name,
			Price: float32(v.Price),
			Meta: &pb.Extra{
				CreatedBy: int32(v.CreatedBy),
				UpdatedBy: int64(v.UpdatedBy),
			},
		})
	}

	return res, nil
}
