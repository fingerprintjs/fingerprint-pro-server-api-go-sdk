package sdk

import (
	"bytes"
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
	body        io.Reader
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

func (f *FingerprintApiService) GetEvent(ctx context.Context, requestId string) (EventsGetResponse, *http.Response, error) {
	request := apiRequest{
		definition:  createGetEventDefinition(),
		pathParams:  []string{requestId},
		queryParams: nil,
		method:      http.MethodGet,
	}

	var eventResponse EventsGetResponse
	response, err := f.doRequest(ctx, request, &eventResponse)

	return eventResponse, response, err
}

func (f *FingerprintApiService) UpdateEvent(ctx context.Context, body EventsUpdateRequest, requestId string) (*http.Response, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request := apiRequest{
		definition: createUpdateEventDefinition(),
		pathParams: []string{requestId},
		body:       bytes.NewBuffer(bodyBytes),
		method:     http.MethodPut,
	}

	httpResponse, err := f.doRequest(ctx, request, nil)

	return httpResponse, err
}

func (f *FingerprintApiService) GetVisits(ctx context.Context, visitorId string, opts *FingerprintApiGetVisitsOpts) (VisitorsGetResponse, *http.Response, error) {
	request := apiRequest{
		definition:  createGetVisitsDefinition(),
		pathParams:  []string{visitorId},
		queryParams: opts.ToUrlValuesMap(),
		method:      http.MethodGet,
	}

	var response VisitorsGetResponse
	httpResponse, err := f.doRequest(ctx, request, &response)

	return response, httpResponse, err
}

func (f *FingerprintApiService) GetRelatedVisitors(ctx context.Context, visitorId string) (RelatedVisitorsResponse, *http.Response, error) {
	request := apiRequest{
		definition: createGetRelatedVisitorsDefinition(),
		queryParams: map[string]any{
			"visitor_id": visitorId,
		},
		method: http.MethodGet,
	}

	var response RelatedVisitorsResponse
	httpResponse, err := f.doRequest(ctx, request, &response)

	return response, httpResponse, err
}

func (f *FingerprintApiService) prepareRequest(ctx context.Context, requestUrl *url.URL, method string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, requestUrl.String(), body)

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

	request, err := f.prepareRequest(ctx, requestUrl, apiRequest.method, apiRequest.body)
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
