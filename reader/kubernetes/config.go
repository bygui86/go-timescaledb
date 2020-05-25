package kubernetes

import (
	"errors"

	"github.com/bygui86/go-timescaledb/reader/logging"
	"github.com/bygui86/go-timescaledb/reader/utils"
)

const (
	restHostEnvVar = "KUBE_HOST"
	restPortEnvVar = "KUBE_PORT"

	restHostDefault = "localhost"
	restPortDefault = 9091
)

func loadConfig() *config {
	logging.Log.Debug("Load Kubernetes configurations")
	return &config{
		restHost: utils.GetStringEnv(restHostEnvVar, restHostDefault),
		restPort: utils.GetIntEnv(restPortEnvVar, restPortDefault),
	}
}

func (c *config) checkConfig() error {
	if c.restPort < 1024 {
		return errors.New("kubernetes port can't be less than 1024")
	}
	return nil
}
