package collectors

import (
	"strconv"
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// ServerCollector is a collector for the Server-related API
type ServerCollector struct {
	dedibackupQuotaSpaceMetric     *prometheus.Desc
	dedibackupQuotaSpaceUsedMetric *prometheus.Desc
	dedibackupQuotaFilesMetric     *prometheus.Desc
	dedibackupQuotaFilesUsedMetric *prometheus.Desc
}

// NewServerCollector is a helper function to spawn a new ServerCollector
func NewServerCollector() *ServerCollector {
	return &ServerCollector{
		dedibackupQuotaSpaceMetric: prometheus.NewDesc(
			"dedibox_dedibackup_quota_space_total_bytes",
			"Get Dedibackup total space quota",
			[]string{"server_id", "active"},
			nil,
		),
		dedibackupQuotaSpaceUsedMetric: prometheus.NewDesc(
			"dedibox_dedibackup_quota_space_used_bytes",
			"Get Dedibackup space quota used",
			[]string{"server_id", "active"},
			nil,
		),
		dedibackupQuotaFilesMetric: prometheus.NewDesc(
			"dedibox_dedibackup_quota_files_total",
			"Get Dedibackup total quota files",
			[]string{"server_id", "active"},
			nil,
		),
		dedibackupQuotaFilesUsedMetric: prometheus.NewDesc(
			"dedibox_dedibackup_quota_files_used",
			"Get Dedibackup used quota files",
			[]string{"server_id", "active"},
			nil,
		),
	}
}

// Describe report all the metrics of the ServerCollector
func (collector *ServerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.dedibackupQuotaSpaceMetric
	ch <- collector.dedibackupQuotaSpaceUsedMetric
	ch <- collector.dedibackupQuotaFilesMetric
	ch <- collector.dedibackupQuotaFilesUsedMetric
}

// Collect gather all the metrics of the ServerCollector
func (collector *ServerCollector) Collect(ch chan<- prometheus.Metric) {
	dedibackups, err := online.GetDedibackups()
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "server",
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
