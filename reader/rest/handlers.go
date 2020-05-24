package rest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bygui86/go-timescaledb/reader/commons"
	"github.com/bygui86/go-timescaledb/reader/database"
	"github.com/bygui86/go-timescaledb/reader/logging"
	"github.com/bygui86/go-timescaledb/reader/monitoring"
)

const (
	locationKey = "location"
	startKey    = "start" // RFC3339 = "2006-01-02T15:04:05Z07:00"
	endKey      = "end"   // RFC3339 = "2006-01-02T15:04:05Z07:00"
	limitKey    = "limit"

	startDefaultDiff = 24 * time.Hour
	limitDefault     = 100
)

func (s *Server) getConditions(writer http.ResponseWriter, request *http.Request) {
	span, ctx := retrieveSpanAndCtx(request, "get-conditions-handler")
	defer span.Finish()

	logging.Log.Info("Get conditions")

	params := s.getParameters(request)
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

func (s *Server) getParameters(request *http.Request) *commons.QueryParameters {
	params := &commons.QueryParameters{}

	now := time.Now().UTC() // endDefault
	startDefault := now.Add(-startDefaultDiff)

	// location
	params.Location = request.FormValue(locationKey)

	// limit
	limitStr := request.FormValue(limitKey)
	if limitStr != "" {
		limitTmp, limitErr := strconv.Atoi(limitStr)
		if limitErr != nil {
			logging.SugaredLog.Errorf("Error parsing limit %s to int: %s - default back to %d",
				limitStr, limitErr.Error(), limitDefault)
			params.Limit = limitDefault
		} else {
			params.Limit = int64(limitTmp)
		}
	} else {
		params.Limit = limitDefault
	}

	// end
	endStr := request.FormValue(endKey)
	if endStr != "" {
		endTmp, endErr := time.Parse(time.RFC3339, endStr)
		if endErr != nil {
			logging.SugaredLog.Errorf("Error parsing end %s to time: %s - default back to %v",
				endStr, endErr.Error(), now)
			params.EndTime = now
		} else {
			params.EndTime = endTmp
		}
	} else {
		params.EndTime = now
	}

	// start
	startStr := request.FormValue(startKey)
	if startStr != "" {
		startTmp, startErr := time.Parse(time.RFC3339, startStr)
		if startErr != nil {
			logging.SugaredLog.Errorf("Error parsing start %s to time: %s - default back to %v",
				startStr, startErr.Error(), startDefault)
			params.StartTime = startDefault
		} else {
			params.StartTime = startTmp
		}
	} else {
		params.StartTime = startDefault
	}

	// check start vs end
	if params.EndTime.After(now) {
		logging.SugaredLog.Errorf("Selected end %v is in the future - default back to %v",
			params.EndTime, now)
		params.EndTime = now
	}

	if params.StartTime.After(now) {
		logging.SugaredLog.Errorf("Selected start %v is in the future - default back to %v",
			params.StartTime, startDefault)
		params.StartTime = startDefault
	}

	if params.StartTime.After(params.EndTime) {
		logging.SugaredLog.Errorf("Selected start %v is after end %v - set start %v before end",
			params.StartTime, params.EndTime, startDefaultDiff)
		params.StartTime = params.EndTime.Add(-startDefaultDiff)
	}

	return params
}
