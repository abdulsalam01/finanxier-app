package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity/base"
	"github.com/golang-jwt/jwt/v4"
)

// JWTAuthMiddleware is a middleware for validating JWT tokens.
func JWTAuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				err          error
				jsonResponse []byte
			)

			// Set the Content-Type header to indicate JSON response.
			w.Header().Set("Content-Type", "application/json")
			mapData := base.Response[interface{}]{
				Data:    nil,
				Message: "Successfully executed",
				Success: false,
			}

			authHeader := r.Header.Get(constant.HTTPHeaderAuthorization)
			if authHeader == "" {
				mapData.Message = "Authorization header is missing"
				// Safe operation, without breaking changes or flow.
				jsonResponse, err = json.Marshal(mapData) //nolint:all

				// Write the JSON response.
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(jsonResponse)
				return
			}

			// Split the token type and the token itself.
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 || bearerToken[0] != constant.HTTPHeaderBearer {
				mapData.Message = "Invalid token format"
				// Safe operation, without breaking changes or flow.
				jsonResponse, err = json.Marshal(mapData) //nolint:all

				// Write the JSON response.
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(jsonResponse)
				return
			}

			// Parse the token.
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				mapData.Message = "Invalid token"
				// Safe operation, without breaking changes or flow.
				jsonResponse, err = json.Marshal(mapData) //nolint:all

				// Write the JSON response.
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(jsonResponse)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				// Write the JSON response.
				w.WriteHeader(http.StatusUnauthorized)

				return
			}

			// Token is valid, pass the request to the next middleware or handler.
			ctx := context.WithValue(r.Context(), constant.UserContext, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
