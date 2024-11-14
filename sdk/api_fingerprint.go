/*
 * Fingerprint Pro Server API
 *
 * Fingerprint Pro Server API allows you to get information about visitors and about individual events in a server environment. It can be used for data exports, decision-making, and data analysis scenarios. Server API is intended for server-side usage, it's not intended to be used from the client side, whether it's a browser or a mobile device.
 *
 * API version: 3
 * Contact: support@fingerprint.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk

import (
	"context"
	"net/http"
	"strings"
)

const IntegrationInfo = "fingerprint-pro-server-go-sdk/7.0.0-test.1"

type FingerprintApiServiceInterface interface {
	/*
	   FingerprintApiService Delete data by visitor ID
	   Request deleting all data associated with the specified visitor ID. This API is useful for compliance with privacy regulations. ### Which data is deleted? - Browser (or device) properties - Identification requests made from this browser (or device)  #### Browser (or device) properties - Represents the data that Fingerprint collected from this specific browser (or device) and everything inferred and derived from it. - Upon request to delete, this data is deleted asynchronously (typically within a few minutes) and it will no longer be used to identify this browser (or device) for your [Fingerprint Application](https://dev.fingerprint.com/docs/glossary#fingerprint-application).  #### Identification requests made from this browser (or device) - Fingerprint stores the identification requests made from a browser (or device) for up to 30 (or 90) days depending on your plan. To learn more, see [Data Retention](https://dev.fingerprint.com/docs/regions#data-retention). - Upon request to delete, the identification requests that were made by this browser   - Within the past 10 days are deleted within 24 hrs.   - Outside of 10 days are allowed to purge as per your data retention period.  ### Corollary After requesting to delete a visitor ID, - If the same browser (or device) requests to identify, it will receive a different visitor ID. - If you request [`/events` API](https://dev.fingerprint.com/reference/getevent) with a `request_id` that was made outside of the 10 days, you will still receive a valid response. - If you request [`/visitors` API](https://dev.fingerprint.com/reference/getvisits) for the deleted visitor ID, the response will include identification requests that were made outside of those 10 days.  ### Interested? Please [contact our support team](https://fingerprint.com/support/) to enable it for you. Otherwise, you will receive a 403.
	    * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param visitorId The [visitor ID](https://dev.fingerprint.com/reference/get-function#visitorid) you want to delete.

	*/
	DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error)

	/*
	   FingerprintApiService Get event by request ID
	   Get a detailed analysis of an individual identification event, including Smart Signals.  Please note that the response includes mobile signals (e.g. `rootApps`) even if the request originated from a non-mobile platform. It is highly recommended that you **ignore** the mobile signals for such requests.   Use `requestId` as the URL path parameter. This API method is scoped to a request, i.e. all returned information is by `requestId`.
	    * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param requestId The unique [identifier](https://dev.fingerprint.com/reference/get-function#requestid) of each identification request.
	       @return EventsGetResponse
	*/
	GetEvent(ctx context.Context, requestId string) (EventsGetResponse, *http.Response, error)

	/*
	   FingerprintApiService Get Related Visitors
	   Related visitors API lets you link web visits and in-app browser visits that originated from the same mobile device. It searches the past 6 months of identification events to find the visitor IDs that belong to the same mobile device as the given visitor ID.  ⚠️ Please note that this API is not enabled by default and is billable separately. ⚠️  If you would like to use Related visitors API, please contact our [support team](https://fingerprint.com/support). To learn more, see [Related visitors API reference](https://dev.fingerprint.com/reference/related-visitors-api).
	    * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param visitorId The [visitor ID](https://dev.fingerprint.com/reference/get-function#visitorid) for which you want to find the other visitor IDs that originated from the same mobile device.
	       @return RelatedVisitorsResponse
	*/
	GetRelatedVisitors(ctx context.Context, visitorId string) (RelatedVisitorsResponse, *http.Response, error)

	/*
	   FingerprintApiService Get visits by visitor ID
	   Get a history of visits (identification events) for a specific `visitorId`. Use the `visitorId` as a URL path parameter. Only information from the _Identification_ product is returned.  #### Headers  * `Retry-After` — Present in case of `429 Too many requests`. Indicates how long you should wait before making a follow-up request. The value is non-negative decimal integer indicating the seconds to delay after the response is received.
	    * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param visitorId Unique [visitor identifier](https://dev.fingerprint.com/reference/get-function#visitorid) issued by Fingerprint Pro.
	    * @param optional nil or *FingerprintApiGetVisitsOpts - Optional Parameters:
	        * @param "RequestId" (string) -  Filter visits by `requestId`.   Every identification request has a unique identifier associated with it called `requestId`. This identifier is returned to the client in the identification [result](https://dev.fingerprint.com/reference/get-function#requestid). When you filter visits by `requestId`, only one visit will be returned.
	    * @param "LinkedId" (string) -  Filter visits by your custom identifier.   You can use [`linkedId`](https://dev.fingerprint.com/reference/get-function#linkedid) to associate identification requests with your own identifier, for example: session ID, purchase ID, or transaction ID. You can then use this `linked_id` parameter to retrieve all events associated with your custom identifier.
	    * @param "Limit" (int32) -  Limit scanned results.   For performance reasons, the API first scans some number of events before filtering them. Use `limit` to specify how many events are scanned before they are filtered by `requestId` or `linkedId`. Results are always returned sorted by the timestamp (most recent first). By default, the most recent 100 visits are scanned, the maximum is 500.
	    * @param "PaginationKey" (string) -  Use `paginationKey` to get the next page of results.   When more results are available (e.g., you requested 200 results using `limit` parameter, but a total of 600 results are available), the `paginationKey` top-level attribute is added to the response. The key corresponds to the `requestId` of the last returned event. In the following request, use that value in the `paginationKey` parameter to get the next page of results:  1. First request, returning most recent 200 events: `GET api-base-url/visitors/:visitorId?limit=200` 2. Use `response.paginationKey` to get the next page of results: `GET api-base-url/visitors/:visitorId?limit=200&paginationKey=1683900801733.Ogvu1j`  Pagination happens during scanning and before filtering, so you can get less visits than the `limit` you specified with more available on the next page. When there are no more results available for scanning, the `paginationKey` attribute is not returned.
	    * @param "Before" (int64) -  ⚠️ Deprecated pagination method, please use `paginationKey` instead. Timestamp (in milliseconds since epoch) used to paginate results.
	   @return VisitorsGetResponse
	*/
	GetVisits(ctx context.Context, visitorId string, opts *FingerprintApiGetVisitsOpts) (VisitorsGetResponse, *http.Response, error)

	/*
	   FingerprintApiService Update an event with a given request ID
	   Change information in existing events specified by `requestId` or *flag suspicious events*.  When an event is created, it is assigned `linkedId` and `tag` submitted through the JS agent parameters. This information might not be available on the client so the Server API allows for updating the attributes after the fact.  **Warning** It's not possible to update events older than 10 days.
	    * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param body
	    * @param requestId The unique event [identifier](https://dev.fingerprint.com/reference/get-function#requestid).

	*/
	UpdateEvent(ctx context.Context, body EventsUpdateRequest, requestId string) (*http.Response, error)
}

