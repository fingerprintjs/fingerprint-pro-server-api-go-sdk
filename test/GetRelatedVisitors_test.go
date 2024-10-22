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

const visitorId = "XIkiQhRyp7edU9SA0jBb"

func TestGetRelatedVisitors(t *testing.T) {
	t.Run("Returns visits", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.RelatedVisitorsResponse]("../test/mocks/related-visitors/get_related_visitors_200.json")

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

		res, _, err := client.FingerprintApi.GetRelatedVisitors(ctx, mockResponse.RelatedVisitors[0].VisitorId)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, mockResponse)
	})

	t.Run("Returns TooManyRequestsError", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorCommon429Response]("../test/mocks/shared/429_error_too_many_requests.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.NoError(t, parseErr)

			configFile := config.ReadConfig("../config.json")

			query := r.URL.Query()
			assert.Equal(t, query.Get("ii"), fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, query.Get("visitor_id"), visitorId)
			assert.Equal(t, r.URL.Path, "/related-visitors")

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

		_, res, err := client.FingerprintApi.GetRelatedVisitors(ctx, visitorId)
		assert.IsType(t, err, &sdk.TooManyRequestsError{})
		assert.Error(t, err)
		assert.Equal(t, 429, res.StatusCode)

		var tooManyRequestsError *sdk.TooManyRequestsError
		errors.As(err, &tooManyRequestsError)

		assert.IsType(t, tooManyRequestsError, &sdk.TooManyRequestsError{})

		assert.Equal(t, tooManyRequestsError.Code(), mockResponse.Error_.Code)
		assert.Equal(t, tooManyRequestsError.Error(), mockResponse.Error_.Message)
	})

	t.Run("Returns ErrorVisitor400Response on 400", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorVisitor400Response]("../test/mocks/shared/400_error_incorrect_visitor_id.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			query := r.URL.Query()
			assert.Equal(t, query.Get("ii"), fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, query.Get("visitor_id"), visitorId)
			assert.Equal(t, r.URL.Path, "/related-visitors")

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			err := json.NewEncoder(w).Encode(mockResponse)

			if err != nil {
				panic(err)
			}
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)

		client := sdk.NewAPIClient(cfg)

		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
			Key: "api_key",
		})

		_, res, err := client.FingerprintApi.GetRelatedVisitors(ctx, visitorId)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 400, res.StatusCode)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorVisitor400Response)
		assert.IsType(t, errorModel, &sdk.ErrorVisitor400Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorCommon403Response on 403", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorCommon403Response]("../test/mocks/shared/403_error_feature_not_enabled.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			query := r.URL.Query()
			assert.Equal(t, query.Get("ii"), fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, query.Get("visitor_id"), visitorId)
			assert.Equal(t, r.URL.Path, "/related-visitors")

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(403)
			err := json.NewEncoder(w).Encode(mockResponse)

			if err != nil {
				panic(err)
			}
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)

		client := sdk.NewAPIClient(cfg)

		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
			Key: "api_key",
		})

		_, res, err := client.FingerprintApi.GetRelatedVisitors(ctx, visitorId)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 403, res.StatusCode)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorCommon403Response)
		assert.IsType(t, errorModel, &sdk.ErrorCommon403Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorVisitor404Response on 404", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorVisitor404Response]("../test/mocks/shared/404_error_visitor_not_found.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			query := r.URL.Query()
			assert.Equal(t, query.Get("ii"), fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, query.Get("visitor_id"), visitorId)
			assert.Equal(t, r.URL.Path, "/related-visitors")

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			err := json.NewEncoder(w).Encode(mockResponse)

			if err != nil {
				panic(err)
			}
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)

		client := sdk.NewAPIClient(cfg)

		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
			Key: "api_key",
		})

		_, res, err := client.FingerprintApi.GetRelatedVisitors(ctx, visitorId)

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 404, res.StatusCode)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorVisitor404Response)
		assert.IsType(t, errorModel, &sdk.ErrorVisitor404Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

}
