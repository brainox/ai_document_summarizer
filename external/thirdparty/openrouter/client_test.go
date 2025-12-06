package openrouter

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	apiKey := "test-key"
	model := "test-model"

	client := NewClient(apiKey, model)

	if client.APIKey != apiKey {
		t.Errorf("Expected APIKey to be %s, got %s", apiKey, client.APIKey)
	}

	if client.Model != model {
		t.Errorf("Expected Model to be %s, got %s", model, client.Model)
	}
}

func TestAnalyze(t *testing.T) {
	// This is a unit test skeleton
	// In production, you would mock the HTTP client
	client := NewClient("test-key", "test-model")

	if client == nil {
		t.Error("Expected client to be initialized")
	}
}
