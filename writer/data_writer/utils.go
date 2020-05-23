package data_writer

import (
	"context"
	"math/rand"
	"time"

	"github.com/bygui86/go-timescaledb/writer/database"
	"github.com/bygui86/go-timescaledb/writer/logging"
	"github.com/bygui86/go-timescaledb/writer/monitoring"

	"github.com/opentracing/opentracing-go"

	"github.com/bygui86/go-timescaledb/writer/commons"
)

const (
	locationA = "zurich"
	locationB = "new-york"
	locationC = "milan"

	randMin = -20.0
	randMax = 35.0
)

func (w *Writer) startWriter() {
	rand.Seed(time.Now().UnixNano())
	w.ticker = time.NewTicker(w.config.interval)
	counter := 0
	for {
		select {
		case <-w.ticker.C:

			span := opentracing.StartSpan("insert-condition")
			ctx := opentracing.ContextWithSpan(context.Background(), span)

			location := locationA
			if counter%2 == 0 {
				location = locationB
			}
			if counter%3 == 0 {
				location = locationC
			}

			condition := &commons.Condition{
				Time:        time.Now(),
				Location:    location,
				Temperature: randMin + rand.Float64()*(randMax-randMin),
			}

			span.SetTag("condition", condition.String())
			span.LogKV("condition", condition.String())

			err := database.InsertCondition(w.db, condition, ctx)
			if err != nil {
				// logging
				logging.SugaredLog.Errorf("Error inserting condition %s: %s", condition.String(), err.Error())
				// tracing
				span.SetTag("condition-inserted", false)
				span.SetTag("error", err)
				span.LogKV("condition-inserted", false, "error", err)
				// monitoring
				monitoring.IncreaseInsertionErrors()
			} else {
				// logging
				logging.SugaredLog.Infof("Condition %s successfully inserted", condition.String())
				// tracing
				span.SetTag("condition-inserted", true)
				span.LogKV("condition-inserted", true)
				// monitoring
				monitoring.IncreaseInsertions()
			}

			span.Finish()
			counter++

		case <-w.ctx.Done():
			return
		}
	}
}
