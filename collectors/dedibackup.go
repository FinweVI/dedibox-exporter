package collectors

import (
	"strconv"
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// DedibackupCollector is a collector for the Dedibackup-related API
type DedibackupCollector struct {
	dedibackupQuotaSpaceMetric     *prometheus.Desc
	dedibackupQuotaSpaceUsedMetric *prometheus.Desc
	dedibackupQuotaFilesMetric     *prometheus.Desc
	dedibackupQuotaFilesUsedMetric *prometheus.Desc
}

// NewDedibackupCollector is a helper function to spawn a new DedibackupCollector
func NewDedibackupCollector() *DedibackupCollector {
	return &DedibackupCollector{
		dedibackupQuotaSpaceMetric: prometheus.NewDesc(
			"dedibox_dedibackup_quota_space_total_bytes",
			"Dedibackup total space quota",
			[]string{"server_id", "active"},
			nil,
		),
		dedibackupQuotaSpaceUsedMetric: prometheus.NewDesc(
			"dedibox_dedibackup_quota_space_usage_bytes",
			"Dedibackup usage space quota",
			[]string{"server_id", "active"},
			nil,
		),
		dedibackupQuotaFilesMetric: prometheus.NewDesc(
			"dedibox_dedibackup_quota_files_total",
			"Dedibackup total files quota",
			[]string{"server_id", "active"},
			nil,
		),
		dedibackupQuotaFilesUsedMetric: prometheus.NewDesc(
			"dedibox_dedibackup_quota_files_usage",
			"Dedibackup usage files quota",
			[]string{"server_id", "active"},
			nil,
		),
	}
}

// Describe report all the metrics of the DedibackupCollector
func (collector *DedibackupCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.dedibackupQuotaSpaceMetric
	ch <- collector.dedibackupQuotaSpaceUsedMetric
	ch <- collector.dedibackupQuotaFilesMetric
	ch <- collector.dedibackupQuotaFilesUsedMetric
}

// Collect gather all the metrics of the DedibackupCollector
func (collector *DedibackupCollector) Collect(ch chan<- prometheus.Metric) {
	dedibackups, err := online.GetDedibackups()
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "dedibackup",
			"provider":  "online.net",
			"source":    "api",
		}).Error("Unable to retrieve informations")
		log.Debug(err)
		return
	}

	for _, ddbkp := range dedibackups {
		var dedibackupLabels []string
		// we get the integer part from "sd-1234"
		splt := strings.Split(ddbkp.Login, "-")
		sid := splt[len(splt)-1]

		dedibackupLabels = append(dedibackupLabels, sid)
		dedibackupLabels = append(dedibackupLabels, strconv.FormatBool(ddbkp.Active))

		ch <- prometheus.MustNewConstMetric(collector.dedibackupQuotaSpaceMetric, prometheus.GaugeValue, float64(ddbkp.QuotaSpace), dedibackupLabels...)
		ch <- prometheus.MustNewConstMetric(collector.dedibackupQuotaSpaceUsedMetric, prometheus.GaugeValue, float64(ddbkp.QuotaSpaceUsed), dedibackupLabels...)
		ch <- prometheus.MustNewConstMetric(collector.dedibackupQuotaFilesMetric, prometheus.GaugeValue, float64(ddbkp.QuotaFiles), dedibackupLabels...)
		ch <- prometheus.MustNewConstMetric(collector.dedibackupQuotaFilesUsedMetric, prometheus.GaugeValue, float64(ddbkp.QuotaFilesUsed), dedibackupLabels...)
	}
}
