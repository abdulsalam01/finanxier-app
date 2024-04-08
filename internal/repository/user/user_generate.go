package user

import (
	"context"
	"time"

	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity"
	"github.com/golang-jwt/jwt/v4"
)

func (c *UsersRepo) JWTGenerator(ctx context.Context, user entity.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims.
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims.
	claims["authorized"] = true
	claims["user"] = user.Username
	claims["exp"] = time.Now().Add(constant.SessionTTL).Unix()

	// Generate encoded token.
	tokenString, err := token.SignedString([]byte(c.jwtSetupKey))
	if err != nil {
		return constant.DefaultString, err
	}

	return tokenString, nil
}
