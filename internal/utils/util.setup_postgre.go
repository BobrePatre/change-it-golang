package utils

import (
	"time"

	"change-it/internal/config"
	"change-it/internal/constants"
	"change-it/internal/datasources/drivers"
	"github.com/jmoiron/sqlx"
)

func SetupPostgresConnection() (*sqlx.DB, error) {
	var dsn string
	switch config.AppConfig.Environment {
	case constants.EnvironmentDevelopment:
		dsn = config.AppConfig.DBPostgreDsn
	case constants.EnvironmentProduction:
		dsn = config.AppConfig.DBPostgreURL
	}

	databaseConfig := drivers.SQLXConfig{
		DriverName:     config.AppConfig.DBPostgreDriver,
		DataSourceName: dsn,
		MaxOpenConns:   100,
		MaxIdleConns:   10,
		MaxLifetime:    15 * time.Minute,
	}

	conn, err := databaseConfig.InitializeSQLXDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
