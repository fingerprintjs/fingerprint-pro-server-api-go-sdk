package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

func handlePotentialTooManyRequestsResponse(httpResponse *http.Response, err Error) Error {
	if err == nil {
		return nil
	}

	var e *ApiError

	if httpResponse.StatusCode != http.StatusTooManyRequests {
		return WrapWithApiError(err)
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

			code := TOOMANYREQUESTS429
			msg := e.error
			if model.Error_ != nil {
				if model.Error_.Message != "" {
					msg = model.Error_.Message
				}
				if model.Error_.Code != nil {
					code = *model.Error_.Code
				}
			}

			return &TooManyRequestsError{
				error:      msg,
				code:       code,
				retryAfter: retryAfter,
				body:       e.body,
				model:      e.model,
			}
		}

		retryAfter := getRetryAfterFromHeader(httpResponse)

		return &TooManyRequestsError{
			error:      e.error,
			code:       TOOMANYREQUESTS429,
			retryAfter: retryAfter,
			body:       e.body,
			model:      e.model,
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

func handleErrorResponse(body []byte, httpResponse *http.Response, definition requestDefinition) *ApiError {
	apiError := ApiError{
		body:  body,
		error: httpResponse.Status,
		code:  FAILED,
	}

	modelFactory := definition.StatusCodeResultsFactoryMap[httpResponse.StatusCode]

	if modelFactory != nil {
		model := modelFactory()

		err := json.Unmarshal(body, model)

		if err != nil {
			var plain ErrorPlainResponse
			if errPlain := json.Unmarshal(body, &plain); errPlain == nil {
				apiError.error = plain.Error_
				apiError.model = &plain
				return &apiError
			}
			apiError.error = err.Error()

			return &apiError
		}

		switch v := model.(type) {
		case *ErrorResponse:
			if v.Error_ != nil {
				if v.Error_.Code != nil {
					apiError.code = *v.Error_.Code
				}
				apiError.error = v.Error_.Message
			}

		case *ErrorPlainResponse:
			apiError.error = v.Error_
		}

		apiError.model = model
	}

	return &apiError
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
