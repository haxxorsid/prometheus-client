package directinstrumentation

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "prometheus_client"
	subsystem = "directinstrumentation"
)

var (
	opsProcessed = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,             //optional
		Subsystem: subsystem,             //optional
		Name:      "processed_ops_total", //metric name will be: prometheus_client_directinstrumentation_processed_ops_total
		Help:      "The total number of processed events1",
	}, []string{"path"})
	memorySummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace:  namespace,
		Subsystem:  subsystem,
		Name:       "summary_memory_celsius",
		Help:       "The memory.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
	memoryHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "histogram_memory_celsius",
		Help:      "The memory.",
		Buckets:   prometheus.LinearBuckets(20, 5, 5), //0-20,21-25,26-30,31-35,36-40,41-infinity
		//Buckets: prometheus.ExponentialBuckets(10, 2, 3) //0-10, 11-20, 21-40, 41-infinity
	})
)

func RecordMetrics() {
	go func() {
		for {
			opsProcessed.WithLabelValues("/machine1/path1").Inc()
			time.Sleep(2 * time.Second)
			opsProcessed.WithLabelValues("/machine2/path2").Inc()
			memoryHistogram.Observe(float64(rand.Intn(100)))
			memorySummary.Observe(float64(rand.Intn(100)))
		}
	}()
}

func Register(registry *prometheus.Registry) {
	registry.MustRegister(
		opsProcessed,
		memoryHistogram,
		memorySummary,
		//collectors.NewGoCollector(),
		//collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})
	)
}

// Registering Metrics

// #1 - Register() method ABOVE

// #2 - USING SYNC

// var registry sync.Once
// func Register(registry *prometheus.Registry) {
//	 registry.Do(func() {
//		registry.MustRegister(opsProcessed1)
//	})
//}

// #3 - USING INIT

// func init() {
//	registry.MustRegister(opsProcessed1)
// }
