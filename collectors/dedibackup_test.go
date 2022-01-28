package collectors

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestDedibackupCollector(t *testing.T) {
	c := NewDedibackupCollector(apiClient)

	expected := `
	        # HELP dedibox_dedibackup_quota_files_total Dedibackup total files quota
	        # TYPE dedibox_dedibackup_quota_files_total gauge
	        dedibox_dedibackup_quota_files_total{active="false",server_id="5678"} 70
	        dedibox_dedibackup_quota_files_total{active="true",server_id="1234"} 50
	        # HELP dedibox_dedibackup_quota_files_usage Dedibackup usage files quota
	        # TYPE dedibox_dedibackup_quota_files_usage gauge
	        dedibox_dedibackup_quota_files_usage{active="false",server_id="5678"} 50
	        dedibox_dedibackup_quota_files_usage{active="true",server_id="1234"} 5
	        # HELP dedibox_dedibackup_quota_space_total_bytes Dedibackup total space quota
	        # TYPE dedibox_dedibackup_quota_space_total_bytes gauge
	        dedibox_dedibackup_quota_space_total_bytes{active="false",server_id="5678"} 5000
	        dedibox_dedibackup_quota_space_total_bytes{active="true",server_id="1234"} 1000
	        # HELP dedibox_dedibackup_quota_space_usage_bytes Dedibackup usage space quota
	        # TYPE dedibox_dedibackup_quota_space_usage_bytes gauge
	        dedibox_dedibackup_quota_space_usage_bytes{active="false",server_id="5678"} 70
	        dedibox_dedibackup_quota_space_usage_bytes{active="true",server_id="1234"} 10
	`

	err := testutil.CollectAndCompare(
		c,
		strings.NewReader(expected),
		"dedibox_dedibackup_quota_space_total_bytes",
		"dedibox_dedibackup_quota_space_usage_bytes",
		"dedibox_dedibackup_quota_files_total",
		"dedibox_dedibackup_quota_files_usage",
	)

	if err != nil {
		t.Error(err)
	}

	lintProb, err := testutil.CollectAndLint(
		c,
		"dedibox_dedibackup_quota_space_total_bytes",
		"dedibox_dedibackup_quota_space_usage_bytes",
		"dedibox_dedibackup_quota_files_total",
		"dedibox_dedibackup_quota_files_usage",
	)
	if err != nil {
		panic(err)
	} else {
		for _, prob := range lintProb {
			t.Logf("%s: %s", prob.Metric, prob.Text)
		}
	}
}
