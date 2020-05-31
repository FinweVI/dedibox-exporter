package collectors

import (
	"fmt"
	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
)

type planCollector struct {
	dediPlan *prometheus.Desc
}

func NewPlanCollector() *planCollector {
	return &planCollector{
		dediPlan: prometheus.NewDesc(
			"dedibox_plan",
			"Get Dedibox plan availability",
			[]string{"name", "datacenter"},
			nil,
		),
	}
}

func (collector *planCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.dediPlan
}

func (collector *planCollector) Collect(ch chan<- prometheus.Metric) {
	plans, err := online.GetPlans()
	if err != nil {
		fmt.Printf("Unable to get plans informations!")
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