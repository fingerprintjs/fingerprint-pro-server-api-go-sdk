# FingerprintApi

All URIs are relative to *https://api.fpjs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](FingerprintApi.md#GetEvent) | **Get** /events/{request_id} | Get event by requestId
[**GetVisits**](FingerprintApi.md#GetVisits) | **Get** /visitors/{visitor_id} | Get visits by visitorId

# **GetEvent**
> EventResponse GetEvent(ctx, requestId)
Get event by requestId

This endpoint allows you to retrieve an individual analysis event with all the information from each activated product (Identification, Bot Detection, and others). Products that are not activated for your application or not relevant to the event's detected platform (web, iOS, Android) are not included in the response.   Use `requestId` as the URL path parameter. This API method is scoped to a request, i.e. all returned information is by `requestId`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| The unique [identifier](https://dev.fingerprint.com/docs/js-agent#requestid) of each analysis request. | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVisits**
> Response GetVisits(ctx, visitorId, optional)
Get visits by visitorId

This endpoint allows you to get a history of visits for a specific `visitorId`. Use the `visitorId` as a URL path parameter. Only information from the _Identification_ product is returned.  #### Headers  * `Retry-After` — Present in case of `429 Too many requests`. Indicates how long you should wait before making a follow-up request. The value is non-negative decimal integer indicating the seconds to delay after the response is received. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **visitorId** | **string**| Unique identifier of the visitor issued by Fingerprint Pro. | 
 

### Optional Parameters
Optional parameters are passed through a pointer to a FingerprintApiGetVisitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **optional.String**| Filter visits by &#x60;requestId&#x60;.   Every identification request has a unique identifier associated with it called &#x60;requestId&#x60;. This identifier is returned to the client in the identification [result](https://dev.fingerprint.com/docs/js-agent#requestid). When you filter visits by &#x60;requestId&#x60;, only one visit will be returned.  | 
 **linkedId** | **optional.String**| Filter visits by your custom identifier.   You can use [&#x60;linkedId&#x60;](https://dev.fingerprint.com/docs/js-agent#linkedid) to associate identification requests with your own identifier, for example: session ID, purchase ID, or transaction ID. You can then use this &#x60;linked_id&#x60; parameter to retrieve all events associated with your custom identifier.  | 
 **limit** | **optional.Int32**| Limit scanned results.   For performance reasons, the API first scans some number of events before filtering them. Use &#x60;limit&#x60; to specify how many events are scanned before they are filtered by &#x60;requestId&#x60; or &#x60;linkedId&#x60;. Results are always returned sorted by the timestamp (most recent first). By default, the most recent 100 visits are scanned, the maximum is 500.  | 
 **paginationKey** | **optional.String**| Use &#x60;paginationKey&#x60; to get the next page of results.   When more results are available (e.g., you requested 200 results using &#x60;limit&#x60; parameter, but a total of 600 results are available), the &#x60;paginationKey&#x60; top-level attribute is added to the response. The key corresponds to the &#x60;requestId&#x60; of the last returned event. In the following request, use that value in the &#x60;paginationKey&#x60; parameter to get the next page of results:  1. First request, returning most recent 200 events: &#x60;GET api-base-url/visitors/:visitorId?limit&#x3D;200&#x60; 2. Use &#x60;response.paginationKey&#x60; to get the next page of results: &#x60;GET api-base-url/visitors/:visitorId?limit&#x3D;200&amp;paginationKey&#x3D;1683900801733.Ogvu1j&#x60;  Pagination happens during scanning and before filtering, so you can get less visits than the &#x60;limit&#x60; you specified with more available on the next page. When there are no more results available for scanning, the &#x60;paginationKey&#x60; attribute is not returned.  | 
 **before** | **optional.Int64**| ⚠️ Deprecated pagination method, please use &#x60;paginationKey&#x60; instead. Timestamp (in milliseconds since epoch) used to paginate results.  | 

### Return type

[**Response**](Response.md)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

