package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	// Generic info.
	Driver string `json:"driver" yaml:"driver"`
	Port   string `json:"port" yaml:"port"`
	Host   string `json:"host" yaml:"host"`

	// Data info.
	Name     string                 `json:"name" yaml:"name"`
	Username string                 `json:"user" yaml:"user"`
	Password string                 `json:"password" yaml:"password"`
	Extras   map[string]interface{} `json:"extras" yaml:"extras"`
}

type DatabaseHelper struct {
	Database *pgxpool.Pool
}
