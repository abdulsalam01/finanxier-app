package user

import (
	"context"

	"github.com/finanxier-app/internal/repository/base"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:generate mockgen -destination=user_mock.go -source=types.go -package=user
type User struct {
	ID           string `json:"id" db:"id" table:"products"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password_hash" db:"password_hash"`

	base.Metadata
	base.ExtraAttribute
}

type databaseResource interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type UsersRepo struct {
	database    databaseResource
	jwtSetupKey string
}
