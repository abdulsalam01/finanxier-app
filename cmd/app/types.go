package app

import "github.com/jackc/pgx/v5/pgxpool"

type BaseInitializer struct {
	Database *pgxpool.Pool
}
