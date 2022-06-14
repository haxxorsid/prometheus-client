package main

import (
	"net/http"

	"github.com/haxxorsid/prometheus-client/pk1"
	"github.com/haxxorsid/prometheus-client/pk2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	registry := prometheus.NewRegistry()

	pk1.Register(registry)
	pk2.Register(registry)
	pk1.RecordMetrics()
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.ListenAndServe(":5555", nil)
}
