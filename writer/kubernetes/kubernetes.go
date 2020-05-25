package kubernetes

import (
	"context"
	"errors"
	"time"

	"github.com/bygui86/go-timescaledb/writer/logging"
)

func New() (*Server, error) {
	logging.Log.Info("Create new Kubernetes probe-server")

	cfg := loadConfig()
	cfgErr := cfg.checkConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	server := &Server{
		config: cfg,
	}
	server.newRouter()
	server.newHTTPServer()
	return server, nil
}

func (s *Server) Start() error {
	logging.Log.Info("Start Kubernetes probe-server")

	if s.httpServer != nil && !s.running {
		go func() {
			err := s.httpServer.ListenAndServe()
			if err != nil {
				logging.SugaredLog.Errorf("Kubernetes probe-server start failed: %s", err.Error())
			}
		}()
		s.running = true
		logging.SugaredLog.Infof("Kubernetes probe-server listen on port %d", s.config.restPort)
		return nil
	}

	return errors.New("Kubernetes probe-server start failed: HTTP server not initialized or HTTP server already running")
}

func (s *Server) Shutdown(timeout int) {
	logging.Log.Info("Shutdown Kubernetes probe-server")

	if s.httpServer != nil && s.running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logging.SugaredLog.Errorf("Kubernetes probe-server shutdown failed: %s", err.Error())
		}
		s.running = false
		return
	}

	logging.Log.Error("Kubernetes probe-server shutdown failed: HTTP server not initialized or HTTP server not running")
}
