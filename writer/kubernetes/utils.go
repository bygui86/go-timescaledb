package kubernetes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-timescaledb/writer/commons"
	"github.com/bygui86/go-timescaledb/writer/logging"
)

func (s *Server) newRouter() {
	logging.SugaredLog.Debugf("Setup new Kubernetes router")

	s.router = mux.NewRouter().StrictSlash(true)

	s.router.HandleFunc("/live", livenessHandler)
	s.router.HandleFunc("/ready", readinessHandler)
}

func (s *Server) newHTTPServer() {
	logging.SugaredLog.Debugf("Setup new Kubernetes HTTP server on port %d...", s.config.restPort)

	if s.config != nil {
		s.httpServer = &http.Server{
			Addr:    fmt.Sprintf(commons.HttpServerHostFormat, s.config.restHost, s.config.restPort),
			Handler: s.router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: commons.HttpServerWriteTimeoutDefault,
			ReadTimeout:  commons.HttpServerReadTimeoutDefault,
			IdleTimeout:  commons.HttpServerIdelTimeoutDefault,
		}
		return
	}

	logging.Log.Error("Kubernetes HTTP server creation failed: Kubernetes configurations not loaded")
}
