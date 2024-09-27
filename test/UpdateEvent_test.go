package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/config"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  true,
		}, "123")

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("Returns ErrorEvent404Response on 404", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorEvent404Response]("../test/mocks/update_event_404_error.json")

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  true,
		}, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 404)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorEvent404Response)
		assert.IsType(t, errorModel, &sdk.ErrorEvent404Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorUpdateEvent400Response on 400", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorUpdateEvent400Response]("../test/mocks/update_event_400_error.json")

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  true,
		}, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 400)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorUpdateEvent400Response)
		assert.IsType(t, errorModel, &sdk.ErrorUpdateEvent400Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorCommon403Response on 403", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorCommon403Response]("../test/mocks/update_event_403_error.json")

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  true,
		}, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 403)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorCommon403Response)
		assert.IsType(t, errorModel, &sdk.ErrorCommon403Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorUpdateEvent409Response on 409", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorUpdateEvent409Response]("../test/mocks/update_event_409_error.json")

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

		res, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventUpdateRequest{
			LinkedId: "linked_id",
			Tag:      nil,
			Suspect:  true,
		}, "123")

		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 409)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorUpdateEvent409Response)
		assert.IsType(t, errorModel, &sdk.ErrorUpdateEvent409Response{})
		assert.Equal(t, errorModel, &mockResponse)
	})
}
