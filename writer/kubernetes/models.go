package kubernetes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	config     *config
	router     *mux.Router
	httpServer *http.Server
	running    bool
}

type config struct {
	restHost string
	restPort int
}

type Liveness struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type Readiness struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
