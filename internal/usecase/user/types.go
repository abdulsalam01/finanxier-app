package user

import (
	"context"

	"github.com/finanxier-app/internal/entity"
)

type userResource interface {
	// Get area.
	GetByID(ctx context.Context, id string) (entity.User, error)
	GetFirstOne(ctx context.Context) (entity.User, error)
	GetCurrentUser(ctx context.Context, token entity.User) (entity.User, error)

	// Management area.
	JWTGenerator(ctx context.Context, user entity.User) (string, error)
}

type userUsecase struct {
	userRepo userResource
}
