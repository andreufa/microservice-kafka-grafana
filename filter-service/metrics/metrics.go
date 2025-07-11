package metrics

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestMetrics = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "filter_service",
	Subsystem: "requests",
	Name:      "total",
	Help:      "Count of request to filter-service",
}, []string{"method", "handler", "code"})

func ObserveRequest(method string, handler string, code int) {
	requestMetrics.With(prometheus.Labels{
		"method":  method,
		"handler": handler,
		"code":    strconv.Itoa(code),
	}).Inc()
}

func Listen(addres string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(addres, mux)
}
