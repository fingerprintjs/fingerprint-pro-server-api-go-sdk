package test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/stretchr/testify/assert"
)

// pathParamEndpoint is an operation that takes an ID as a URL path parameter, so that every such
// operation can be checked against the same set of values.
type pathParamEndpoint struct {
	name         string
	prefix       string
	paramName    string
	responseBody string
	call         func(ctx context.Context, client *sdk.APIClient, id string) sdk.Error
}

var pathParamEndpoints = []pathParamEndpoint{
	{
		name:         "GetEvent",
		prefix:       "/events/",
		paramName:    "requestId",
		responseBody: "{}",
		call: func(ctx context.Context, client *sdk.APIClient, id string) sdk.Error {
			_, _, err := client.FingerprintApi.GetEvent(ctx, id)
			return err
		},
	},
	{
		name:      "UpdateEvent",
		prefix:    "/events/",
		paramName: "requestId",
		call: func(ctx context.Context, client *sdk.APIClient, id string) sdk.Error {
			_, err := client.FingerprintApi.UpdateEvent(ctx, sdk.EventsUpdateRequest{LinkedId: "linked_id"}, id)
			return err
		},
	},
	{
		name:         "GetVisits",
		prefix:       "/visitors/",
		paramName:    "visitorId",
		responseBody: "{}",
		call: func(ctx context.Context, client *sdk.APIClient, id string) sdk.Error {
			_, _, err := client.FingerprintApi.GetVisits(ctx, id, nil)
			return err
		},
	},
	{
		name:      "DeleteVisitorData",
		prefix:    "/visitors/",
		paramName: "visitorId",
		call: func(ctx context.Context, client *sdk.APIClient, id string) sdk.Error {
			_, err := client.FingerprintApi.DeleteVisitorData(ctx, id)
			return err
		},
	},
}

type capturedPathParamRequest struct {
	received      bool
	requestTarget string
	escapedPath   string
	decodedPath   string
}

// callWithPathParam performs endpoint.call against a local server and returns what that server saw
// along with the error the SDK returned.
func callWithPathParam(t *testing.T, endpoint pathParamEndpoint, id string) (capturedPathParamRequest, sdk.Error) {
	t.Helper()

	var captured capturedPathParamRequest

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.RequestURI is the literal request target, before any parsing or normalization.
		captured.received = true
		captured.requestTarget = strings.SplitN(r.RequestURI, "?", 2)[0]
		captured.escapedPath = r.URL.EscapedPath()
		captured.decodedPath = r.URL.Path

		apiKey := r.Header.Get("Auth-Api-Key")
		assert.Equal(t, "api_key", apiKey)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if endpoint.responseBody != "" {
			_, err := w.Write([]byte(endpoint.responseBody))
			assert.Nil(t, err)
		}
	}))
	defer ts.Close()

	cfg := sdk.NewConfiguration()
	cfg.ChangeBasePath(ts.URL)

	client := sdk.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{Key: "api_key"})

	// The error is bound before returning: operand evaluation order within a return statement is
	// only specified for function calls, so reading `captured` inline would be relying on luck.
	err := endpoint.call(ctx, client, id)

	return captured, err
}

// TestPathParamIsEncodedAsSingleOpaqueSegment verifies a malformed path parameter is percent-encoded
// into a single opaque path segment rather than being interpreted as part of the path structure,
// across every operation that takes an ID in the path.
func TestPathParamIsEncodedAsSingleOpaqueSegment(t *testing.T) {
	cases := []struct {
		name    string
		value   string
		encoded string
	}{
		{name: "path traversal", value: "../events", encoded: "..%2Fevents"},
		{name: "nested path traversal", value: "../../events", encoded: "..%2F..%2Fevents"},
		{name: "leading slash", value: "/events/123", encoded: "%2Fevents%2F123"},
		{name: "absolute url", value: "https://domain.tld/evil", encoded: "https:%2F%2Fdomain.tld%2Fevil"},
		{name: "query injection", value: "123?limit=1", encoded: "123%3Flimit=1"},
		{name: "fragment injection", value: "123#fragment", encoded: "123%23fragment"},
		{name: "whitespace", value: "hello world", encoded: "hello%20world"},
		{name: "empty", value: "", encoded: ""},
	}

	for _, endpoint := range pathParamEndpoints {
		for _, c := range cases {
			t.Run(endpoint.name+": "+c.name, func(t *testing.T) {
				captured, err := callWithPathParam(t, endpoint, c.value)

				assert.Nil(t, err)

				expectedPath := endpoint.prefix + c.encoded

				assert.Equal(t, expectedPath, captured.requestTarget)
				assert.Equal(t, expectedPath, captured.escapedPath)

				// Decoding the path yields exactly the ID that was passed in, so the encoding
				// never changes which resource is addressed.
				assert.Equal(t, endpoint.prefix+c.value, captured.decodedPath)
			})
		}
	}
}

// TestInvalidPathParamReturnsErrorWithoutSendingRequest covers a path parameter of exactly "." or
// "..", an RFC 3986 dot-segment. Such a value does not address a resource, so the SDK rejects it up
// front instead of sending a request that URL normalizers would resolve to a different endpoint.
func TestInvalidPathParamReturnsErrorWithoutSendingRequest(t *testing.T) {
	for _, endpoint := range pathParamEndpoints {
		for _, value := range []string{".", ".."} {
			t.Run(endpoint.name+": "+value, func(t *testing.T) {
				captured, err := callWithPathParam(t, endpoint, value)

				assert.Error(t, err)
				assert.False(t, captured.received)

				// errors.As is how consumers are expected to detect this error.
				var invalidArgumentError *sdk.InvalidArgumentError
				assert.True(t, errors.As(err, &invalidArgumentError))

				assert.Equal(t, sdk.ErrorCode_REQUEST_CANNOT_BE_PARSED, invalidArgumentError.Code())
				assert.Equal(t, endpoint.paramName, invalidArgumentError.Parameter())
				assert.Equal(t, value, invalidArgumentError.Value())
				assert.Nil(t, invalidArgumentError.Body())
				assert.Nil(t, invalidArgumentError.Model())

				// The message names both the offending parameter and its value.
				assert.Contains(t, invalidArgumentError.Error(), endpoint.paramName)
				assert.Contains(t, invalidArgumentError.Error(), value)
			})
		}
	}
}

// TestPathParamLeavesValidIdsUntouched makes sure the encoding does not mangle the IDs issued by
// Fingerprint. Only a segment made up solely of dots is rejected; anything else containing a dot
// must pass through unencoded, since "." is otherwise an unreserved character.
func TestPathParamLeavesValidIdsUntouched(t *testing.T) {
	values := []string{
		"1708102555327.NLOjmg",
		"Ibk1527CUFmcnjLwIs4A9",
		"0KSh65EnVoB85JBmloQK",
		"...",
		"..a",
		"a..",
		"-_.~",
	}

	for _, endpoint := range pathParamEndpoints {
		for _, value := range values {
			t.Run(endpoint.name+": "+value, func(t *testing.T) {
				captured, err := callWithPathParam(t, endpoint, value)

				assert.Nil(t, err)
				assert.Equal(t, endpoint.prefix+value, captured.requestTarget)
				assert.Equal(t, endpoint.prefix+value, captured.decodedPath)
			})
		}
	}
}
