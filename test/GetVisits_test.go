package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/config"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVisits(t *testing.T) {
	t.Run("Returns visits", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.VisitorsGetResponse]("../test/mocks/get_visitors_200_limit_1.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "api_key"})

		res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", nil)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.VisitorId, mockResponse.VisitorId)
	})

	t.Run("Returns visits with pagination", func(t *testing.T) {
		opts := sdk.FingerprintApiGetVisitsOpts{
			RequestId:     "request_id",
			PaginationKey: "1683900801733.Ogvu1j",
			Limit:         500,
			LinkedId:      "request_id",
		}

		mockResponse := GetMockResponse[sdk.VisitorsGetResponse]("../test/mocks/get_visitors_200_limit_500.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.Nil(t, parseErr)

			assert.Equal(t, r.Form.Get("request_id"), opts.RequestId)
			assert.Equal(t, r.Form.Get("paginationKey"), fmt.Sprint(opts.PaginationKey))
			assert.Equal(t, r.Form.Get("limit"), fmt.Sprint(opts.Limit))
			assert.Equal(t, r.Form.Get("linked_id"), opts.LinkedId)

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
		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "api_key"})

		res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", &opts)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.VisitorId, mockResponse.VisitorId)
		assert.Equal(t, res.Visits, mockResponse.Visits)
		assert.Equal(t, res.PaginationKey, mockResponse.PaginationKey)
		assert.Equal(t, res.LastTimestamp, mockResponse.LastTimestamp)
	})

	t.Run("Returns TooManyRequestsError", func(t *testing.T) {
		opts := sdk.FingerprintApiGetVisitsOpts{
			RequestId:     "request_id",
			PaginationKey: "1683900801733.Ogvu1j",
			Limit:         500,
			LinkedId:      "request_id",
		}

		mockResponse := GetMockResponse[sdk.VisitorsGetResponse]("../test/mocks/get_visitors_429_too_many_requests.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.Nil(t, parseErr)

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
		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "api_key"})

		res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", &opts)
		assert.IsType(t, err, &sdk.TooManyRequestsError{})
		assert.Error(t, err)
		assert.NotNil(t, res)

		var tooManyRequestsError *sdk.TooManyRequestsError
		errors.As(err, &tooManyRequestsError)

		errorModel := tooManyRequestsError.Model().(*sdk.ErrorPlainResponse)
		assert.IsType(t, errorModel, &sdk.ErrorPlainResponse{})
		assert.Equal(t, int64(10), tooManyRequestsError.RetryAfter())
	})

	t.Run("Handles TooManyRequestsError without retry-after header", func(t *testing.T) {
		opts := sdk.FingerprintApiGetVisitsOpts{
			RequestId:     "request_id",
			PaginationKey: "1683900801733.Ogvu1j",
			Limit:         500,
			LinkedId:      "request_id",
		}

		mockResponse := GetMockResponse[sdk.VisitorsGetResponse]("../test/mocks/get_visitors_429_too_many_requests.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.Nil(t, parseErr)

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
		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "api_key"})

		res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", &opts)
		assert.IsType(t, err, &sdk.TooManyRequestsError{})
		assert.Error(t, err)
		assert.NotNil(t, res)

		var tooManyRequestsError *sdk.TooManyRequestsError
		errors.As(err, &tooManyRequestsError)

		errorModel := tooManyRequestsError.Model().(*sdk.ErrorPlainResponse)
		assert.IsType(t, errorModel, &sdk.ErrorPlainResponse{})
		assert.Equal(t, int64(0), tooManyRequestsError.RetryAfter())
	})

	t.Run("Returns ErrorPlainResponse on 400", func(t *testing.T) {
		opts := sdk.FingerprintApiGetVisitsOpts{
			RequestId:     "request_id",
			PaginationKey: "1683900801733.Ogvu1j",
			Limit:         500,
			LinkedId:      "request_id",
		}

		mockResponse := GetMockResponse[sdk.ErrorPlainResponse]("../test/mocks/get_visitors_400_bad_request.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.Nil(t, parseErr)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(mockResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)
		client := sdk.NewAPIClient(cfg)
		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "api_key"})

		res, _, err := client.FingerprintApi.GetVisits(ctx, "visitor_id", &opts)
		assert.IsType(t, err, &sdk.ApiError{})
		assert.Error(t, err)
		assert.NotNil(t, res)

		errorModel := err.Model().(*sdk.ErrorPlainResponse)
		assert.IsType(t, errorModel, &sdk.ErrorPlainResponse{})
		assert.Equal(t, err.Code(), sdk.FAILED)
		assert.Equal(t, err.Error(), mockResponse.Error_)
	})

}
