package sdk

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type FingerprintApiService struct {
	cfg *Configuration
}

func (f *FingerprintApiService) makeRequest(ctx context.Context, requestUrl *url.URL, method string) (*http.Request, error) {
	request, err := http.NewRequest(method, requestUrl.String(), nil)

	if err != nil {
		return request, err
	}

	handleAuth(ctx, request)

	request.Header.Add("User-Agent", f.cfg.UserAgent)

	return request, nil
}

func (f *FingerprintApiService) path(definition requestDefinition, params ...string) string {
	return f.cfg.basePath + definition.Path(params...)
}

func (f *FingerprintApiService) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	_, response, err := doRequest(func() (any, *http.Response, error) {
		definition := createDeleteVisitorDataDefinition()

		path := f.path(definition, visitorId)

		requestUrl, err := url.Parse(path)

		if err != nil {
			return nil, nil, err
		}

		query := getQueryWithIntegrationInfo(requestUrl)
		requestUrl.RawQuery = query.Encode()

		request, err := f.makeRequest(ctx, requestUrl, http.MethodDelete)

		if err != nil {
			return nil, nil, err
		}

		httpResponse, err := f.cfg.HTTPClient.Do(request)

		if err != nil {
			return nil, httpResponse, err
		}

		if isResponseOk(httpResponse) {
			return nil, httpResponse, err
		}

		body, err := io.ReadAll(httpResponse.Body)

		if err != nil {
			return nil, httpResponse, err
		}

		return handleErrorResponse[any](body, httpResponse, definition, nil)
	})

	return response, err
}

func (f *FingerprintApiService) GetEvent(ctx context.Context, requestId string) (EventResponse, *http.Response, error) {
	return doRequest(func() (EventResponse, *http.Response, error) {
		definition := createGetEventDefinition()

		var eventResponse EventResponse

		path := f.path(definition, requestId)

		requestUrl, err := url.Parse(path)

		if err != nil {
			return eventResponse, nil, err
		}

		query := getQueryWithIntegrationInfo(requestUrl)
		requestUrl.RawQuery = query.Encode()

		request, err := f.makeRequest(ctx, requestUrl, http.MethodGet)

		if err != nil {
			return eventResponse, nil, err
		}

		httpResponse, err := f.cfg.HTTPClient.Do(request)

		if err != nil {
			return eventResponse, httpResponse, err
		}

		body, err := io.ReadAll(httpResponse.Body)

		if err != nil {
			return eventResponse, httpResponse, err
		}

		if isResponseOk(httpResponse) {
			err = json.Unmarshal(body, &eventResponse)

			return eventResponse, httpResponse, err
		}

		return handleErrorResponse(body, httpResponse, definition, eventResponse)
	})
}

func (f *FingerprintApiService) GetVisits(ctx context.Context, visitorId string, opts *FingerprintApiGetVisitsOpts) (Response, *http.Response, error) {
	return doRequest(func() (Response, *http.Response, error) {
		definition := createGetVisitsDefinition()

		var response Response

		path := f.path(definition, visitorId)

		requestUrl, err := url.Parse(path)

		if err != nil {
			return response, nil, err
		}

		query := getQueryWithIntegrationInfo(requestUrl)

		if opts != nil {
			addStructToURLQuery(&query, *opts)
		}

		requestUrl.RawQuery = query.Encode()

		request, err := f.makeRequest(ctx, requestUrl, http.MethodGet)

		if err != nil {
			return response, nil, err
		}

		httpResponse, err := f.cfg.HTTPClient.Do(request)

		if err != nil {
			return response, httpResponse, err
		}

		body, err := io.ReadAll(httpResponse.Body)

		if err != nil {
			return response, httpResponse, err
		}

		if isResponseOk(httpResponse) {
			err = json.Unmarshal(body, &response)

			return response, httpResponse, err
		}

		return handleErrorResponse(body, httpResponse, definition, response)
	})

}
