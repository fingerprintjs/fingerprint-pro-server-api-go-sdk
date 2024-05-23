package sdk

import "net/http"

func doRequest[T any](handler func() (T, *http.Response, error)) (T, *http.Response, error) {
	result, httpResponse, err := handler()

	handlePotentialManyRequestsResponse(httpResponse, err)

	return result, httpResponse, err
}
