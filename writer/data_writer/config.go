package data_writer

import (
	"time"

	"github.com/bygui86/go-timescaledb/writer/logging"
	"github.com/bygui86/go-timescaledb/writer/utils"
)

const (
	intervalEnvVar = "WRITER_INTERVAL" // in seconds

	intervalEnvVarDefault = 1
)

func loadConfig() *config {
	logging.Log.Debug("Load WRITER configurations")
	return &config{
		interval: time.Duration(utils.GetIntEnv(intervalEnvVar, intervalEnvVarDefault)) * time.Second,
	}
}
