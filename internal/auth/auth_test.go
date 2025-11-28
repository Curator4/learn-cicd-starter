package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers     http.Header
		expectedKey string
		expectError bool
	}{
		"valid key": {
			headers:     http.Header{"Authorization": []string{"ApiKey abc123"}},
			expectedKey: "abc123",
			expectError: false,
		},
		"missing auth header": {
			headers:     http.Header{"Authorizationn": []string{"ApiKey abc123"}},
			expectedKey: "",
			expectError: true,
		},
		"auth header with wrong format": {
			headers:     http.Header{"Authorization": []string{"ApiKey: abc123"}},
			expectedKey: "",
			expectError: true,
		},
		"empty auth header": {
			headers:     http.Header{"Authorization": []string{""}},
			expectedKey: "",
			expectError: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)
			if tc.expectedKey != key {
				t.Errorf("expected %#v, got %#v", tc.expectedKey, key)
			}
			if tc.expectError && err == nil {
				t.Errorf("expcted an error, got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("expected no error, got %v:", err)
			}
		})
	}
}
