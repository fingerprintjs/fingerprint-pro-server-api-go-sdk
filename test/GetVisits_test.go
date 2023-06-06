package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v3/config"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v3/sdk"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnsVisits(t *testing.T) {
	mockResponse := GetMockResponse("../test/mocks/visits_limit_1.json")

	ts := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		configFile := config.ReadConfig("../config.json")
		integrationInfo := r.URL.Query().Get("ii")
		assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))

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

	res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", nil)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.VisitorId, mockResponse.VisitorId)
}

func TestReturnsVisitsWithPagination(t *testing.T) {
	opts := sdk.FingerprintApiGetVisitsOpts{
		RequestId:     optional.NewString("request_id"),
		PaginationKey: optional.NewString("1683900801733.Ogvu1j"),
		Limit:         optional.NewInt32(500),
		LinkedId:      optional.NewString("request_id"),
	}

	mockResponse := GetMockResponse("../test/mocks/visits_limit_500.json")

	ts := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		parseErr := r.ParseForm()

		assert.NoError(t, parseErr)

		assert.Equal(t, r.Form.Get("request_id"), opts.RequestId.Value())
		assert.Equal(t, r.Form.Get("paginationKey"), fmt.Sprint(opts.PaginationKey.Value()))
		assert.Equal(t, r.Form.Get("limit"), fmt.Sprint(opts.Limit.Value()))
		assert.Equal(t, r.Form.Get("linked_id"), opts.LinkedId.Value())

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

	res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", &opts)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.VisitorId, mockResponse.VisitorId)
	assert.Equal(t, res.Visits, mockResponse.Visits)
	assert.Equal(t, res.PaginationKey, mockResponse.PaginationKey)
	assert.Equal(t, res.LastTimestamp, mockResponse.LastTimestamp)
}

func TestHandlesTooManyRequestsError(t *testing.T) {
	opts := sdk.FingerprintApiGetVisitsOpts{
		RequestId:     optional.NewString("request_id"),
		PaginationKey: optional.NewString("1683900801733.Ogvu1j"),
		Limit:         optional.NewInt32(500),
		LinkedId:      optional.NewString("request_id"),
	}

	mockResponse := GetMockResponse("../test/mocks/visits_too_many_requests_error.json")

	ts := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		parseErr := r.ParseForm()

		assert.NoError(t, parseErr)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Retry-After", "10")

		w.WriteHeader(http.StatusTooManyRequests)

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

	res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", &opts)

	assert.IsType(t, err, sdk.GenericSwaggerError{})
	assert.Error(t, err)
	assert.NotNil(t, res)

	errorModel := err.(sdk.GenericSwaggerError).Model().(sdk.ManyRequestsResponse)

	assert.IsType(t, errorModel, sdk.ManyRequestsResponse{})
	assert.Equal(t, int64(10), errorModel.RetryAfter)
}

func TestHandlesTooManyRequestsErrorWithoutRetryAfterHeader(t *testing.T) {
	opts := sdk.FingerprintApiGetVisitsOpts{
		RequestId:     optional.NewString("request_id"),
		PaginationKey: optional.NewString("1683900801733.Ogvu1j"),
		Limit:         optional.NewInt32(500),
		LinkedId:      optional.NewString("request_id"),
	}

	mockResponse := GetMockResponse("../test/mocks/visits_too_many_requests_error.json")

	ts := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		parseErr := r.ParseForm()

		assert.NoError(t, parseErr)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusTooManyRequests)

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

	res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", &opts)

	assert.IsType(t, err, sdk.GenericSwaggerError{})
	assert.Error(t, err)
	assert.NotNil(t, res)

	errorModel := err.(sdk.GenericSwaggerError).Model().(sdk.ManyRequestsResponse)

	assert.IsType(t, errorModel, sdk.ManyRequestsResponse{})
	assert.Equal(t, int64(0), errorModel.RetryAfter)
}
