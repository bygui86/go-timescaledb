package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/bygui86/go-timescaledb/writer/commons"
	"github.com/bygui86/go-timescaledb/writer/logging"
	"github.com/bygui86/go-timescaledb/writer/monitoring"
)

func InsertCondition(db *sql.DB, condition *commons.Condition, ctx context.Context) error {
	startTimer := time.Now()

	span := opentracing.StartSpan(
		"insert-condition-db",
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer span.Finish()

	span.SetTag("condition", condition.String())
	span.LogKV("condition", condition.String())

	result, err := db.ExecContext(ctx, insertConditionQuery, condition.Time, condition.Location, condition.Temperature)
	if err != nil {
		// tracing
		span.SetTag("condition-inserted", false)
		span.SetTag("error", err)
		span.LogKV("condition-inserted", false, "error", err)
		return err
	}

	// WARN: not supported by library
	// lastId, lastIdErr := result.LastInsertId()
	// if lastIdErr != nil {
	// 	// tracing
	// 	span.SetTag("condition-inserted", false)
	// 	span.SetTag("error", lastIdErr)
	// 	span.LogKV("condition-inserted", false, "error", lastIdErr)
	// 	return lastIdErr
	// }

	rowsAffected, rowsErr := result.RowsAffected()
	if rowsErr != nil {
		// tracing
		span.SetTag("condition-inserted", false)
		span.SetTag("error", rowsErr)
		span.LogKV("condition-inserted", false, "error", rowsErr)
		return rowsErr
	}

	logging.SugaredLog.Debugf("Rows affected: %d", rowsAffected)

	// tracing
	span.SetTag("condition-inserted", true)
	span.LogKV("condition-inserted", true)

	// monitoring
	monitoring.ObserveInsertionTime(float64(time.Now().Sub(startTimer).Milliseconds()))

	return nil
}
