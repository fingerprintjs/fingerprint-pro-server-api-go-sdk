package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v6/config"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v6/sdk"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteVisitorData(t *testing.T) {
	t.Run("Returns 200 on success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/visitors/123")
			assert.Equal(t, r.Method, http.MethodDelete)

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)

		client := sdk.NewAPIClient(cfg)

		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
			Key: "api_key",
		})

		res, err := client.FingerprintApi.DeleteVisitorData(ctx, "123")

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("Returns ErrorVisitor404Response on 404", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorVisitor404Response]("../test/mocks/404_error_visitor_not_found.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/visitors/123")
			assert.Equal(t, r.Method, http.MethodDelete)

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

		res, err := client.FingerprintApi.DeleteVisitorData(ctx, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 404)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorVisitor404Response)
		assert.IsType(t, errorModel, &sdk.ErrorVisitor404Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns TooManyRequestsError on 429", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorCommon429Response]("../test/mocks/429_error_too_many_requests.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/visitors/123")
			assert.Equal(t, r.Method, http.MethodDelete)

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(429)
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

		res, err := client.FingerprintApi.DeleteVisitorData(ctx, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 429)

		var tooManyRequestsError *sdk.TooManyRequestsError
		errors.As(err, &tooManyRequestsError)

		assert.IsType(t, tooManyRequestsError, &sdk.TooManyRequestsError{})

		assert.Equal(t, tooManyRequestsError.Code(), mockResponse.Error_.Code)
		assert.Equal(t, tooManyRequestsError.Error(), mockResponse.Error_.Message)
	})

	t.Run("Returns ErrorVisitor400Response on 400", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorVisitor400Response]("../test/mocks/400_error_incorrect_visitor_id.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/visitors/123")
			assert.Equal(t, r.Method, http.MethodDelete)

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

		res, err := client.FingerprintApi.DeleteVisitorData(ctx, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 400)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorVisitor400Response)
		assert.IsType(t, errorModel, &sdk.ErrorVisitor400Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorCommon403Response on 403", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorCommon403Response]("../test/mocks/403_error_feature_not_enabled.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/visitors/123")
			assert.Equal(t, r.Method, http.MethodDelete)

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

		res, err := client.FingerprintApi.DeleteVisitorData(ctx, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 403)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorCommon403Response)
		assert.IsType(t, errorModel, &sdk.ErrorCommon403Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

}
