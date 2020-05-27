package rest

import (
	"github.com/bygui86/go-timescaledb/reader/logging"
	"github.com/bygui86/go-timescaledb/reader/utils"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostEnvVarDefault = ""
	restPortEnvVarDefault = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load REST configurations")
	return &config{
		restHost: utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		restPort: utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}
