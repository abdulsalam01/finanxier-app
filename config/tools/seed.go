package tools

import (
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/api-sekejap/config"
	"github.com/api-sekejap/internal/constant"
	"github.com/api-sekejap/internal/entity"
	"github.com/api-sekejap/internal/repository/channel"
	"github.com/api-sekejap/internal/repository/feature"
	db "github.com/api-sekejap/pkg/database"
	"github.com/jackc/pgx/v5"
)

func SchemaSeed(ctx context.Context, base db.DatabaseHelper) error {
	var (
		err    error
		runner seederRunner
	)

	// Init map.
	runner = seederRunner{
		Data: make(map[string]interface{}),
		Type: make(map[string]string),
	}

	err = runner.normalizeSeeders(config.KeyPath, config.TypeParser)
	if err != nil {
		return err
	}
	err = runner.normalizeSeeders(config.SeederPath, config.DataParser)
	if err != nil {
		return err
	}

	// Executor runners.
	return runner.exec(ctx, []seederResources{
		ChannelSeeder{},
		FeatureSeeder{},
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

		// Use the built-in function to obtain filename.
		fileNormalize := filepath.Base(filePath)
		// Normalize the fileName without extension, as same as name of constants againts file.
		fileNormalizeWithoutExt := strings.TrimSuffix(fileNormalize, filepath.Ext(fileNormalize))

		switch typ {
		case 0: // Data params structure.
			key, instance := s.getTypeInstance(fileNormalizeWithoutExt)
			err = json.Unmarshal(raw, &instance)
			if err != nil {
				// Handle the error from processing the JSON file.
				return err
			}

			s.Data[key] = instance
		case 1: // Type params structure.
			instance := seederType{}
			err = json.Unmarshal(raw, &instance)
			if err != nil {
				// Handle the error from processing the JSON file.
				return err
			}

			s.Type[instance.Type] = instance.Type
		}
		return nil
	})

	// Check for errors from filepath.WalkDir.
	if err != nil {
		return err
	}

	return nil
}

func (s *seederRunner) getTypeInstance(fileName string) (string, any) {
	var (
		key      string
		instance any
	)

	types, ok := s.Type[fileName]
	if !ok {
		// Handle unknown type.
		key = constant.DefaultString
		instance = new(struct{})
	}

	switch types {
	case constant.ChannelsTable:
		key = constant.ChannelsTable
		instance = new([]entity.Channel)
	case constant.UsersTable:
		key = constant.UsersTable
		instance = new([]entity.User)
	case constant.FeaturesTable:
		key = constant.FeaturesTable
		instance = new([]entity.Feature)
	}

	return key, instance
}

// Each struct function seeder implementations.
func (us UserSeeder) Seed(ctx context.Context, data seederRunner, base db.DatabaseHelper) error {
	return nil
}
func (fs FeatureSeeder) Seed(ctx context.Context, data seederRunner, base db.DatabaseHelper) error {
	var (
		err        error
		dataParser []entity.Feature
	)

	dataRawParser := data.Data[constant.FeaturesTable]
	vParserRaw, ok := dataRawParser.(*[]entity.Feature)
	if !ok {
		return errors.New("Failed parsing")
	}

	vParser := *vParserRaw
	dataParser = append(dataParser, vParser...)

	instance := feature.New(base.Database)
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
func (cs ChannelSeeder) Seed(ctx context.Context, data seederRunner, base db.DatabaseHelper) error {
	var (
		err        error
		dataParser []entity.Channel
	)

	dataRawParser := data.Data[constant.ChannelsTable]
	vParserRaw, ok := dataRawParser.(*[]entity.Channel)
	if !ok {
		return errors.New("Failed parsing")
	}

	vParser := *vParserRaw
	dataParser = append(dataParser, vParser...)

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