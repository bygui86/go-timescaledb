package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/bygui86/go-timescaledb/reader/commons"
	"github.com/bygui86/go-timescaledb/reader/logging"
	"github.com/bygui86/go-timescaledb/reader/monitoring"
)

func GetConditions(db *sql.DB, params *commons.QueryParameters, ctx context.Context) ([]*commons.Condition, error) {
	startTimer := time.Now()

	span := opentracing.StartSpan(
		"get-conditions-db",
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer span.Finish()

	var whereStatement string
	if params.Location == "" {
		whereStatement = dateRangeOnlyWhereConditionsQuery
		span.SetTag("start", params.StartTime)
		span.SetTag("end", params.EndTime)
		span.LogKV("start", params.StartTime, "end", params.EndTime)
	} else {
		whereStatement = locationDateRangeWhereConditionsQuery
		span.SetTag("location", params.Location)
		span.SetTag("start", params.StartTime)
		span.SetTag("end", params.EndTime)
		span.LogKV("location", params.Location, "start", params.StartTime, "end", params.EndTime)
	}

	query := fmt.Sprintf(queryFormat,
		selectConditionsQuery, whereStatement, orderbyConditionsQuery,
		fmt.Sprintf(limitConditionsQuery, params.Limit))

	logging.SugaredLog.Debugf("Query: %s - Values: %s, %v, %v",
		query, params.Location, params.StartTime, params.EndTime)

	span.SetTag("query", query)
	span.LogKV("query", query)

	var rows *sql.Rows
	var err error
	if params.Location == "" {
		rows, err = db.QueryContext(ctx, query, params.StartTime, params.EndTime)
	} else {
		rows, err = db.QueryContext(ctx, query, params.Location, params.StartTime, params.EndTime)
	}
	if err != nil {
		// tracing
		span.SetTag("results", 0)
		span.SetTag("error", err)
		span.LogKV("results", 0, "error", err)
		return nil, err
	}
	defer rows.Close()

	conditions := make([]*commons.Condition, 0)
	for rows.Next() {
		var cond commons.Condition
		if err := rows.Scan(&cond.Time, &cond.Location, &cond.Temperature); err != nil {
			return nil, err
		}
		conditions = append(conditions, &cond)
	}

	logging.SugaredLog.Debugf("Found %d results", len(conditions))

	// tracing
	span.SetTag("results", len(conditions))
	span.LogKV("results", len(conditions))

	// monitoring
	if params.Location == "" {
		monitoring.ObserveSelectTime(float64(time.Now().Sub(startTimer).Milliseconds()))
	} else {
		monitoring.ObserveFilteredSelectTime(float64(time.Now().Sub(startTimer).Milliseconds()))
	}

	return conditions, nil
}
