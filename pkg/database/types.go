package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseHelper struct {
	Database *pgxpool.Pool
}
