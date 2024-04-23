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
		res.Products = append(res.Products, v.Normalize())
	}

	return res, nil
}

func (h *Handler) GetByID(ctx context.Context, req *pb.ProductParams) (*pb.Product, error) {
	var (
		res *pb.Product
		err error
	)

	// Normalize input.
	args := req.Id

	product, err := h.productUsecase.GetByID(ctx, args)
	if err != nil {
		return res, err
	}

	// Normalize response.
	res = product.Normalize()
	return res, nil
}
