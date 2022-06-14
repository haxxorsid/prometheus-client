package main

import (
	"net/http"

	"github.com/haxxorsid/prometheus-client/collector"
	"github.com/haxxorsid/prometheus-client/directinstrumentation"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	registry := prometheus.NewRegistry()

	directinstrumentation.Register(registry)
	collector.Register(registry)
	directinstrumentation.RecordMetrics()
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.ListenAndServe(":5555", nil)
}
