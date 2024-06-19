package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/config"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEvent(t *testing.T) {
	t.Run("Returns event", func(t *testing.T) {
		mockResponse := GetMockEventResponse("../test/mocks/get_event_200.json")

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

	})

	t.Run("Returns event with errors in all signals", func(t *testing.T) {
		mockResponse := GetMockEventResponse("../test/mocks/get_event_200_all_errors.json")

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

	})

	t.Run("Returns event with unexpected fields", func(t *testing.T) {
		mockResponse := GetMockEventResponse("../test/mocks/get_event_200_extra_fields.json")

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
	})

	t.Run("Returns 403 error", func(t *testing.T) {
		mockResponse := GetEvent403ErrorMockResponse("../test/mocks/get_event_403_error.json")

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

			w.WriteHeader(403)

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

		assert.Error(t, err)
		assert.IsType(t, err, sdk.ApiError{})
		assert.NotNil(t, res)

		var apiError sdk.ApiError
		errors.As(err, &apiError)

		errorModel := err.(sdk.ApiError).Model().(*sdk.ErrorCommon403Response)

		assert.IsType(t, errorModel, &sdk.ErrorCommon403Response{})

	})

	t.Run("Returns botd too many requests error", func(t *testing.T) {
		mockResponse := GetMockEventResponse("../test/mocks/get_event_200_botd_too_many_requests_error.json")

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
		assert.Equal(t, res.Products.Botd.Error_.Code, "TooManyRequests")

	})

	t.Run("Returns identification too many requests error", func(t *testing.T) {
		mockResponse := GetMockEventResponse("../test/mocks/get_event_200_identification_too_many_requests_error.json")

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
		assert.Equal(t, res.Products.Identification.Error_.Code, "429 Too Many Requests")

	})

	t.Run("Returns response when JSON field type is not matching response schema", func(t *testing.T) {
		// Changed fields: products.identification.data.incognito bool -> int
		malformedResponse := `{
  "products": {
    "identification": {
      "data": {
        "visitorId": "Ibk1527CUFmcnjLwIs4A9",
        "requestId": "0KSh65EnVoB85JBmloQK",

        "incognito": 1,

        "linkedId": "somelinkedId",
        "tag": {},
        "time": "2019-05-21T16:40:13Z",
        "timestamp": 1582299576512,
        "url": "https://www.example.com/login",
        "ip": "61.127.217.15",
        "ipLocation": {
          "accuracyRadius": 10,
          "latitude": 49.982,
          "longitude": 36.2566,
          "postalCode": "61202",
          "timezone": "Europe/Dusseldorf",
          "city": {
            "name": "Dusseldorf"
          },
          "continent": {
            "code": "EU",
            "name": "Europe"
          },
          "country": {
            "code": "DE",
            "name": "Germany"
          },
          "subdivisions": [
            {
              "isoCode": "63",
              "name": "North Rhine-Westphalia"
            }
          ]
        },
        "browserDetails": {
          "browserName": "Chrome",
          "browserMajorVersion": "74",
          "browserFullVersion": "74.0.3729",
          "os": "Windows",
          "osVersion": "7",
          "device": "Other",
          "userAgent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64) ...."
        },
        "confidence": {
          "score": 0.97,
          "revision": "v1.1"
        },
        "visitorFound": true,
        "firstSeenAt": {
          "global": "2022-03-16T11:26:45.362Z",
          "subscription": "2022-03-16T11:31:01.101Z"
        },
        "lastSeenAt": {
          "global": "2022-03-16T11:28:34.023Z",
          "subscription": null
        }
      }
    },
    "botd": {
      "data": {
        "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36",

        "requestId": "1708102555327.NLOjmg",

        "bot": {
          "result": "notDetected"
        },
        "url": "https://www.example.com/login",
        "ip": "61.127.217.15",
        "time": "2019-05-21T16:40:13Z"
      }
    }
  }
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.NoError(t, parseErr)

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			w.Header().Set("Content-Type", "application/json")

			_, err := w.Write([]byte(malformedResponse))

			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)
		client := sdk.NewAPIClient(cfg)
		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "api_key"})

		res, _, err := client.FingerprintApi.GetEvent(ctx, "request_id")
		assert.Error(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.Products.Identification.Data.Incognito, false)
	})

	t.Run("Returns response when URL in JSON is not RFC compilant", func(t *testing.T) {
		// Changed fields: products.botd.data.url -> not RFC compliant URL
		malformedResponse := `{
  "products": {
    "identification": {
      "data": {
        "visitorId": "Ibk1527CUFmcnjLwIs4A9",
        "requestId": "0KSh65EnVoB85JBmloQK",
        "incognito": true,
        "linkedId": "somelinkedId",
        "tag": {},
        "time": "2019-05-21T16:40:13Z",
        "timestamp": 1582299576512,
        "url": "https://www.example.com/login",
        "ip": "61.127.217.15",
        "ipLocation": {
          "accuracyRadius": 10,
          "latitude": 49.982,
          "longitude": 36.2566,
          "postalCode": "61202",
          "timezone": "Europe/Dusseldorf",
          "city": {
            "name": "Dusseldorf"
          },
          "continent": {
            "code": "EU",
            "name": "Europe"
          },
          "country": {
            "code": "DE",
            "name": "Germany"
          },
          "subdivisions": [
            {
              "isoCode": "63",
              "name": "North Rhine-Westphalia"
            }
          ]
        },
        "browserDetails": {
          "browserName": "Chrome",
          "browserMajorVersion": "74",
          "browserFullVersion": "74.0.3729",
          "os": "Windows",
          "osVersion": "7",
          "device": "Other",
          "userAgent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64) ...."
        },
        "confidence": {
          "score": 0.97,
          "revision": "v1.1"
        },
        "visitorFound": true,
        "firstSeenAt": {
          "global": "2022-03-16T11:26:45.362Z",
          "subscription": "2022-03-16T11:31:01.101Z"
        },
        "lastSeenAt": {
          "global": "2022-03-16T11:28:34.023Z",
          "subscription": null
        }
      }
    },
    "botd": {
      "data": {
        "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36",

        "requestId": "1708102555327.NLOjmg",

        "bot": {
          "result": "notDetected"
        },
        "url": "https://www.example.com/{{{login",
        "ip": "61.127.217.15",
        "time": "2019-05-21T16:40:13Z"
      }
    }
  }
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.NoError(t, parseErr)

			apiKey := r.Header.Get("Auth-Api-Key")
			assert.Equal(t, apiKey, "api_key")

			w.Header().Set("Content-Type", "application/json")

			_, err := w.Write([]byte(malformedResponse))

			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		cfg := sdk.NewConfiguration()
		cfg.ChangeBasePath(ts.URL)
		client := sdk.NewAPIClient(cfg)
		ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "api_key"})

		res, _, err := client.FingerprintApi.GetEvent(ctx, "request_id")
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.Products.Botd.Data.Url, "https://www.example.com/{{{login")
	})

}
