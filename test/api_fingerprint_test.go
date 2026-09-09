package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/stretchr/testify/assert"
)

func TestApiFingerprint(t *testing.T) {
	t.Run("Create with empty config", func(t *testing.T) {
		client := sdk.NewAPIClient(nil)

		assert.NotNil(t, client)
	})

	t.Run("Handles error response with missing code field without panic", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"error": {"message": "Forbidden request without code"}}`))
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)
		client := sdk.NewAPIClient(cfg)

		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "test_key"})
		_, _, err := client.FingerprintApi.GetEvent(ctx, "req_123")

		assert.NotNil(t, err)
		assert.Equal(t, "Forbidden request without code", err.Error())
	})

	t.Run("Handles 429 response with missing code field without panic", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Retry-After", "5")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "Too many requests"}`))
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)
		client := sdk.NewAPIClient(cfg)

		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "test_key"})
		_, _, err := client.FingerprintApi.GetEvent(ctx, "req_123")

		assert.NotNil(t, err)
		assert.Equal(t, "Too many requests", err.Error())
	})
}

