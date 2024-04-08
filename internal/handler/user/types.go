package user

import (
	"context"

	"github.com/finanxier-app/cmd/app"
	"github.com/finanxier-app/internal/entity"
)

type userUsecase interface {
	GetCurrentUser(ctx context.Context, token any) (entity.User, error)
	GenerateValidJWT(ctx context.Context) (string, error)
}

type Handler struct {
	userUsecase userUsecase
	baseApp     app.BaseAppInitializer
}
