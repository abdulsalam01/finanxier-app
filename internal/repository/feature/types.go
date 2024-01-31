package feature

import (
	"context"

	"github.com/api-sekejap/internal/repository/base"
	"github.com/jackc/pgx/v5/pgconn"
)

type Features struct {
	ID          int    `json:"id" db:"id" table:"features"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	Meta  base.Metadata
	Extra base.ExtraAttribute
}

type databaseResource interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type FeaturesRepo struct {
	database databaseResource
}
