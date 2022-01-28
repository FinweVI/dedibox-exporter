package online

import (
	"testing"
)

func TestGetDDoS(t *testing.T) {
	_, err := apiClient.GetDDoS()
	if err != nil {
		t.Error(err)
	}
}
