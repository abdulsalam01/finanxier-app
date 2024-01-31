package tools

import (
	"context"

	db "github.com/api-sekejap/pkg/database"
)

type seederRunner struct {
	Data []interface{} `json:"data"`
	Type string        `json:"type"`
}

// Seeder interface for all seeders.
type seederResources interface {
	Seed(ctx context.Context, data seederRunner, base db.DatabaseHelper) error
}

// Implement Seeder for all tables.
type ChannelSeeder struct{}
type UserSeeder struct{}
type FeatureSeeder struct{}
