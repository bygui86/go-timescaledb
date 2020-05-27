package config

type Config struct {
	enableMonitoring bool
	enableTracing    bool
	tracingTech      string
	enableKubeProbes bool
	shutdownTimeout  int
}
