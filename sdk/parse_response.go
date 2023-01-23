package sdk

import (
	"net/http"
	"strconv"
)

func ParseResponse(httpResponse *http.Response, error interface{}) {
	if error == nil {
		return
	}

	switch e := error.(type) {
	case *GenericSwaggerError:
		switch model := e.model.(type) {
		case ManyRequestsResponse:
			retryAfter := getRetryAfterFromHeader(httpResponse)

			e.model = ManyRequestsResponse{
				RetryAfter: retryAfter,
				Error_:     model.Error_,
			}
		}

	}

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
