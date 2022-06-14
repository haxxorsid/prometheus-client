package directinstrumentation

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

const (
	namespace = "msp_controller"
	subsystem = "directinstrumentation"
)

var (
	opsProcessed1 = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "processed_ops_total",
		Help:      "The total number of processed events1",
	}, []string{"path"})
)

func RecordMetrics() {
	go func() {
		for {
			opsProcessed1.WithLabelValues("path1").Inc()
			time.Sleep(2 * time.Second)
			opsProcessed1.WithLabelValues("path2").Inc()
		}
	}()
}

func Register(registry *prometheus.Registry) {
	registry.MustRegister(opsProcessed1, collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
}
