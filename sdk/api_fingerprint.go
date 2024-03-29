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
	"fmt"
	"github.com/antihax/optional"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type FingerprintApiService service

/*
FingerprintApiService Get event by requestId
This endpoint allows you to get a detailed analysis of an individual request.  **Only for Enterprise customers:** Please note that the response includes mobile signals (e.g. `rootApps`) even if the request originated from a non-mobile platform. It is highly recommended that you **ignore** the mobile signals for such requests.   Use `requestId` as the URL path parameter. This API method is scoped to a request, i.e. all returned information is by `requestId`.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param requestId The unique [identifier](https://dev.fingerprint.com/docs/js-agent#requestid) of each analysis request.

@return EventResponse
*/
func (a *FingerprintApiService) GetEvent(ctx context.Context, requestId string) (EventResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = http.MethodGet
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue EventResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/events/{request_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"request_id"+"}", fmt.Sprintf("%v", requestId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("ii", "fingerprint-pro-server-go-sdk/5.0.2")
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Auth-API-Key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))

		ParseResponse(localVarHttpResponse, &err)

		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v EventResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()

				ParseResponse(localVarHttpResponse, &newErr)

				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v

			ParseResponse(localVarHttpResponse, &newErr)

			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorEvent403Response
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()

				ParseResponse(localVarHttpResponse, &newErr)

				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v

			ParseResponse(localVarHttpResponse, &newErr)

			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v ErrorEvent404Response
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()

				ParseResponse(localVarHttpResponse, &newErr)

				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v

			ParseResponse(localVarHttpResponse, &newErr)

			return localVarReturnValue, localVarHttpResponse, newErr
		}

		ParseResponse(localVarHttpResponse, &newErr)

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
FingerprintApiService Get visits by visitorId
This endpoint allows you to get a history of visits for a specific `visitorId`. Use the `visitorId` as a URL path parameter. Only information from the _Identification_ product is returned.  #### Headers  * `Retry-After` — Present in case of `429 Too many requests`. Indicates how long you should wait before making a follow-up request. The value is non-negative decimal integer indicating the seconds to delay after the response is received.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param visitorId Unique identifier of the visitor issued by Fingerprint Pro.
 * @param optional nil or *FingerprintApiGetVisitsOpts - Optional Parameters:
     * @param "RequestId" (optional.String) -  Filter visits by `requestId`.   Every identification request has a unique identifier associated with it called `requestId`. This identifier is returned to the client in the identification [result](https://dev.fingerprint.com/docs/js-agent#requestid). When you filter visits by `requestId`, only one visit will be returned.
     * @param "LinkedId" (optional.String) -  Filter visits by your custom identifier.   You can use [`linkedId`](https://dev.fingerprint.com/docs/js-agent#linkedid) to associate identification requests with your own identifier, for example: session ID, purchase ID, or transaction ID. You can then use this `linked_id` parameter to retrieve all events associated with your custom identifier.
     * @param "Limit" (optional.Int32) -  Limit scanned results.   For performance reasons, the API first scans some number of events before filtering them. Use `limit` to specify how many events are scanned before they are filtered by `requestId` or `linkedId`. Results are always returned sorted by the timestamp (most recent first). By default, the most recent 100 visits are scanned, the maximum is 500.
     * @param "PaginationKey" (optional.String) -  Use `paginationKey` to get the next page of results.   When more results are available (e.g., you requested 200 results using `limit` parameter, but a total of 600 results are available), the `paginationKey` top-level attribute is added to the response. The key corresponds to the `requestId` of the last returned event. In the following request, use that value in the `paginationKey` parameter to get the next page of results:  1. First request, returning most recent 200 events: `GET api-base-url/visitors/:visitorId?limit=200` 2. Use `response.paginationKey` to get the next page of results: `GET api-base-url/visitors/:visitorId?limit=200&paginationKey=1683900801733.Ogvu1j`  Pagination happens during scanning and before filtering, so you can get less visits than the `limit` you specified with more available on the next page. When there are no more results available for scanning, the `paginationKey` attribute is not returned.
     * @param "Before" (optional.Int64) -  ⚠️ Deprecated pagination method, please use `paginationKey` instead. Timestamp (in milliseconds since epoch) used to paginate results.
@return Response
*/

type FingerprintApiGetVisitsOpts struct {
	RequestId     optional.String
	LinkedId      optional.String
	Limit         optional.Int32
	PaginationKey optional.String
	Before        optional.Int64
}

func (a *FingerprintApiService) GetVisits(ctx context.Context, visitorId string, localVarOptionals *FingerprintApiGetVisitsOpts) (Response, *http.Response, error) {
	var (
		localVarHttpMethod  = http.MethodGet
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/visitors/{visitor_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"visitor_id"+"}", fmt.Sprintf("%v", visitorId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("ii", "fingerprint-pro-server-go-sdk/5.0.2")
	if localVarOptionals != nil {

		if localVarOptionals.RequestId.IsSet() {
			localVarQueryParams.Add("request_id", parameterToString(localVarOptionals.RequestId.Value(), ""))
		}
		if localVarOptionals.LinkedId.IsSet() {
			localVarQueryParams.Add("linked_id", parameterToString(localVarOptionals.LinkedId.Value(), ""))
		}
		if localVarOptionals.Limit.IsSet() {
			localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
		}
		if localVarOptionals.PaginationKey.IsSet() {
			localVarQueryParams.Add("paginationKey", parameterToString(localVarOptionals.PaginationKey.Value(), ""))
		}
		if localVarOptionals.Before.IsSet() {
			localVarQueryParams.Add("before", parameterToString(localVarOptionals.Before.Value(), ""))
		}
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Auth-API-Key"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))

		ParseResponse(localVarHttpResponse, &err)

		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Response
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()

				ParseResponse(localVarHttpResponse, &newErr)

				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v

			ParseResponse(localVarHttpResponse, &newErr)

			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v ErrorVisits403
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()

				ParseResponse(localVarHttpResponse, &newErr)

				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v

			ParseResponse(localVarHttpResponse, &newErr)

			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v ManyRequestsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()

				ParseResponse(localVarHttpResponse, &newErr)

				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v

			ParseResponse(localVarHttpResponse, &newErr)

			return localVarReturnValue, localVarHttpResponse, newErr
		}

		ParseResponse(localVarHttpResponse, &newErr)

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
