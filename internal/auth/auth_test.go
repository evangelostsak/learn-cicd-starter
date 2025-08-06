package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 123456")
	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if apiKey != "123456" {
		t.Errorf("expected apiKey '123456', got '%s'", apiKey)
	}
}
