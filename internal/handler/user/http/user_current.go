package http

import (
	"errors"
	"net/http"

	"github.com/finanxier-app/internal/constant"
)

func (h *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		ctx = r.Context()
	)

	user := ctx.Value(constant.UserContext)
	if user == nil {
		return nil, errors.New("user unauthorized")
	}

	return h.userUsecase.GetCurrentUser(ctx, user)
}
