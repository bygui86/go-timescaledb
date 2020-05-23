package data_writer

import (
	"context"
	"database/sql"
	"time"
)

type Writer struct {
	config  *config
	ctx     context.Context
	db      *sql.DB
	ticker  *time.Ticker
	running bool
}

type config struct {
	interval time.Duration
}
