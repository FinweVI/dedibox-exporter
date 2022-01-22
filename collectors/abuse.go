package collectors

import (
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// AbuseCollector is a collector for the abuses-related API
type AbuseCollector struct {
	abuseMetric      *prometheus.Desc
	abuseCountMetric *prometheus.Desc
}

// NewAbuseCollector is a helper function to spawn a new AbuseCollector
func NewAbuseCollector() *AbuseCollector {
	return &AbuseCollector{
		abuseMetric: prometheus.NewDesc(
			"dedibox_pending_abuse",
			"Pending abuses",
			[]string{"service", "category"},
			nil,
		),
		abuseCountMetric: prometheus.NewDesc(
			"dedibox_pending_abuse_count",
			"Pending abuse count",
			[]string{},
			nil,
		),
	}
}

// Describe report all the metrics of the AbuseCollector
func (collector *AbuseCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.abuseMetric
	ch <- collector.abuseCountMetric
}

// Collect gather all the metrics of the AbuseCollector
func (collector *AbuseCollector) Collect(ch chan<- prometheus.Metric) {
	abuses, err := online.GetAbuses()
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "abuse",
			"provider":  "online.net",
			"source":    "api",
		}).Error("Unable to retrieve informations")
		log.Debug(err)
		return
	}

	ch <- prometheus.MustNewConstMetric(collector.abuseCountMetric, prometheus.GaugeValue, float64(len(abuses)))

	for _, abs := range abuses {
		var abuseLabels []string
		abuseLabels = append(abuseLabels, strings.ToLower(abs.Service))
		abuseLabels = append(abuseLabels, strings.ToLower(abs.Category))

		ch <- prometheus.MustNewConstMetric(collector.abuseMetric, prometheus.CounterValue, float64(0), abuseLabels...)
	}
}
