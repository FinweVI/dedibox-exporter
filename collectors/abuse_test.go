package collectors

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestAbuseCollector(t *testing.T) {
	c := NewAbuseCollector(apiClient)

	const metadata = `
		# HELP dedibox_pending_abuse Pending abuses
		# TYPE dedibox_pending_abuse counter
	`

	expected := `
		dedibox_pending_abuse{category="string",service="string"} 0
	`

	if err := testutil.CollectAndCompare(c, strings.NewReader(metadata+expected), "dedibox_pending_abuse"); err != nil {
		t.Error(err)
	}

	lintProb, err := testutil.CollectAndLint(c, "dedibox_pending_abuse")
	if err != nil {
		panic(err)
	} else {
		for _, prob := range lintProb {
			t.Logf("%s: %s", prob.Metric, prob.Text)
		}
	}
}
