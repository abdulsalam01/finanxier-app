package product

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity/base"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		ctx = r.Context()
		id  = chi.URLParam(r, "id")
	)

	if id == constant.DefaultString {
		return nil, errors.New("err")
	}

	return h.productUsecase.GetByID(ctx, id)
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		err   error
		ctx   = r.Context()
		query = r.URL.Query()

		limit   = query.Get("limit")
		offset  = query.Get("offset")
		iLimit  = constant.DefaultLimit
		iOffset = constant.DefaultOffset
	)

	// Validate queryParams.
	if limit != constant.DefaultString {
		iLimit, err = strconv.Atoi(limit)
		if err != nil {
			iLimit = constant.DefaultLimit
		}
	}
	if offset == constant.DefaultString {
		iOffset, err = strconv.Atoi(offset)
		if err != nil {
			iOffset = constant.DefaultOffset
		}
	}

	// Setup as base pagination results.
	params := base.PaginationRequest{
		Limit:  iLimit,
		Offset: iOffset,
	}
	return h.productUsecase.GetByParams(ctx, params)
}
