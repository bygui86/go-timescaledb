package data_writer

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bygui86/go-timescaledb/writer/database"
	"github.com/bygui86/go-timescaledb/writer/logging"
)

func New(ctx context.Context, enableTracing bool) (*Writer, error) {
	logging.Log.Info("Create new WRITER")

	cfg := loadConfig()

	var db *sql.DB
	var dbErr error
	if enableTracing {
		db, dbErr = database.NewWithWrappedTracing()
	} else {
		db, dbErr = database.New()
	}
	if dbErr != nil {
		return nil, dbErr
	}

	writer := &Writer{
		config: cfg,
		ctx:    ctx,
		db:     db,
	}

	return writer, nil
}

func (w *Writer) Start() error {
	logging.Log.Info("Start WRITER")

	if !w.running {
		go w.startWriter()
		w.running = true
		logging.SugaredLog.Infof("WRITER started")
		return nil
	}

	return fmt.Errorf("WRITER start failed: writer already running")
}

func (w *Writer) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown WRITER, timeout %d", timeout)

	if w.running {
		w.ticker.Stop()
		w.ctx.Done()
		time.Sleep(time.Duration(timeout) * time.Second)
		w.running = false
		return
	}

	logging.Log.Error("WRITER shutdown failed: writer not running")
}
