package collector

import (
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "prometheus_client"
	subsystem = "collector"
)

type volumeCollector struct {
	volumeBytesTotal *prometheus.Desc
	volumeBytesFree  *prometheus.Desc
	volumeBytesUsed  *prometheus.Desc
}

func newVolumeCollector() *volumeCollector {
	return &volumeCollector{
		volumeBytesTotal: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "bytes_total"),
			"Total size of the volume/disk",
			[]string{"volume_name", "volume_path"}, nil,
		),
		volumeBytesFree: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "bytes_free"),
			"Free size of the volume/disk",
			[]string{"volume_name", "volume_path"}, nil,
		),
		volumeBytesUsed: prometheus.NewDesc(prometheus.BuildFQName(namespace, subsystem, "bytes_used"),
			"Used size of volume/disk",
			[]string{"volume_name", "volume_path"}, nil,
		),
	}
}

func (collector *volumeCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.volumeBytesTotal
	ch <- collector.volumeBytesFree
	ch <- collector.volumeBytesUsed
}

func (collector *volumeCollector) Collect(ch chan<- prometheus.Metric) {

	var metricValue float64
	if 1 == 1 {
		metricValue = float64(rand.Intn(100))
	}
	ch <- prometheus.MustNewConstMetric(collector.volumeBytesTotal, prometheus.GaugeValue, metricValue, "log", "path")
	ch <- prometheus.MustNewConstMetric(collector.volumeBytesFree, prometheus.GaugeValue, metricValue, "log", "path")
	ch <- prometheus.MustNewConstMetric(collector.volumeBytesUsed, prometheus.GaugeValue, float64(10), "log", "path")

}

func Register(registry *prometheus.Registry) {
	collector := newVolumeCollector()
	registry.MustRegister(collector)
}
