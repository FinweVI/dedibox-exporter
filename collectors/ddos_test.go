package collectors

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestDDoSCollector(t *testing.T) {
	c := NewDDoSCollector(apiClient)

	const metadata = `
		# HELP dedibox_ddos DDoS attacks on your services
		# TYPE dedibox_ddos counter
	`

	expected := `
		dedibox_ddos{attack_type="string",mitigation_system="arbor",target="string"} 1
	`

	if err := testutil.CollectAndCompare(c, strings.NewReader(metadata+expected), "dedibox_ddos"); err != nil {
		t.Error(err)
	}

	lintProb, err := testutil.CollectAndLint(c, "dedibox_ddos")
	if err != nil {
		panic(err)
	} else {
		for _, prob := range lintProb {
			t.Logf("%s: %s", prob.Metric, prob.Text)
		}
	}
}
