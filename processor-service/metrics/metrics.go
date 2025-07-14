package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestMetrics = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "base_service",
	Subsystem: "kafka_message",
	Name:      "total",
	Help:      "Count of kafka message to processor-service",
}, []string{"topic", "error"})

func ObserveRequest(topic string, errorMsg string) {
	requestMetrics.With(prometheus.Labels{
		"topic": topic,
		"error": errorMsg,
	}).Inc()
}

func Listen(addres string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(addres, mux)
}
