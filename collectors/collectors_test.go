package collectors

import (
	"os"
	"testing"

	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/FinweVI/dedibox-exporter/online/testutil"
)

var (
	apiClient *online.Client
)

func TestMain(m *testing.M) {
	serverURL, teardown := testutil.SetupMockAPI()
	defer teardown()

	apiClient = mockClient(serverURL)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func mockClient(serverURL string) *online.Client {
	var err error

	apiClient, err = online.NewClient(
		online.BaseURL(serverURL),
		online.AuthToken("testToken"),
	)
	if err != nil {
		panic(err)
	}

	return apiClient
}
