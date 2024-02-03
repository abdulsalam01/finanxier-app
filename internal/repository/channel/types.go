package channel

import (
	"context"

	"github.com/api-sekejap/internal/repository/base"
	"github.com/jackc/pgx/v5/pgconn"
)

type Channels struct {
	ID          int    `json:"id" db:"id" table:"channels"`
	PackageID   int    `json:"package_id" db:"package_id"`
	Name        string `json:"name" db:"name"`
	Link        string `json:"link" db:"link"`
	AssetUrl    string `json:"asset_url" db:"asset_url"`
	Description string `json:"description" db:"description"`

	base.Metadata
	base.ExtraAttribute
}

type databaseResource interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type ChannelsRepo struct {
	database databaseResource
}
