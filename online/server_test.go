package online

import (
	"testing"
)

func TestGetDedibackups(t *testing.T) {
	_, err := apiClient.GetDedibackups()
	if err != nil {
		t.Error(err)
	}
}
