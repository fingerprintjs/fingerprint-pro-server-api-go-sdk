package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/config"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var trueValue = true

func TestUpdateEvent(t *testing.T) {
	t.Run("Returns 200 on success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  &trueValue,
		}, "123")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("Update with empty body", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, "{}", string(body))

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{}, "123")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("Update with empty tag", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, "{}", string(body))

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

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

		var tag sdk.ModelMap

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{
			Tag: &tag,
		}, "123")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("Update with empty tag struct", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, "{}", string(body))

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

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

		tag := sdk.ModelMap{}

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{
			Tag: &tag,
		}, "123")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("Update with just suspect=false", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, "{\"suspect\":false}", string(body))

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

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

		suspect := false

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{
			Suspect: &suspect,
		}, "123")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("Returns ErrorResponse on 404", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorResponse]("../test/mocks/errors/404_visitor_not_found.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  &trueValue,
		}, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 404)

		errorModel := err.(*sdk.ApiError).Model().(*sdk.ErrorResponse)
		assert.IsType(t, errorModel, &sdk.ErrorResponse{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorResponse on 400", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorResponse]("../test/mocks/errors/400_request_body_invalid.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  &trueValue,
		}, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 400)

		errorModel := err.(*sdk.ApiError).Model().(*sdk.ErrorResponse)
		assert.IsType(t, errorModel, &sdk.ErrorResponse{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorResponse on 403", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorResponse]("../test/mocks/errors/403_token_required.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  &trueValue,
		}, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 403)

		errorModel := err.(*sdk.ApiError).Model().(*sdk.ErrorResponse)
		assert.IsType(t, errorModel, &sdk.ErrorResponse{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorResponse on 409", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorResponse]("../test/mocks/errors/409_state_not_ready.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPut)

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(409)
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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  &trueValue,
		}, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 409)

		errorModel := err.(*sdk.ApiError).Model().(*sdk.ErrorResponse)
		assert.IsType(t, errorModel, &sdk.ErrorResponse{})
		assert.Equal(t, errorModel, &mockResponse)
	})
}
