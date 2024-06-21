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
	definition    requestDefinition
	pathParams    []string
	method        string
	requestParams interface{}
	context       context.Context
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

func (f *FingerprintApiService) doRequest(apiRequest apiRequest, result interface{}) (*http.Response, error) {
	path := f.getPath(apiRequest.definition, apiRequest.pathParams...)
	requestUrl, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	query := getQueryWithIntegrationInfo(requestUrl)
	if apiRequest.requestParams != nil {
		addStructToURLQuery(&query, apiRequest.requestParams)
	}
	requestUrl.RawQuery = query.Encode()

	request, err := f.prepareRequest(apiRequest.context, requestUrl, http.MethodDelete)
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

func (f *FingerprintApiService) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	var request = apiRequest{
		definition:    createDeleteVisitorDataDefinition(),
		pathParams:    []string{visitorId},
		requestParams: nil,
		method:        http.MethodDelete,
		context:       ctx,
	}

	return f.doRequest(request, nil)
}

func (f *FingerprintApiService) GetEvent(ctx context.Context, requestId string) (EventResponse, *http.Response, error) {
	definition := createGetEventDefinition()

	var eventResponse EventResponse

	var request = apiRequest{
		definition:    definition,
		pathParams:    []string{requestId},
		requestParams: nil,
		method:        http.MethodGet,
		context:       ctx,
	}

	response, err := f.doRequest(request, &eventResponse)

	return eventResponse, response, err
}

func (f *FingerprintApiService) GetVisits(ctx context.Context, visitorId string, opts *FingerprintApiGetVisitsOpts) (Response, *http.Response, error) {
	definition := createGetVisitsDefinition()

	var response Response

	var apiRequest = apiRequest{
		definition:    definition,
		pathParams:    []string{visitorId},
		requestParams: opts,
		method:        http.MethodGet,
		context:       ctx,
	}

	httpResponse, err := f.doRequest(apiRequest, &response)

	return response, httpResponse, err
}
