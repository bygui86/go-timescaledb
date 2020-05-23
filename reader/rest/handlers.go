package rest

import (
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/bygui86/go-timescaledb/reader/commons"
	"github.com/bygui86/go-timescaledb/reader/database"
	"github.com/bygui86/go-timescaledb/reader/logging"
	"github.com/bygui86/go-timescaledb/reader/monitoring"
)

const (
	locationKey = "location"
	startKey    = "start"
	endKey      = "end"
	limitKey    = "limit"

	limitDefault = 100
)

// TODO extract parameters from request: start, end
func (s *Server) getConditions(writer http.ResponseWriter, request *http.Request) {
	span, ctx := retrieveSpanAndCtx(request, "get-conditions-handler")
	defer span.Finish()

	logging.Log.Info("Get conditions")

	now := time.Now().UTC()
	params := &commons.QueryParameters{
		StartTime: now.Add(-24 * time.Hour),
		EndTime:   now,
		Limit:     limitDefault,
	}

	params.Location = request.FormValue(locationKey)
	// TODO start
	// TODO end
	limitStr := request.FormValue(limitKey)
	if limitStr != "" {
		limit, limitErr := strconv.Atoi(limitStr)
		if limitErr != nil {
			// logging
			errMsg := "Get conditions failed: " + limitErr.Error()
			sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

			// tracing
			span.SetTag("results", 0)
			span.SetTag("error", errMsg)
			span.LogKV("results", 0, "error", errMsg)

			// monitoring
			monitoring.IncreaseRequestErrors()
			return
		}
		params.Limit = int64(limit)
	}
	logging.SugaredLog.Debugf("Parameters: %s", params.String())

	conditions, err := database.GetConditions(s.db, params, ctx)
	if err != nil {
		// logging
		errMsg := "Get conditions failed: " + err.Error()
		sendErrorResponse(writer, http.StatusInternalServerError, errMsg)

		// tracing
		span.SetTag("results", 0)
		span.SetTag("error", errMsg)
		span.LogKV("results", 0, "error", errMsg)

		// monitoring
		monitoring.IncreaseRequestErrors()
		return
	}

	sendJsonResponse(writer, http.StatusOK, conditions)

	// tracing
	span.SetTag("results", len(conditions))
	span.LogKV("results", len(conditions))

	// monitoring
	monitoring.IncreaseRequests()
}
