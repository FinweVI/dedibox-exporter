package online

import (
	"os"
	"testing"

	"github.com/FinweVI/dedibox-exporter/online/testutil"
)

var (
	apiClient *Client
)

func TestMain(m *testing.M) {
	serverURL, teardown := testutil.SetupMockAPI()
	defer teardown()

	apiClient = mockClient(serverURL)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func mockClient(serverURL string) *Client {
	var err error

	apiClient, err = NewClient(BaseURL(serverURL), AuthToken("testToken"))
	if err != nil {
		panic(err)
	}

	return apiClient
}
