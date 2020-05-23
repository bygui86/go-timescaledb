package main

import (
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/openzipkin/zipkin-go/reporter"

	"github.com/bygui86/go-timescaledb/reader/config"
	"github.com/bygui86/go-timescaledb/reader/logging"
	"github.com/bygui86/go-timescaledb/reader/monitoring"
	"github.com/bygui86/go-timescaledb/reader/rest"
	"github.com/bygui86/go-timescaledb/reader/tracing"
)

const (
	serviceName = "http-server"

	zipkinHost = "localhost"
	zipkinPort = 9411
)

var (
	monitoringServer *monitoring.Server
	jaegerCloser     io.Closer
	zipkinReporter   reporter.Reporter
	restServer       *rest.Server
)

func main() {
	initLogging()

	logging.SugaredLog.Infof("Start %s", serviceName)

	cfg := loadConfig()

	if cfg.GetEnableMonitoring() {
		monitoringServer = startMonitoringServer()
	}

	if cfg.GetEnableTracing() {
		switch cfg.GetTracingTech() {
		case config.TracingTechJaeger:
			jaegerCloser = initJaegerTracer()
		case config.TracingTechZipkin:
			zipkinReporter = initZipkinTracer()
		}
	}

	restServer = startRestServer()

	logging.SugaredLog.Infof("%s up and running", serviceName)

	startSysCallChannel()

	shutdownAndWait(1)
}

func initLogging() {
	err := logging.InitGlobalLogger()
	if err != nil {
		logging.SugaredLog.Errorf("Logging setup failed: %s", err.Error())
		os.Exit(501)
	}
}

func loadConfig() *config.Config {
	logging.Log.Debug("Load configurations")
	return config.LoadConfig()
}

func startMonitoringServer() *monitoring.Server {
	logging.Log.Debug("Start monitoring")
	server := monitoring.New()
	logging.Log.Debug("Monitoring server successfully created")

	server.Start()
	logging.Log.Debug("Monitoring successfully started")

	return server
}

func initJaegerTracer() io.Closer {
	logging.Log.Debug("Init Jaeger tracer")
	closer, err := tracing.InitTestingJaeger(serviceName)
	if err != nil {
		logging.SugaredLog.Errorf("Jaeger tracer setup failed: %s", err.Error())
		os.Exit(501)
	}
	return closer
}

func initZipkinTracer() reporter.Reporter {
	logging.Log.Debug("Init Zipkin tracer")
	zReporter, err := tracing.InitTestingZipkin(serviceName, zipkinHost, zipkinPort)
	if err != nil {
		logging.SugaredLog.Errorf("Zipkin tracer setup failed: %s", err.Error())
		os.Exit(501)
	}
	return zReporter
}

func startRestServer() *rest.Server {
	logging.Log.Debug("Start REST server")

	server, newErr := rest.New(true)
	if newErr != nil {
		logging.SugaredLog.Errorf("REST server creation failed: %s", newErr.Error())
		os.Exit(501)
	}
	logging.Log.Debug("REST server successfully created")

	startErr := server.Start()
	if startErr != nil {
		logging.SugaredLog.Errorf("REST server start failed: %s", startErr.Error())
		os.Exit(502)
	}
	logging.Log.Debug("REST server successfully started")

	monitoring.RegisterCustomMetrics()

	return server
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if restServer != nil {
		restServer.Shutdown(timeout)
	}

	if jaegerCloser != nil {
		err := jaegerCloser.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Jaeger tracer closure failed: %s", err.Error())
		}
	}

	if zipkinReporter != nil {
		err := zipkinReporter.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Zipkin tracer closure failed: %s", err.Error())
		}
	}

	if monitoringServer != nil {
		monitoringServer.Shutdown(timeout)
	}

	time.Sleep(time.Duration(timeout+1) * time.Second)
}
