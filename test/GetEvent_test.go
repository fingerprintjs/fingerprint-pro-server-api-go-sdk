package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/config"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnsEvent(t *testing.T) {
	mockResponse := GetMockEventResponse("../test/mocks/get_event.json")

	ts := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		configFile := config.ReadConfig("../config.json")
		integrationInfo := r.URL.Query().Get("ii")
		assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
		assert.Equal(t, r.URL.Path, "/events/123")

		apiKey := r.Header.Get("Auth-Api-Key")
		assert.Equal(t, apiKey, "api_key")

		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(mockResponse)

		if err != nil {
			log.Fatal(err)
		}
	}))
	defer ts.Close()

	cfg := sdk.NewConfiguration()
	cfg.ChangeBasePath(ts.URL)

	client := sdk.NewAPIClient(cfg)

	ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
		Key: "api_key",
	})

	res, _, err := client.FingerprintApi.GetEvent(ctx, "123")

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res, mockResponse)
}
