package collectors

import (
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type DDoSCollector struct {
	ddosMetric      *prometheus.Desc
	ddosCountMetric *prometheus.Desc
}

func NewDDoSCollector() *DDoSCollector {
	return &DDoSCollector{
		ddosMetric: prometheus.NewDesc(
			"dedibox_ddos",
			"DDoS attacks on your services",
			[]string{"target", "mitigation_system", "attack_type"},
			nil,
		),
		ddosCountMetric: prometheus.NewDesc(
			"dedibox_ddos_count",
			"Recent DDoS attacks count",
			[]string{},
			nil,
		),
	}
}

func (collector *DDoSCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.ddosMetric
	ch <- collector.ddosCountMetric
}

func (collector *DDoSCollector) Collect(ch chan<- prometheus.Metric) {
	ddosList, err := online.GetDDoS()
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "ddos",
			"provider":  "online.net",
			"source":    "api",
		}).Error("Unable to retrieve informations")
		log.Debug(err)
		return
	}

	ch <- prometheus.MustNewConstMetric(collector.ddosCountMetric, prometheus.GaugeValue, float64(len(ddosList)))

	for _, ddos := range ddosList {
		var ddosLabels []string
		ddosLabels = append(ddosLabels, strings.ToLower(ddos.Target))
		ddosLabels = append(ddosLabels, strings.ToLower(ddos.MitigationSystem))
		ddosLabels = append(ddosLabels, strings.ToLower(ddos.AttackType))

		var sts float64 = 0
		if !ddos.EndDate.IsZero() {
			// If EndDate is not set to the Zero time, the DDoS is not ongoing
			sts = 1
		}

		ch <- prometheus.MustNewConstMetric(collector.ddosMetric, prometheus.CounterValue, sts, ddosLabels...)
	}
}
