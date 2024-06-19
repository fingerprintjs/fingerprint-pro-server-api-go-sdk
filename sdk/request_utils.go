package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

func handlePotentialManyRequestsResponse(httpResponse *http.Response, error error) error {
	if error == nil {
		return error
	}

	var e ApiError

	if errors.As(error, &e) {
		if model, ok := e.model.(*TooManyRequestsResponse); ok {
			retryAfter := getRetryAfterFromHeader(httpResponse)

			return &TooManyRequestsError{
				error:      model.Error_,
				retryAfter: retryAfter,
				body:       e.body,
				model:      e.model,
			}
		}
	}

	return error
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

func getQueryWithIntegrationInfo(url *url.URL) url.Values {
	query := url.Query()

	query.Add("ii", IntegrationInfo)

	return query
}

func handleErrorResponse[T any](body []byte, httpResponse *http.Response, definition requestDefinition, responseModel T) (T, *http.Response, error) {
	genericError := ApiError{
		body:  body,
		error: httpResponse.Status,
	}

	modelFactory := definition.StatusCodeResultsFactoryMap[httpResponse.StatusCode]

	if modelFactory != nil {
		model := modelFactory()

		err := json.Unmarshal(body, &model)

		if err != nil {
			genericError.error = err.Error()

			return responseModel, httpResponse, genericError
		}

		genericError.model = model
	}

	return responseModel, httpResponse, genericError
}

func isResponseOk(httpResponse *http.Response) bool {
	return 199 < httpResponse.StatusCode && httpResponse.StatusCode < 300
}

func handleAuth(ctx context.Context, request *http.Request) {
	if ctx != nil {
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
}
