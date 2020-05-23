package monitoring

import "github.com/prometheus/client_golang/prometheus"

const (
	namespace = "timeseries"
	subsystem = "reader"
)

var (
	requests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "requests_total",
			Help:      "Number of requests managed by the reader",
		},
	)

	requestErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "request_errors_total",
			Help:      "Number of request errors managed by the reader",
		},
	)

	selectTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "select_execution_time_milliseconds",
			Help:      "Execution time (in milliseconds) of plain select statement performed by the reader",
			Buckets:   []float64{1e-10, 1e-8, 1e-6, 1e-4, 1e-2, 0.025, 0.05, 0.075, 0.1, 0.125, 0.25, 0.5, 1, 1.5, 2, 2.5, 5, 7.5, 10, 25, 50, 100, 250, 500, 750, 1000, 2500, 5000, 10000},
		},
	)

	filteredSelectTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "filtered_select_execution_time_milliseconds",
			Help:      "Execution time (in milliseconds) of filtered select statement performed by the reader",
			Buckets:   []float64{1e-10, 1e-8, 1e-6, 1e-4, 1e-2, 0.025, 0.05, 0.075, 0.1, 0.125, 0.25, 0.5, 1, 1.5, 2, 2.5, 5, 7.5, 10, 25, 50, 100, 250, 500, 750, 1000, 2500, 5000, 10000},
		},
	)
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(
		requests,
		requestErrors,
		selectTime,
		filteredSelectTime,
	)
}

func IncreaseRequests() {
	go requests.Inc()
}

func IncreaseRequestErrors() {
	go requestErrors.Inc()
}

func ObserveSelectTime(timing float64) {
	selectTime.Observe(timing)
}

func ObserveFilteredSelectTime(timing float64) {
	filteredSelectTime.Observe(timing)
}
