package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

func (f *FingerprintApiService) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, Error) {
	request := apiRequest{
		definition:  createDeleteVisitorDataDefinition(),
		pathParams:  []string{visitorId},
		queryParams: nil,
		method:      http.MethodDelete,
	}

	return f.doRequest(ctx, request, nil)
}

func (f *FingerprintApiService) GetEvent(ctx context.Context, requestId string) (EventsGetResponse, *http.Response, Error) {
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

func (f *FingerprintApiService) UpdateEvent(ctx context.Context, body EventsUpdateRequest, requestId string) (*http.Response, Error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, WrapWithApiError(err)
	}

	request := apiRequest{
		definition: createUpdateEventDefinition(),
		pathParams: []string{requestId},
		body:       bytes.NewBuffer(bodyBytes),
		method:     http.MethodPut,
	}

	httpResponse, apiError := f.doRequest(ctx, request, nil)

	return httpResponse, apiError
}

func (f *FingerprintApiService) GetVisits(ctx context.Context, visitorId string, opts *FingerprintApiGetVisitsOpts) (VisitorsGetResponse, *http.Response, Error) {
	request := apiRequest{
		definition:  createGetVisitsDefinition(),
		pathParams:  []string{visitorId},
		queryParams: opts.ToQueryParams(),
		method:      http.MethodGet,
	}

	var response VisitorsGetResponse
	httpResponse, err := f.doRequest(ctx, request, &response)

	return response, httpResponse, err
}

func (f *FingerprintApiService) SearchEvents(ctx context.Context, limit int32, opts *FingerprintApiSearchEventsOpts) (SearchEventsResponse, *http.Response, Error) {
	queryParams := opts.ToQueryParams()
	queryParams["limit"] = strconv.Itoa(int(limit))

	request := apiRequest{
		definition:  createSearchEventsDefinition(),
		queryParams: queryParams,
		method:      http.MethodGet,
	}

	var response SearchEventsResponse
	httpResponse, err := f.doRequest(ctx, request, &response)

	return response, httpResponse, err
}

func (f *FingerprintApiService) GetRelatedVisitors(ctx context.Context, visitorId string) (RelatedVisitorsResponse, *http.Response, Error) {
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

func (f *FingerprintApiService) doRequest(ctx context.Context, apiRequest apiRequest, result any) (*http.Response, Error) {
	path := f.getPath(apiRequest.definition, apiRequest.pathParams...)
	requestUrl, err := url.Parse(path)
	if err != nil {
		return nil, WrapWithApiError(err)
	}

	query := requestUrl.Query()

	if apiRequest.queryParams != nil {
		addMapToUrlValues(apiRequest.queryParams, &query)
	}

	addIntegrationInfoToQuery(&query)

	requestUrl.RawQuery = query.Encode()

	request, err := f.prepareRequest(ctx, requestUrl, apiRequest.method, apiRequest.body)
	if err != nil {
		return nil, WrapWithApiError(err)
	}

	httpResponse, err := f.cfg.HTTPClient.Do(request)
	if err != nil {
		return httpResponse, WrapWithApiError(err)
	}

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return httpResponse, WrapWithApiError(err)
	}

	if isResponseOk(httpResponse) {
		if result != nil {
			jsonErr := json.Unmarshal(body, &result)

			if jsonErr != nil {
				return httpResponse, WrapWithApiError(jsonErr)
			}
		}

		return httpResponse, nil
	}

	apiError := handleErrorResponse(body, httpResponse, apiRequest.definition)

	return httpResponse, handlePotentialTooManyRequestsResponse(httpResponse, apiError)
}
