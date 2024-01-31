package tools

import (
	"context"
	"reflect"

	db "github.com/api-sekejap/pkg/database"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (s *seederRunner) CreateFactorySeeder(ctx context.Context, base db.DatabaseHelper) error {
	var (
		err        error
		dataParser []reflect.Value
	)

	dataRawParser := s.Data
	typeParser := reflect.TypeOf(s.getTypeInstance())

	for _, v := range dataRawParser {
		vParserRaw := reflect.ValueOf(v)
		if vParserRaw.Type() != typeParser {
			logrus.Warn("Failed parsing")
			continue
		}

		dataParser = append(dataParser, vParserRaw)
	}

	// Assume you have a method to create an instance of the seeder.
	instanceReflectType := reflect.TypeOf(s.getTypeInstance())
	instance := reflect.New(instanceReflectType).Interface()

	err = base.WithTx(ctx, func(tx pgx.Tx) error {
		createMethod := reflect.ValueOf(instance).MethodByName("Create")
		for _, v := range dataParser {
			results := createMethod.Call([]reflect.Value{reflect.ValueOf(ctx), v})
			if !results[1].IsNil() {
				return results[1].Interface().(error)
			}
		}

		return nil
	})

	return err
}
