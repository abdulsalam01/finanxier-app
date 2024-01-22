package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/api-sekejap/internal/constant"
	"github.com/api-sekejap/internal/entity"
	"github.com/api-sekejap/internal/repository/channel"
	db "github.com/api-sekejap/pkg/database"
	"github.com/jackc/pgx/v5"
)

func SchemaSeed(ctx context.Context, base db.DatabaseHelper) error {
	var (
		err    error
		runner seederRunner
	)

	err = runner.normalizeSeeders(keyPath, typeParser)
	if err != nil {
		return err
	}
	err = runner.normalizeSeeders(keyPath, dataParser)
	if err != nil {
		return err
	}

	// Executor runners.
	return runner.exec(ctx, []seederResources{
		ChannelSeeder{},
		UserSeeder{},
	}, base)
}

// Executor functions.
func (s *seederRunner) exec(ctx context.Context, seeder []seederResources, base db.DatabaseHelper) error {
	var err error

	for _, v := range seeder {
		err = v.Seed(ctx, *s, base)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *seederRunner) normalizeSeeders(path string, typ int) error {
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

		switch typ {
		case 0: // Data params structure.
			instance := s.getTypeInstance()
			err = json.Unmarshal(raw, &instance)
			if err != nil {
				// Handle the error from processing the JSON file.
				return err
			}

			s.Data = instance
			fmt.Printf("%+v", s.Data)
		case 1: // Type params structure.
			err = json.Unmarshal(raw, &s)
			if err != nil {
				// Handle the error from processing the JSON file.
				return err
			}
		}
		return nil
	})

	// Check for errors from filepath.WalkDir.
	if err != nil {
		return err
	}

	return nil
}

func (s *seederRunner) getTypeInstance() any {
	var instance any

	switch s.Type {
	case constant.ChannelsTable:
		instance = new(entity.Channel)
	case constant.UsersTable:
		instance = new(entity.User)
	case constant.FeaturesTable:
		instance = new(entity.Feature)
	default:
		// Handle unknown type.
		instance = new(struct{})
	}

	return instance
}

// Each struct function seeder implementations.
func (us UserSeeder) Seed(ctx context.Context, data seederRunner, base db.DatabaseHelper) error {
	return nil
}
func (cs ChannelSeeder) Seed(ctx context.Context, data seederRunner, base db.DatabaseHelper) error {
	var (
		err error
	)

	dataRawParser := data.Data
	dataParser, ok := dataRawParser.([]entity.Channel)
	if !ok {
		return errors.New("Failed parsing")
	}

	instance := channel.New(base.Database)
	err = base.WithTx(ctx, func(tx pgx.Tx) error {
		for _, v := range dataParser {
			if _, err = instance.Create(ctx, v); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
