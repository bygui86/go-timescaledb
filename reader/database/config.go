package database

import (
	"github.com/bygui86/go-timescaledb/reader/logging"
	"github.com/bygui86/go-timescaledb/reader/utils"
)

const (
	dbHostEnvVar     = "DB_HOST"
	dbPortEnvVar     = "DB_PORT"
	dbUsernameEnvVar = "DB_USERNAME"
	dbPasswordEnvVar = "DB_PASSWORD"
	dbNameEnvVar     = "DB_NAME"
	dbSslModeEnvVar  = "DB_SSL_MODE"

	dbHostEnvVarDefault     = "localhost"
	dbPortEnvVarDefault     = 5432
	dbUsernameEnvVarDefault = "username"
	dbPasswordEnvVarDefault = "password"
	dbNameEnvVarDefault     = "db"
	dbSslModeEnvVarDefault  = "disable"
)

func loadConfig() *config {
	logging.Log.Debug("Load REST configurations")
	return &config{
		dbHost:     utils.GetStringEnv(dbHostEnvVar, dbHostEnvVarDefault),
		dbPort:     utils.GetIntEnv(dbPortEnvVar, dbPortEnvVarDefault),
		dbUsername: utils.GetStringEnv(dbUsernameEnvVar, dbUsernameEnvVarDefault),
		dbPassword: utils.GetStringEnv(dbPasswordEnvVar, dbPasswordEnvVarDefault),
		dbName:     utils.GetStringEnv(dbNameEnvVar, dbNameEnvVarDefault),
		dbSslMode:  utils.GetStringEnv(dbSslModeEnvVar, dbSslModeEnvVarDefault),
	}
}
