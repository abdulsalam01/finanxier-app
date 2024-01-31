package tools

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func SchemaMigrate(connUrl string, version uint) (bool, error) {
	var (
		err     error
		isExist bool
	)

	migrations, err := migrate.New("file://config/database/migrations", connUrl)
	if err != nil {
		return isExist, err
	}

	ver, _, err := migrations.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return isExist, err
	}
	if ver == version { // Indicating the migrations is exists.
		return true, nil
	}

	err = migrations.Migrate(version)
	if err != nil {
		return isExist, err
	}

	defer migrations.Close()
	return isExist, nil
}
