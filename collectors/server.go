package collectors

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
)

type serverCollector struct {
	dedibackupQuotaSpaceMetric     *prometheus.Desc
	dedibackupQuotaSpaceUsedMetric *prometheus.Desc
	dedibackupQuotaFilesMetric     *prometheus.Desc
	dedibackupQuotaFilesUsedMetric *prometheus.Desc
}

func NewServerCollector() *serverCollector {
	return &serverCollector{
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

func (collector *serverCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.dedibackupQuotaSpaceMetric
	ch <- collector.dedibackupQuotaSpaceUsedMetric
	ch <- collector.dedibackupQuotaFilesMetric
	ch <- collector.dedibackupQuotaFilesUsedMetric
}

func (collector *serverCollector) Collect(ch chan<- prometheus.Metric) {
	dedibackups, err := online.GetDedibackups()
	if err != nil {
		fmt.Printf("Unable to get dedibackups informations!")
		return
	}

	for _, ddbkp := range dedibackups {
		var dedibackupLabels []string
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
