package user

import (
	"net/http"

	"github.com/finanxier-app/internal/constant"
)

func (h *Handler) GenerateJWT(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var (
		ctx = r.Context()
	)

	// Define locker to prevent duplicate request at the same time.
	lock, err := h.baseApp.Lock(ctx, constant.UserUsecaseGenerate)
	if err != nil {
		return nil, err
	}
	defer lock.Release(ctx)

	return h.userUsecase.GenerateValidJWT(ctx)
}
