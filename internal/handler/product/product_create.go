package product

import (
	"encoding/json"
	"net/http"

	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		err     error
		product entity.ProductRequest

		ctx = r.Context()
	)

	// Define locker to prevent duplicate request at the same time.
	lock, err := h.baseApp.Lock(ctx, constant.ProductUsecaseCreate)
	if err != nil {
		return nil, err
	}
	defer lock.Release(ctx)

	if err = json.NewDecoder(r.Body).Decode(&product); err != nil {
		return nil, err
	}
	if err = h.baseApp.Validate.Struct(product); err != nil {
		return nil, err
	}

	return h.productUsecase.Create(ctx, product.NormalizeRequest())
}
