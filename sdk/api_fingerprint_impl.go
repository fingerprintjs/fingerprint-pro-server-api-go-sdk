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

type apiRequest struct {
	definition  requestDefinition
	pathParams  []string
	method      string
	queryParams map[string]any
}

func (f *FingerprintApiService) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	request := apiRequest{
		definition:  createDeleteVisitorDataDefinition(),
		pathParams:  []string{visitorId},
		queryParams: nil,
		method:      http.MethodDelete,
	}

	return f.doRequest(ctx, request, nil)
}

func (f *FingerprintApiService) GetEvent(ctx context.Context, requestId string) (EventResponse, *http.Response, error) {
	request := apiRequest{
		definition:  createGetEventDefinition(),
		pathParams:  []string{requestId},
		queryParams: nil,
		method:      http.MethodGet,
	}

	var eventResponse EventResponse
	response, err := f.doRequest(ctx, request, &eventResponse)

	return eventResponse, response, err
}

func (f *FingerprintApiService) GetVisits(ctx context.Context, visitorId string, opts *FingerprintApiGetVisitsOpts) (Response, *http.Response, error) {
	request := apiRequest{
		definition:  createGetVisitsDefinition(),
		pathParams:  []string{visitorId},
		queryParams: opts.ToUrlValuesMap(),
		method:      http.MethodGet,
	}

	var response Response
	httpResponse, err := f.doRequest(ctx, request, &response)

	return response, httpResponse, err
}

func (f *FingerprintApiService) prepareRequest(ctx context.Context, requestUrl *url.URL, method string) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, requestUrl.String(), nil)

	if err != nil {
		return request, err
	}

	handleAuth(ctx, request)

	request.Header.Add("User-Agent", f.cfg.UserAgent)

	return request, nil
}

func (f *FingerprintApiService) getPath(definition requestDefinition, params ...string) string {
	return f.cfg.basePath + definition.GetPath(params...)
}

func (f *FingerprintApiService) doRequest(ctx context.Context, apiRequest apiRequest, result any) (*http.Response, error) {
	path := f.getPath(apiRequest.definition, apiRequest.pathParams...)
	requestUrl, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	query := requestUrl.Query()

	if apiRequest.queryParams != nil {
		addMapToUrlValues(apiRequest.queryParams, &query)
	}

	addIntegrationInfoToQuery(&query)

	requestUrl.RawQuery = query.Encode()

	request, err := f.prepareRequest(ctx, requestUrl, apiRequest.method)
	if err != nil {
		return nil, err
	}

	httpResponse, err := f.cfg.HTTPClient.Do(request)
	if err != nil {
		return httpResponse, err
	}

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return httpResponse, err
	}

	if isResponseOk(httpResponse) {
		if result != nil {
			err = json.Unmarshal(body, &result)
		}

		return httpResponse, err
	}

	_, err = handleErrorResponse(body, httpResponse, apiRequest.definition)

	return httpResponse, handlePotentialTooManyRequestsResponse(httpResponse, err)
}
