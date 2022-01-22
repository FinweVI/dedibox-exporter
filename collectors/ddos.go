package collectors

import (
	"strconv"
	"strings"
	"time"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type DDoSCollector struct {
	ddos      *prometheus.Desc
	ddosCount *prometheus.Desc
}

func NewDDoSCollector() *DDoSCollector {
	return &DDoSCollector{
		ddos: prometheus.NewDesc(
			"dedibox_ddos",
			"Dedibox ongoing DDoS activity",
			[]string{"id", "target", "mitigation_system", "attack_type"},
			nil,
		),
		ddosCount: prometheus.NewDesc(
			"dedibox_ddos_count_total",
			"Dedibox total count of ongoing DDoS alerts",
			[]string{},
			nil,
		),
	}
}

func (collector *DDoSCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.ddos
	ch <- collector.ddosCount
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

	ch <- prometheus.MustNewConstMetric(collector.ddosCount, prometheus.GaugeValue, float64(len(ddosList)))

	for _, ddos := range ddosList {
		var ddosLabels []string
		ddosLabels = append(ddosLabels, strconv.Itoa(ddos.ID))
		ddosLabels = append(ddosLabels, strings.ToLower(ddos.Target))
		ddosLabels = append(ddosLabels, strings.ToLower(ddos.MitigationSystem))
		ddosLabels = append(ddosLabels, strings.ToLower(ddos.AttackType))

		var sts float64 = 0
		_, err := time.Parse("2006-01-02", ddos.EndDate)
		// No error means EndDate is available so the DDoS is not ongoing
		if err == nil {
			sts = 1
		}

		ch <- prometheus.MustNewConstMetric(collector.ddos, prometheus.CounterValue, sts, ddosLabels...)
	}
}
