package smartcardfyi_test

import (
	"testing"

	smartcardfyi "github.com/fyipedia/smartcardfyi-go"
)

func TestNewClient(t *testing.T) {
	c := smartcardfyi.NewClient()
	if c.BaseURL != smartcardfyi.DefaultBaseURL {
		t.Errorf("expected %s, got %s", smartcardfyi.DefaultBaseURL, c.BaseURL)
	}
	if c.HTTPClient == nil {
		t.Error("expected non-nil HTTPClient")
	}
}

func TestSearch(t *testing.T) {
	c := smartcardfyi.NewClient()
	result, err := c.Search("javacard")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	if result.Total == 0 {
		t.Error("expected results, got 0")
	}
}
