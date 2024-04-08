package user

import (
	"context"
	"errors"

	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity"
	"github.com/golang-jwt/jwt/v4"
)

func (c *userUsecase) GenerateValidJWT(ctx context.Context) (string, error) {
	var (
		result string
		user   entity.User
		err    error
	)

	user, err = c.userRepo.GetFirstOne(ctx)
	if err != nil {
		return result, err
	}

	result, err = c.userRepo.JWTGenerator(ctx, user)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *userUsecase) GetCurrentUser(ctx context.Context, token any) (entity.User, error) {
	var (
		result entity.User
		err    error
	)

	// Parsing data from claims and get the column as map.
	username, ok := token.(jwt.MapClaims)[constant.UserColumnUsername]
	if !ok {
		return result, errors.New("failed parsing")
	}
	usernameStr, ok := username.(string)
	if !ok {
		return result, errors.New("failed parsing")
	}

	result, err = c.userRepo.GetCurrentUser(ctx, entity.User{Username: usernameStr})
	if err != nil {
		return result, err
	}

	return result, nil
}
