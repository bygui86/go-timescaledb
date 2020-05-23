package rest

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	config     *config
	router     *mux.Router
	httpServer *http.Server
	db         *sql.DB
	running    bool
}

type config struct {
	dbHost     string
	dbPort     int
	dbUsername string
	dbPassword string
	dbName     string
	dbSslMode  string
	restHost   string
	restPort   int
}
