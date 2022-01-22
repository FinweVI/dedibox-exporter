package collectors

import (
	"strings"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// PlanCollector is a collector for the dedibox-related API focused on the plans availability
type PlanCollector struct {
	dediPlan *prometheus.Desc
}

// NewPlanCollector is a helper function to spawn a new PlanCollector
func NewPlanCollector() *PlanCollector {
	return &PlanCollector{
		dediPlan: prometheus.NewDesc(
			"dedibox_plan",
			"Get Dedibox plan availability",
			[]string{"name", "datacenter"},
			nil,
		),
	}
}

// Describe report all the metrics of the PlanCollector
func (collector *PlanCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.dediPlan
}

// Collect gather all the metrics of the PlanCollector
func (collector *PlanCollector) Collect(ch chan<- prometheus.Metric) {
	plans, err := online.GetPlans()
	if err != nil {
		log.WithFields(log.Fields{
			"collector": "dedibox",
			"provider":  "online.net",
			"source":    "api",
		}).Error("Unable to retrieve informations")
		log.Debug(err)
		return
	}

	for _, plan := range plans {
		for _, stock := range plan.Stocks {
			var planLabels []string
			planLabels = append(planLabels, strings.ToLower(plan.Slug))
			planLabels = append(planLabels, strings.ToLower(stock.Datacenter.Name))

			ch <- prometheus.MustNewConstMetric(collector.dediPlan, prometheus.GaugeValue, float64(stock.Stock), planLabels...)
		}
	}
}
