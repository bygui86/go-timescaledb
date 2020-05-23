package tracing

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegerlogzap "github.com/uber/jaeger-client-go/log/zap"
	"github.com/uber/jaeger-lib/metrics"
	jaegerprom "github.com/uber/jaeger-lib/metrics/prometheus"

	"github.com/bygui86/go-timescaledb/reader/logging"
)

/*
	By default, the client sends traces via UDP to the agent at localhost:6831.
	Use JAEGER_AGENT_HOST and JAEGER_AGENT_PORT to send UDP traces to a different host:port.

	If JAEGER_ENDPOINT is set, the client sends traces to the endpoint via HTTP, making the JAEGER_AGENT_HOST and
	JAEGER_AGENT_PORT unused.

	If JAEGER_ENDPOINT is secured, HTTP basic authentication can be performed by setting the JAEGER_USER and
	JAEGER_PASSWORD environment variables.
*/

// Sample configuration for an easy start. Use constant sampling to sample every trace and enable LogSpan to log
// every span to stdout. No metrics are produced.
func InitSampleJaeger(serviceName string) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	stdLogger := jaegerlog.StdLogger
	nullMetricsFactory := metrics.NullFactory

	// Initialize tracing with a logger and a metrics factory
	closer, tracerErr := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(stdLogger),
		jaegercfg.Metrics(nullMetricsFactory),
	)
	if tracerErr != nil {
		return nil, tracerErr
	}

	logging.SugaredLog.Debugf("Jaeger global tracer registered: %t", opentracing.IsGlobalTracerRegistered())
	return closer, nil
}

// Sample configuration for testing. Use constant sampling to sample every trace and enable LogSpan to log every
// span via configured Logger. Use a Prometheus registerer to expose metrics.
func InitTestingJaeger(serviceName string) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	// Initialize tracing with a logger and a metrics factory
	closer, tracerErr := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(jaegerlogzap.NewLogger(logging.Log)),
		jaegercfg.Metrics(jaegerprom.New(jaegerprom.WithRegisterer(prometheus.DefaultRegisterer))),
	)
	if tracerErr != nil {
		return nil, tracerErr
	}

	logging.SugaredLog.Debugf("Jaeger global tracer registered: %t", opentracing.IsGlobalTracerRegistered())
	return closer, nil
}

// Recommended configuration for production.
func InitProductionJaeger(serviceName string) (io.Closer, error) {
	cfg := jaegercfg.Configuration{}

	// Initialize tracing with a logger and a metrics factory
	closer, tracerErr := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(jaegerlogzap.NewLogger(logging.Log)),
		jaegercfg.Metrics(jaegerprom.New(jaegerprom.WithRegisterer(prometheus.DefaultRegisterer))),
	)
	if tracerErr != nil {
		return nil, tracerErr
	}

	logging.SugaredLog.Debugf("Jaeger global tracer registered: %t", opentracing.IsGlobalTracerRegistered())
	return closer, nil
}
