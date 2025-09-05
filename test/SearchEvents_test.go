package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/config"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchEvents(t *testing.T) {
	t.Run("Search with just limit", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.SearchEventsResponse]("../test/mocks/get_event_search_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			assert.Equal(t, "/events/search", r.URL.Path)
			assert.Equal(t, "2", r.URL.Query().Get("limit"))
			assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 2)

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

		res, _, err := client.FingerprintApi.SearchEvents(ctx, 2, nil)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, mockResponse)
		assert.IsType(t, sdk.SearchEventsResponse{}, res)
	})

	t.Run("Search with partial params", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.SearchEventsResponse]("../test/mocks/get_event_search_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			configFile := config.ReadConfig("../config.json")
			query := r.URL.Query()
			integrationInfo := query.Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			assert.Equal(t, "/events/search", r.URL.Path)
			assert.Equal(t, "2", query.Get("limit"), "limit")
			assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 7)
			assert.Equal(t, "", query.Get("suspect"), "suspect")
			assert.False(t, query.Has("suspect"), "has suspect")
			assert.Equal(t, "", query.Get("bot"), "bot")
			assert.False(t, query.Has("bot"), "has bot")
			assert.Equal(t, "10", query.Get("end"), "end")
			assert.Equal(t, "5", query.Get("start"), "start")
			assert.Equal(t, "", query.Get("ip_address"), "ip_address")
			assert.False(t, query.Has("ip_address"), "has ip_address")
			assert.Equal(t, "linked_id", query.Get("linked_id"), "linked_id")
			assert.Equal(t, "false", query.Get("reverse"), "reverse")
			assert.Equal(t, "XIkiQhRyp7edU9SA0jBb", query.Get("visitor_id"), "visitor_id")

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

		var end int64 = 10
		var start int64 = 5
		reverse := false
		linkedId := "linked_id"
		visitorId := "XIkiQhRyp7edU9SA0jBb"
		opts := sdk.FingerprintApiSearchEventsOpts{
			End:       &end,
			Start:     &start,
			LinkedId:  &linkedId,
			Reverse:   &reverse,
			VisitorId: &visitorId,
		}
		res, _, err := client.FingerprintApi.SearchEvents(ctx, 2, &opts)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, mockResponse)
	})

	t.Run("Search with all params", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.SearchEventsResponse]("../test/mocks/get_event_search_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			configFile := config.ReadConfig("../config.json")
			query := r.URL.Query()
			integrationInfo := query.Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			assert.Equal(t, "/events/search", r.URL.Path)
			assert.Equal(t, "2", query.Get("limit"), "limit")
			assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 36, "expected all parameters in query")

			// Existing params
			assert.Equal(t, "true", query.Get("suspect"), "suspect")
			assert.True(t, query.Has("suspect"), "has suspect")
			assert.Equal(t, "bot", query.Get("bot"), "bot")
			assert.True(t, query.Has("bot"), "has bot")
			assert.Equal(t, "10", query.Get("end"), "end")
			assert.Equal(t, "5", query.Get("start"), "start")
			assert.Equal(t, "127.0.0.1", query.Get("ip_address"), "ip_address")
			assert.True(t, query.Has("ip_address"), "has ip_address")
			assert.Equal(t, "linked_id", query.Get("linked_id"), "linked_id")
			assert.Equal(t, "false", query.Get("reverse"), "reverse")
			assert.Equal(t, "XIkiQhRyp7edU9SA0jBb", query.Get("visitor_id"), "visitor_id")

			// New params
			assert.Equal(t, "true", query.Get("vpn"), "vpn")
			assert.Equal(t, "true", query.Get("virtual_machine"), "virtual_machine")
			assert.Equal(t, "true", query.Get("tampering"), "tampering")
			assert.Equal(t, "true", query.Get("anti_detect_browser"), "anti_detect_browser")
			assert.Equal(t, "true", query.Get("incognito"), "incognito")
			assert.Equal(t, "true", query.Get("privacy_settings"), "privacy_settings")
			assert.Equal(t, "true", query.Get("jailbroken"), "jailbroken")
			assert.Equal(t, "true", query.Get("frida"), "frida")
			assert.Equal(t, "true", query.Get("factory_reset"), "factory_reset")
			assert.Equal(t, "true", query.Get("cloned_app"), "cloned_app")
			assert.Equal(t, "true", query.Get("emulator"), "emulator")
			assert.Equal(t, "true", query.Get("root_apps"), "root_apps")
			assert.Equal(t, "high", query.Get("vpn_confidence"), "vpn_confidence")
			assert.Equal(t, "85.5", query.Get("min_suspect_score"), "min_suspect_score")
			assert.Equal(t, "true", query.Get("ip_blocklist"), "ip_blocklist")
			assert.Equal(t, "true", query.Get("datacenter"), "datacenter")
			assert.Equal(t, "true", query.Get("developer_tools"), "developer_tools")
			assert.Equal(t, "true", query.Get("location_spoofing"), "location_spoofing")
			assert.Equal(t, "true", query.Get("mitm_attack"), "mitm_attack")
			assert.Equal(t, "true", query.Get("proxy"), "proxy")
			assert.Equal(t, "testSdkVersion", query.Get("sdk_version"), "sdk_version")
			assert.Equal(t, "testSdkPlatform", query.Get("sdk_platform"), "sdk_platform")
			assert.Equal(t, "testProximityId", query.Get("proximity_id"), "proximity_id")
			assert.Equal(t, "10", query.Get("proximity_precision_radius"), "proximity_precision_radius")

			envs := r.URL.Query()["environment"]
			assert.Equal(t, []string{"env1", "env2"}, envs, "environment")

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

		var (
			end                      int64   = 10
			start                    int64   = 5
			suspect                          = true
			reverse                          = false
			bot                              = "bot"
			ipAddress                        = "127.0.0.1"
			linkedId                         = "linked_id"
			visitorId                        = "XIkiQhRyp7edU9SA0jBb"
			vpn                              = true
			virtualMachine                   = true
			tampering                        = true
			antiDetectBrowser                = true
			incognito                        = true
			privacySettings                  = true
			jailbroken                       = true
			frida                            = true
			factoryReset                     = true
			clonedApp                        = true
			emulator                         = true
			rootApps                         = true
			vpnConfidence                    = "high"
			minSuspectScore          float32 = 85.5
			ipBlocklist                      = true
			datacenter                       = true
			developerTools                   = true
			locationSpoofing                 = true
			mitmAttack                       = true
			proxy                            = true
			sdkVersion                       = "testSdkVersion"
			sdkPlatform                      = "testSdkPlatform"
			environment                      = []string{"env1", "env2"}
			proximityId                      = "testProximityId"
			proximityPrecisionRadius int32   = 10
		)

		opts := sdk.FingerprintApiSearchEventsOpts{
			Suspect:                  &suspect,
			Bot:                      &bot,
			End:                      &end,
			Start:                    &start,
			IpAddress:                &ipAddress,
			LinkedId:                 &linkedId,
			Reverse:                  &reverse,
			VisitorId:                &visitorId,
			Vpn:                      &vpn,
			VirtualMachine:           &virtualMachine,
			Tampering:                &tampering,
			AntiDetectBrowser:        &antiDetectBrowser,
			Incognito:                &incognito,
			PrivacySettings:          &privacySettings,
			Jailbroken:               &jailbroken,
			Frida:                    &frida,
			FactoryReset:             &factoryReset,
			ClonedApp:                &clonedApp,
			Emulator:                 &emulator,
			RootApps:                 &rootApps,
			VpnConfidence:            &vpnConfidence,
			MinSuspectScore:          &minSuspectScore,
			IpBlocklist:              &ipBlocklist,
			Datacenter:               &datacenter,
			DeveloperTools:           &developerTools,
			LocationSpoofing:         &locationSpoofing,
			MitmAttack:               &mitmAttack,
			Proxy:                    &proxy,
			SdkVersion:               &sdkVersion,
			SdkPlatform:              &sdkPlatform,
			Environment:              &environment,
			ProximityId:              &proximityId,
			ProximityPrecisionRadius: &proximityPrecisionRadius,
		}
		res, _, err := client.FingerprintApi.SearchEvents(ctx, 2, &opts)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res, mockResponse)
		assert.IsType(t, sdk.SearchEventsResponse{}, res)
	})
	t.Run("Returns ErrorResponse on 400", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorResponse]("../test/mocks/errors/400_ip_address_invalid.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			assert.Equal(t, "/events/search", r.URL.Path)
			assert.Equal(t, "2", r.URL.Query().Get("limit"))
			assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 2)

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

		_, res, err := client.FingerprintApi.SearchEvents(ctx, 2, nil)
		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

		errorModel := err.(*sdk.ApiError).Model().(*sdk.ErrorResponse)
		assert.IsType(t, errorModel, &sdk.ErrorResponse{})
		assert.Equal(t, errorModel, &mockResponse)
	})

	t.Run("Returns ErrorResponse on 403", func(t *testing.T) {
		mockResponse := GetMockResponse[sdk.ErrorResponse]("../test/mocks/errors/403_feature_not_enabled.json")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			configFile := config.ReadConfig("../config.json")
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", configFile.PackageVersion))

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			assert.Equal(t, "/events/search", r.URL.Path)
			assert.Equal(t, "2", r.URL.Query().Get("limit"))
			assert.Len(t, strings.Split(r.URL.RawQuery, "&"), 2)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
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

		_, res, err := client.FingerprintApi.SearchEvents(ctx, 2, nil)
		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, http.StatusForbidden, res.StatusCode)

		errorModel := err.(*sdk.ApiError).Model().(*sdk.ErrorResponse)
		assert.IsType(t, errorModel, &sdk.ErrorResponse{})
		assert.Equal(t, errorModel, &mockResponse)
	})

}
