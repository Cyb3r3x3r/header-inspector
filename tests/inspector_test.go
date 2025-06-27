package tests

import (
	"testing"

	"github.com/Cyb3r3x3r/header-inspector/internal/inspector"
)

func TestFetchHeaders(t *testing.T) {
	url := "http://example.com"

	headers, err := inspector.FetchHeaders(url)
	if err != nil {
		t.Fatalf("got error while testing fetch headers: %v", err)
	}

	expectedHeaders := []string{
		"Content-Security-Policy",
		"Strict-Transport-Security",
		"X-Content-Type-Options",
		"X-Frame-Options",
		"Referrer-Policy",
		"Permissions-Policy",
		"Access-Control-Allow-Origin",
		"Set-Cookie",
	}

	for _, header := range expectedHeaders {
		if _, exists := headers[header]; !exists {
			t.Errorf("missing header key in result: %s", header)
		}
	}
}
