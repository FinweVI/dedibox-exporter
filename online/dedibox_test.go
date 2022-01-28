package online

import (
	"testing"
)

func TestGetPlans(t *testing.T) {
	_, err := apiClient.GetPlans()
	if err != nil {
		t.Error(err)
	}
}
