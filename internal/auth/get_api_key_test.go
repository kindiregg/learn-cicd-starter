package auth

import (
	"net/http"
	"testing"
)

func TestGetAPI(t *testing.T) {
	tests := []struct {
		name     string
		headers  http.Header
		expected string
		wantErr  bool
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey api-key-123"},
			},
			expected: "api-key-123",
			wantErr:  false,
		},
		{
			name:     "missing header",
			headers:  http.Header{},
			expected: "",
			wantErr:  true,
		},
		{
			name: "misspelled header",
			headers: http.Header{
				"Authorization": []string{"Key api-key-123"},
			},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.expected {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.expected)
			}
		})
	}
}
