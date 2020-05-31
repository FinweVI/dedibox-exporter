package collectors

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
)

type abuseCollector struct {
	abuseMetric *prometheus.Desc
}

func NewAbuseCollector() *abuseCollector {
	return &abuseCollector{
		abuseMetric: prometheus.NewDesc(
			"dedibox_abuse",
			"Online.net's abuse and it's resolution status",
			[]string{"id", "sender", "service", "type"},
			nil,
		),
	}
}

func (collector *abuseCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.abuseMetric
}

func (collector *abuseCollector) Collect(ch chan<- prometheus.Metric) {
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
