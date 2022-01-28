package collectors

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestPlanCollector(t *testing.T) {
	c := NewPlanCollector(apiClient)

	const metadata = `
	        # HELP dedibox_plan Dedibox plans availability
	        # TYPE dedibox_plan gauge
	`

	expected := `
	        dedibox_plan{datacenter="ams1",name="core-7-l-a"} 5
	        dedibox_plan{datacenter="ams1",name="core-7-l-i"} 1
	        dedibox_plan{datacenter="ams1",name="core-7-m-a"} 1
	        dedibox_plan{datacenter="ams1",name="core-7-m-i"} 0
	        dedibox_plan{datacenter="ams1",name="core-7-xl-i"} 5
	        dedibox_plan{datacenter="ams1",name="pro-4-l"} 0
	        dedibox_plan{datacenter="ams1",name="pro-6-m"} 0
	        dedibox_plan{datacenter="ams1",name="pro-6-s"} 0
	        dedibox_plan{datacenter="ams1",name="pro-6-xs"} 0
	        dedibox_plan{datacenter="ams1",name="pro-7-s"} 0
	        dedibox_plan{datacenter="ams1",name="start-2-m-sata"} 0
	        dedibox_plan{datacenter="ams1",name="start-2-m-ssd"} 0
	        dedibox_plan{datacenter="ams1",name="start-2-s-sata"} 4
	        dedibox_plan{datacenter="ams1",name="start-2-s-ssd"} 0
	        dedibox_plan{datacenter="ams1",name="start-2-xs-sata"} 0
	        dedibox_plan{datacenter="ams1",name="start-3-l"} 0
	        dedibox_plan{datacenter="dc2",name="core-5-l"} 0
	        dedibox_plan{datacenter="dc2",name="core-7-m-i"} 0
	        dedibox_plan{datacenter="dc2",name="core-7-xxl-a"} 6
	        dedibox_plan{datacenter="dc2",name="pro-4-l"} 0
	        dedibox_plan{datacenter="dc2",name="pro-4-m"} 0
	        dedibox_plan{datacenter="dc2",name="pro-4-s"} 0
	        dedibox_plan{datacenter="dc2",name="pro-5-m"} 0
	        dedibox_plan{datacenter="dc2",name="pro-5-s"} 0
	        dedibox_plan{datacenter="dc2",name="pro-7-s"} 0
	        dedibox_plan{datacenter="dc2",name="start-1-l"} 0
	        dedibox_plan{datacenter="dc2",name="start-1-m-sata"} 0
	        dedibox_plan{datacenter="dc2",name="start-1-m-ssd"} 0
	        dedibox_plan{datacenter="dc2",name="start-1-m-ssd-160g"} 0
	        dedibox_plan{datacenter="dc2",name="start-2-l"} 0
	        dedibox_plan{datacenter="dc2",name="start-2-m-sata"} 1
	        dedibox_plan{datacenter="dc2",name="start-2-s-sata"} 3
	        dedibox_plan{datacenter="dc2",name="start-2-s-ssd"} 0
	        dedibox_plan{datacenter="dc2",name="start-2-xs-sata"} 0
	        dedibox_plan{datacenter="dc2",name="store-1-s"} 0
	        dedibox_plan{datacenter="dc3",name="core-2-m-sata"} 0
	        dedibox_plan{datacenter="dc3",name="core-2-m-ssd"} 0
	        dedibox_plan{datacenter="dc3",name="core-3-l-sata"} 0
	        dedibox_plan{datacenter="dc3",name="core-3-l-ssd"} 0
	        dedibox_plan{datacenter="dc3",name="core-3-m-sata"} 0
	        dedibox_plan{datacenter="dc3",name="core-3-m-ssd"} 0
	        dedibox_plan{datacenter="dc3",name="core-3-s-sata"} 0
	        dedibox_plan{datacenter="dc3",name="core-3-s-ssd"} 0
	        dedibox_plan{datacenter="dc3",name="core-4-m-sata"} 0
	        dedibox_plan{datacenter="dc3",name="pro-2-m-sata"} 0
	        dedibox_plan{datacenter="dc3",name="pro-3-l-sata"} 0
	        dedibox_plan{datacenter="dc3",name="pro-3-l-ssd"} 0
	        dedibox_plan{datacenter="dc3",name="pro-3-l-ssd-160"} 0
	        dedibox_plan{datacenter="dc3",name="pro-4-l"} 0
	        dedibox_plan{datacenter="dc3",name="start-1-l"} 0
	        dedibox_plan{datacenter="dc3",name="start-1-m-sata"} 0
	        dedibox_plan{datacenter="dc3",name="store-1-l"} 0
	        dedibox_plan{datacenter="dc3",name="store-1-s"} 0
	        dedibox_plan{datacenter="dc3",name="store-2-l"} 0
	        dedibox_plan{datacenter="dc3",name="store-2-m"} 0
	        dedibox_plan{datacenter="dc3",name="store-2-s"} 0
	        dedibox_plan{datacenter="dc3",name="store-2-xxl"} 0
	        dedibox_plan{datacenter="dc3",name="store-4-l"} 0
	        dedibox_plan{datacenter="dc3",name="store-4-m"} 0
	        dedibox_plan{datacenter="dc3",name="store-4-xl"} 0
	        dedibox_plan{datacenter="dc3",name="store-4-xxl"} 0
	        dedibox_plan{datacenter="dc5",name="core-4-l-sata"} 0
	        dedibox_plan{datacenter="dc5",name="core-4-l-ssd"} 0
	        dedibox_plan{datacenter="dc5",name="core-4-m-sata"} 0
	        dedibox_plan{datacenter="dc5",name="core-4-m-ssd"} 0
	        dedibox_plan{datacenter="dc5",name="core-4-s-sata"} 0
	        dedibox_plan{datacenter="dc5",name="core-4-s-ssd"} 0
	        dedibox_plan{datacenter="dc5",name="core-5-l"} 4
	        dedibox_plan{datacenter="dc5",name="core-5-m"} 101
	        dedibox_plan{datacenter="dc5",name="core-5-s"} 6
	        dedibox_plan{datacenter="dc5",name="core-5-xl"} 3
	        dedibox_plan{datacenter="dc5",name="core-6-m"} 1
	        dedibox_plan{datacenter="dc5",name="core-6-xs"} 0
	        dedibox_plan{datacenter="dc5",name="core-7-l-a"} 0
	        dedibox_plan{datacenter="dc5",name="core-7-l-i"} 0
	        dedibox_plan{datacenter="dc5",name="core-7-m-a"} 0
	        dedibox_plan{datacenter="dc5",name="core-7-m-i"} 0
	        dedibox_plan{datacenter="dc5",name="core-7-xl-i"} 0
	        dedibox_plan{datacenter="dc5",name="core-7-xxl-a"} 6
	        dedibox_plan{datacenter="dc5",name="pro-4-l"} 0
	        dedibox_plan{datacenter="dc5",name="pro-4-m"} 0
	        dedibox_plan{datacenter="dc5",name="pro-4-s"} 0
	        dedibox_plan{datacenter="dc5",name="pro-5-l"} 2
	        dedibox_plan{datacenter="dc5",name="pro-5-m"} 0
	        dedibox_plan{datacenter="dc5",name="pro-5-s"} 0
	        dedibox_plan{datacenter="dc5",name="pro-5-s-le"} 0
	        dedibox_plan{datacenter="dc5",name="pro-6-m"} 0
	        dedibox_plan{datacenter="dc5",name="pro-6-s"} 0
	        dedibox_plan{datacenter="dc5",name="pro-6-xs"} 0
	        dedibox_plan{datacenter="dc5",name="pro-7-m"} 0
	        dedibox_plan{datacenter="dc5",name="pro-7-s"} 147
	        dedibox_plan{datacenter="dc5",name="start-2-l"} 0
	        dedibox_plan{datacenter="dc5",name="start-2-m-ssd"} 0
	        dedibox_plan{datacenter="dc5",name="start-3-s-ssd"} 0
	        dedibox_plan{datacenter="dc5",name="store-2-l"} 0
	        dedibox_plan{datacenter="dc5",name="store-4-xl"} 0
	        dedibox_plan{datacenter="dc5",name="store-4-xxl"} 0
	`

	err := testutil.CollectAndCompare(
		c,
		strings.NewReader(metadata+expected),
		"dedibox_plan",
	)

	if err != nil {
		t.Error(err)
	}

	lintProb, err := testutil.CollectAndLint(c, "dedibox_plan")
	if err != nil {
		panic(err)
	} else {
		for _, prob := range lintProb {
			t.Logf("%s: %s", prob.Metric, prob.Text)
		}
	}
}
