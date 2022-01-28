package collectors

import (
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// AbuseCollector is a collector for the abuses-related API
type AbuseCollector struct {
	apiClient *online.Client

	abuseMetric *prometheus.Desc
}

// NewAbuseCollector is a helper function to spawn a new AbuseCollector
func NewAbuseCollector(client *online.Client) *AbuseCollector {
	return &AbuseCollector{
		apiClient: client,

		abuseMetric: prometheus.NewDesc(
			"dedibox_pending_abuse",
			"Pending abuses",
			[]string{"service", "category"},
			nil,
		),
	}
}

// Describe report all the metrics of the AbuseCollector
func (collector *AbuseCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.abuseMetric
}

// Collect gather all the metrics of the AbuseCollector
func (collector *AbuseCollector) Collect(ch chan<- prometheus.Metric) {
	abuses, err := collector.apiClient.GetAbuses()
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "abuse",
			"provider":  "online.net",
			"source":    "api",
		}).Error("Unable to retrieve informations")
		log.Debug(err)
		return
	}

	for _, abs := range abuses {
		var abuseLabels []string
		abuseLabels = append(abuseLabels, strings.ToLower(abs.Service))
		abuseLabels = append(abuseLabels, strings.ToLower(abs.Category))

		ch <- prometheus.MustNewConstMetric(collector.abuseMetric, prometheus.CounterValue, float64(0), abuseLabels...)
	}
}
