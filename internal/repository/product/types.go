package product

import (
	"context"

	"github.com/finanxier-app/internal/repository/base"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:generate mockgen -destination=product_mock.go -source=types.go -package=product
type Product struct {
	ID    string  `json:"id" db:"id" table:"products"`
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`

	base.Metadata
	base.ExtraAttribute
}

type databaseResource interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type ProductsRepo struct {
	database databaseResource
}
