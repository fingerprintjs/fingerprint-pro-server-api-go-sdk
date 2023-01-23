/*
 * Fingerprint Pro Server API
 *
 * Fingerprint Pro Server API allows you to get information about visitors and about individual events in a server environment. This API can be used for data exports, decision-making, and data analysis scenarios.
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
This endpoint allows you to get events with all the information from each activated product (Fingerprint Pro or Bot Detection). Use the requestId as a URL path :request_id parameter. This API method is scoped to a request, i.e. all returned information is by requestId.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param requestId requestId is the unique identifier of each request

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

	localVarQueryParams.Add("ii", "fingerprint-pro-server-go-sdk/2.0.0-test.1")
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
This endpoint allows you to get a history of visits with all available information. Use the visitorId as a URL path parameter. This API method is scoped to a visitor, i.e. all returned information is by visitorId.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param visitorId
 * @param optional nil or *FingerprintApiGetVisitsOpts - Optional Parameters:
     * @param "RequestId" (optional.String) -  Filter visits by requestId
     * @param "LinkedId" (optional.String) -  Filter visits by custom identifier
     * @param "Limit" (optional.Int32) -  Limit scanned results
     * @param "Before" (optional.Int32) -  Used to paginate results
@return Response
*/

type FingerprintApiGetVisitsOpts struct {
	RequestId optional.String
	LinkedId  optional.String
	Limit     optional.Int32
	Before    optional.Int32
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

	localVarQueryParams.Add("ii", "fingerprint-pro-server-go-sdk/2.0.0-test.1")
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
