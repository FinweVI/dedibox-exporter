package online

import (
	"testing"
)

func TestGetAbuses(t *testing.T) {
	_, err := apiClient.GetAbuses()
	if err != nil {
		t.Fatal(err)
	}
}
