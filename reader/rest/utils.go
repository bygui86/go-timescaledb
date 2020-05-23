package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-timescaledb/reader/commons"
	"github.com/bygui86/go-timescaledb/reader/logging"
)

const (
	// urls
	rootEndpoint = "/conditions"

	contentTypeHeaderKey       = "Content-Type"
	contentTypeApplicationJson = "application/json"
)

// SERVER

func (s *Server) setupRouter() {
	logging.Log.Debug("Create new router")

	s.router = mux.NewRouter().StrictSlash(true)

	s.router.HandleFunc(rootEndpoint, s.getConditions).Methods(http.MethodGet)
}

func (s *Server) setupHTTPServer() {
	logging.SugaredLog.Debugf("Create new HTTP server on port %d", s.config.restPort)

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

	logging.Log.Error("HTTP server creation failed: REST server configurations not loaded")
}

// HANDLERS

func sendJsonResponse(writer http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	writer.Header().Set(contentTypeHeaderKey, contentTypeApplicationJson)
	writer.WriteHeader(code)
	_, err := writer.Write(response)
	if err != nil {
		logging.SugaredLog.Errorf("Error sending JSON response: %s", err.Error())
	}
}

func sendErrorResponse(writer http.ResponseWriter, code int, message string) {
	sendJsonResponse(writer, code, map[string]string{"error": message})
}
