package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

func handlePotentialTooManyRequestsResponse(httpResponse *http.Response, err error) error {
	if err == nil {
		return err
	}

	var e ApiError

	if httpResponse.StatusCode != http.StatusTooManyRequests {
		return err
	}

	if errors.As(err, &e) {
		if model, ok := e.model.(*ErrorPlainResponse); ok {
			retryAfter := getRetryAfterFromHeader(httpResponse)

			return &TooManyRequestsError{
				error:      model.Error_,
				code:       TOOMANYREQUESTS429,
				retryAfter: retryAfter,
				body:       e.body,
				model:      e.model,
			}
		}

		if model, ok := e.model.(*ErrorResponse); ok {
			retryAfter := getRetryAfterFromHeader(httpResponse)

			return &TooManyRequestsError{
				error:      model.Error_.Message,
				code:       *model.Error_.Code,
				retryAfter: retryAfter,
				body:       e.body,
				model:      e.model,
			}
		}
	}

	return err
}

func getRetryAfterFromHeader(httpResponse *http.Response) int64 {
	header := httpResponse.Header.Get("Retry-After")

	if header != "" {
		parsed, err := strconv.ParseInt(header, 10, 32)

		if err == nil {
			return parsed
		}
	}

	return 0
}

func addIntegrationInfoToQuery(query *url.Values) {
	query.Add("ii", IntegrationInfo)
}

// UNKNOWN_ERROR_CODE is a fallback error code for unknown errors, usually used when ErrorPlainResponse is returned
var UNKNOWN_ERROR_CODE ErrorCode = "UNKNOWN"

func handleErrorResponse(body []byte, httpResponse *http.Response, definition requestDefinition) (*http.Response, error) {
	apiError := ApiError{
		body:  body,
		error: httpResponse.Status,
		code:  UNKNOWN_ERROR_CODE,
	}

	modelFactory := definition.StatusCodeResultsFactoryMap[httpResponse.StatusCode]

	if modelFactory != nil {
		model := modelFactory()

		err := json.Unmarshal(body, &model)

		if err != nil {
			apiError.error = err.Error()

			return httpResponse, apiError
		}

		switch v := model.(type) {
		case *ErrorResponse:
			apiError.code = *v.Error_.Code
			apiError.error = v.Error_.Message

		case *ErrorPlainResponse:
			apiError.error = v.Error_
		}

		apiError.model = model
	}

	return httpResponse, apiError
}

func isResponseOk(httpResponse *http.Response) bool {
	return 199 < httpResponse.StatusCode && httpResponse.StatusCode < 300
}

func handleAuth(ctx context.Context, request *http.Request) {
	if ctx == nil {
		return
	}

	if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
		var key string
		if auth.Prefix != "" {
			key = auth.Prefix + " " + auth.Key
		} else {
			key = auth.Key
		}

		request.Header.Add("Auth-API-Key", key)
	}
}
