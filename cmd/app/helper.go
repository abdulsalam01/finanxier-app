package app

import (
	"context"
	"fmt"
	"time"

	"github.com/api-sekejap/config"
	"github.com/api-sekejap/internal/constant"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Base initializer function to return base of application requirements.
// This function will contains:
// 1. Database.
// 2. Logging.
// 3. Redis.
// 4. ErrorWrapper.
func Initializer(ctx context.Context, config *config.Config) (BaseInitializer, error) {
	var (
		initializer BaseInitializer
		err         error
	)

	/*
	 * Configuration layer.
	 * Database section.
	 */
	dbUrlConnection := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	dbDurationMax, err := time.ParseDuration(config.Database.Extras[constant.DatabaseTimeout].(string))
	if err != nil {
		logrus.Fatalf("Unable to parse duration of connection config pool: %v\n", err)
		return initializer, err
	}

	dbConfig, err := pgxpool.ParseConfig(dbUrlConnection)
	if err != nil {
		logrus.Fatalf("Unable to parse connection config pool: %v\n", err)
		return initializer, err
	}

	// Setup base db connection.
	dbConfig.MaxConns = int32(config.Database.Extras[constant.DatabaseMaxConnection].(int))
	dbConfig.MinConns = int32(config.Database.Extras[constant.DatabaseMinConnection].(int))
	dbConfig.MaxConnLifetime = dbDurationMax
	dbConfig.ConnConfig.ConnectTimeout = dbDurationMax
	// Tie with database pool configuration.
	dbPool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		logrus.Fatalf("Unable to create connection pool: %v\n", err)
		return initializer, err
	}
	defer dbPool.Close()

	/*
	 * Configuration layer.
	 * Logger section.
	 */

	/*
	 * Configuration layer.
	 * Redis section.
	 */

	/*
	 * Configuration layer.
	 * Errorwraper section.
	 */

	// Initializer all here.
	initializer = BaseInitializer{
		Database: dbPool,
	}
	return initializer, nil
}
