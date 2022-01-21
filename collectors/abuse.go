package collectors

import (
	"fmt"
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
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
			"dedibox_abuse",
			"Online.net's pending abuse and it's details",
			[]string{"id", "sender", "service", "type"},
			nil,
		),
		abuseCountMetric: prometheus.NewDesc(
			"dedibox_abuse_count_total",
			"Online.net's total unresolved abuses",
			[]string{},
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
	abuses, err := online.GetAbuses()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Unable to get abuses informations!")
		return
	}

	for _, abs := range abuses {
		var abuseLabels []string
		abuseLabels = append(abuseLabels, strconv.Itoa(abs.ID))
		abuseLabels = append(abuseLabels, strings.ToLower(abs.Sender))
		abuseLabels = append(abuseLabels, strings.ToLower(abs.Service))
		abuseLabels = append(abuseLabels, strings.ToLower(abs.Category))

		var sts float64 = 0
		if abs.Status == "Resolved" {
			sts = 1
		}

		ch <- prometheus.MustNewConstMetric(collector.abuseMetric, prometheus.CounterValue, sts, abuseLabels...)
	}
}
