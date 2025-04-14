package openai_test

import (
	"testing"

	openai "github.com/openai/go-sdk/pkg/openai"
)

func TestNewApiClient(t *testing.T) {
	client, err := openai.NewApiClient("test")
	if err != nil {
		t.Fatalf("Error creating new API client: %v", err)
	}
	if client == nil {
		t.Fatal("Expected a non-nil client")
	}
}
