package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		header    string
		wantKey   string
		wantError bool
	}{
		{name: "valid", header: "ApiKey secret123", wantKey: "secret123"},
		{name: "missing", header: "", wantError: true},
		{name: "wrong scheme", header: "Bearer secret123", wantError: true},
		{name: "malformed", header: "ApiKey", wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}
			key, err := GetAPIKey(headers)
			if tt.wantError {
				if err == nil {
					t.Fatalf("expected error, got key %q", key)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if key != tt.wantKey {
				t.Fatalf("got %q, want %q", key, tt.wantKey)
			}
		})
	}
}
