package config

import (
	"context"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SchemaSeed(database *pgxpool.Pool) error {
	var (
		err    error
		runner seederRunner[interface{}]
	)

	err = runner.normalizeSeeders("./database/seeders")
	if err != nil {
		return err
	}

	return nil
}

func (s *seederRunner[T]) exec(ctx context.Context, database *pgxpool.Pool) error {
	var err error

	tx, err := database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	// Defer last function to run the connection commit or revert.
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			return
		}

		tx.Commit(ctx)
	}()

	return nil
}

func (s *seederRunner[T]) normalizeSeeders(path string) error {
	var (
		err error
	)

	err = filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			// Handle the error from filepath.WalkDir.
			return err
		}

		if d.IsDir() {
			// It's a directory, continue walking.
			return nil
		}

		raw, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		err = json.Unmarshal(raw, &s)
		if err != nil {
			// Handle the error from processing the JSON file.
			return err
		}

		return nil
	})

	// Check for errors from filepath.WalkDir.
	if err != nil {
		return err
	}

	return nil
}
