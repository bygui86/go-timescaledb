package monitoring

import "github.com/prometheus/client_golang/prometheus"

const (
	namespace = "timeseries"
	subsystem = "writer"
)

var (
	insertions = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "insertion_total",
			Help:      "Number of insertion performed by the writer",
		},
	)

	insertionErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "insertion_errors_total",
			Help:      "Number of insertion error performed by the writer",
		},
	)

	insertionTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "insertion_execution_time_milliseconds",
			Help:      "Execution time (in milliseconds) of insertion performed by the writer",
			Buckets:   []float64{1e-10, 1e-8, 1e-6, 1e-4, 1e-2, 0.025, 0.05, 0.075, 0.1, 0.125, 0.25, 0.5, 1, 1.5, 2, 2.5, 5, 7.5, 10, 25, 50, 100, 250, 500, 750, 1000, 2500, 5000, 10000},
		},
	)
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(
		insertions,
		insertionErrors,
		insertionTime,
	)
}

func IncreaseInsertions() {
	go insertions.Inc()
}

func IncreaseInsertionErrors() {
	go insertionErrors.Inc()
}

func ObserveInsertionTime(timing float64) {
	insertionTime.Observe(timing)
}
