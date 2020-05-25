package kubernetes

import (
	"encoding/json"
	"net/http"

	"github.com/bygui86/go-timescaledb/writer/logging"
)

// TODO in case of error, return it as response
// TODO implement db connection check

const (
	contentTypeKey  = "Content-Type"
	contentTypeJson = "application/json"

	responseStatusOk    = "OK"
	responseStatusError = "ERROR"

	responseStatusOkCode    = 200
	responseStatusErrorCode = 500
)

func livenessHandler(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Debug("Liveness probe invoked")

	writer.Header().Set(contentTypeKey, contentTypeJson)
	err := json.NewEncoder(writer).Encode(
		Liveness{Status: responseStatusOk, Code: responseStatusOkCode},
	)
	if err != nil {
		logging.SugaredLog.Errorf("Error encoding liveness to JSON: %s", err.Error())
	}
}

func readinessHandler(writer http.ResponseWriter, request *http.Request) {
	logging.Log.Debug("Readiness probe invoked")

	generalStatus := responseStatusOk
	generalCode := responseStatusOkCode

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(
		Readiness{
			Status: generalStatus,
			Code:   generalCode,
		},
	)
	if err != nil {
		logging.SugaredLog.Errorf("Error encoding liveness to JSON: %s", err.Error())
	}
}
