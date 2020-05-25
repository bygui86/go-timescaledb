package config

import (
	"github.com/bygui86/go-timescaledb/reader/logging"
	"github.com/bygui86/go-timescaledb/reader/utils"
)

const (
	enableMonitoringEnvVar = "ENABLE_MONITORING"  // bool
	enableTracingEnvVar    = "ENABLE_TRACING"     // bool
	tracingTechEnvVar      = "TRACING_TECH"       //  available values: jaeger, zipkin
	enableKubeProbesEnvVar = "ENABLE_KUBE_PROBES" // bool

	enableMonitoringDefault = true
	enableTracingDefault    = true
	tracingTechDefault      = TracingTechJaeger
	enableKubeProbesDefault = true
)

func LoadConfig() *Config {
	logging.Log.Info("Load global configurations")

	tracingTech := utils.GetStringEnv(tracingTechEnvVar, tracingTechDefault)
	if tracingTech != TracingTechJaeger && tracingTech != TracingTechZipkin {
		logging.SugaredLog.Warnf("Tracing technology %s not supported, fallback to %s",
			tracingTech, TracingTechJaeger)
		tracingTech = TracingTechJaeger
	}

	return &Config{
		enableMonitoring: utils.GetBoolEnv(enableMonitoringEnvVar, enableMonitoringDefault),
		enableTracing:    utils.GetBoolEnv(enableTracingEnvVar, enableTracingDefault),
		tracingTech:      tracingTech,
		enableKubeProbes: utils.GetBoolEnv(enableKubeProbesEnvVar, enableKubeProbesDefault),
	}
}