type requestDefinition struct {
	StatusCodeResultsFactoryMap map[int]func() any
	GetPath                     func(params ...string) string
}

func createDeleteVisitorDataDefinition() requestDefinition {
	return requestDefinition{
		GetPath: func(args ...string) string {
			pathParams := []string{"visitor_id"}

			path := "/visitors/{visitor_id}"

			for i, arg := range args {
				path = strings.Replace(path, "{"+pathParams[i]+"}", arg, -1)
			}

			return path
		},
		StatusCodeResultsFactoryMap: map[int]func() any{
			400: func() any { return &ErrorResponse{} },
			403: func() any { return &ErrorResponse{} },
			404: func() any { return &ErrorResponse{} },
			429: func() any { return &ErrorResponse{} },
		},
	}
}

func createGetEventDefinition() requestDefinition {
	return requestDefinition{
		GetPath: func(args ...string) string {
			pathParams := []string{"request_id"}

			path := "/events/{request_id}"

			for i, arg := range args {
				path = strings.Replace(path, "{"+pathParams[i]+"}", arg, -1)
			}

			return path
		},
		StatusCodeResultsFactoryMap: map[int]func() any{
			200: func() any { return &EventsGetResponse{} },
			403: func() any { return &ErrorResponse{} },
			404: func() any { return &ErrorResponse{} },
		},
	}
}

func createGetRelatedVisitorsDefinition() requestDefinition {
	return requestDefinition{
		GetPath: func(args ...string) string {
			pathParams := []string{}

			path := "/related-visitors"

			for i, arg := range args {
				path = strings.Replace(path, "{"+pathParams[i]+"}", arg, -1)
			}

			return path
		},
		StatusCodeResultsFactoryMap: map[int]func() any{
			200: func() any { return &RelatedVisitorsResponse{} },
			400: func() any { return &ErrorResponse{} },
			403: func() any { return &ErrorResponse{} },
			404: func() any { return &ErrorResponse{} },
			429: func() any { return &ErrorResponse{} },
		},
	}
}

func createGetVisitsDefinition() requestDefinition {
	return requestDefinition{
		GetPath: func(args ...string) string {
			pathParams := []string{"visitor_id"}

			path := "/visitors/{visitor_id}"

			for i, arg := range args {
				path = strings.Replace(path, "{"+pathParams[i]+"}", arg, -1)
			}

			return path
		},
		StatusCodeResultsFactoryMap: map[int]func() any{
			200: func() any { return &VisitorsGetResponse{} },
			400: func() any { return &ErrorPlainResponse{} },
			403: func() any { return &ErrorPlainResponse{} },
			429: func() any { return &ErrorPlainResponse{} },
		},
	}
}

type FingerprintApiGetVisitsOpts struct {
	RequestId     string
	LinkedId      string
	Limit         int32
	PaginationKey string
	Before        int64
}

func (o *FingerprintApiGetVisitsOpts) ToUrlValuesMap() map[string]any {
	data := make(map[string]any)

	if o == nil {
		return data
	}

	data["request_id"] = o.RequestId
	data["linked_id"] = o.LinkedId
	data["limit"] = o.Limit
	data["paginationKey"] = o.PaginationKey
	data["before"] = o.Before

	return data
}

func createUpdateEventDefinition() requestDefinition {
	return requestDefinition{
		GetPath: func(args ...string) string {
			pathParams := []string{"request_id"}

			path := "/events/{request_id}"

			for i, arg := range args {
				path = strings.Replace(path, "{"+pathParams[i]+"}", arg, -1)
			}

			return path
		},
		StatusCodeResultsFactoryMap: map[int]func() any{
			400: func() any { return &ErrorResponse{} },
			403: func() any { return &ErrorResponse{} },
			404: func() any { return &ErrorResponse{} },
			409: func() any { return &ErrorResponse{} },
		},
	}
}
