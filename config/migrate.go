package config

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func SchemaMigrate(connUrl string, version uint) error {
	var err error

	migations, err := migrate.New("file://config/database/migrations", connUrl)
	if err != nil {
		return err
	}

	err = migations.Migrate(version)
	if err != nil {
		return err
	}

	defer migations.Close()
	return nil
}
