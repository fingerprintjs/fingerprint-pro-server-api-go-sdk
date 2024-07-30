# FingerprintApi

All URIs are relative to *https://api.fpjs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVisitorData**](FingerprintApi.md#DeleteVisitorData) | **Delete** /visitors/{visitor_id} | Delete data by visitor ID
[**GetEvent**](FingerprintApi.md#GetEvent) | **Get** /events/{request_id} | Get event by request ID
[**GetVisits**](FingerprintApi.md#GetVisits) | **Get** /visitors/{visitor_id} | Get visits by visitor ID

# **DeleteVisitorData**
> DeleteVisitorData(ctx, visitorId)
Delete data by visitor ID

Request deleting all data associated with the specified visitor ID. This API is useful for compliance with privacy regulations. All delete requests are queued:   * Recent data (10 days or newer) belonging to the specified visitor will be deleted within 24 hours. * Data from older (11 days or more) identification events  will be deleted after 90 days.  If you are interested in using this API, please [contact our support team](https://fingerprint.com/support/) to enable it for you. Otherwise, you will receive a 403. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **visitorId** | **string**| The [visitor ID](https://dev.fingerprint.com/docs/js-agent#visitorid) you want to delete. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvent**
> EventResponse GetEvent(ctx, requestId)
Get event by request ID

Get a detailed analysis of an individual identification event, including Smart Signals.  Please note that the response includes mobile signals (e.g. `rootApps`) even if the request originated from a non-mobile platform. It is highly recommended that you **ignore** the mobile signals for such requests.   Use `requestId` as the URL path parameter. This API method is scoped to a request, i.e. all returned information is by `requestId`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| The unique [identifier](https://dev.fingerprint.com/docs/js-agent#requestid) of each identification request. | 

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
Get visits by visitor ID

Get a history of visits (identification events) for a specific `visitorId`. Use the `visitorId` as a URL path parameter. Only information from the _Identification_ product is returned.  #### Headers  * `Retry-After` — Present in case of `429 Too many requests`. Indicates how long you should wait before making a follow-up request. The value is non-negative decimal integer indicating the seconds to delay after the response is received. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **visitorId** | **string**| Unique [visitor identifier](https://dev.fingerprint.com/docs/js-agent#visitorid) issued by Fingerprint Pro. | 
 

### Optional Parameters
Optional parameters are passed through a pointer to a FingerprintApiGetVisitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string**| Filter visits by `requestId`.   Every identification request has a unique identifier associated with it called `requestId`. This identifier is returned to the client in the identification [result](https://dev.fingerprint.com/docs/js-agent#requestid). When you filter visits by `requestId`, only one visit will be returned.  | 
 **linkedId** | **string**| Filter visits by your custom identifier.   You can use [`linkedId`](https://dev.fingerprint.com/docs/js-agent#linkedid) to associate identification requests with your own identifier, for example: session ID, purchase ID, or transaction ID. You can then use this `linked_id` parameter to retrieve all events associated with your custom identifier.  | 
 **limit** | **int32**| Limit scanned results.   For performance reasons, the API first scans some number of events before filtering them. Use `limit` to specify how many events are scanned before they are filtered by `requestId` or `linkedId`. Results are always returned sorted by the timestamp (most recent first). By default, the most recent 100 visits are scanned, the maximum is 500.  | 
 **paginationKey** | **string**| Use `paginationKey` to get the next page of results.   When more results are available (e.g., you requested 200 results using `limit` parameter, but a total of 600 results are available), the `paginationKey` top-level attribute is added to the response. The key corresponds to the `requestId` of the last returned event. In the following request, use that value in the `paginationKey` parameter to get the next page of results:  1. First request, returning most recent 200 events: `GET api-base-url/visitors/:visitorId?limit=200` 2. Use `response.paginationKey` to get the next page of results: `GET api-base-url/visitors/:visitorId?limit=200&paginationKey=1683900801733.Ogvu1j`  Pagination happens during scanning and before filtering, so you can get less visits than the `limit` you specified with more available on the next page. When there are no more results available for scanning, the `paginationKey` attribute is not returned.  | 
 **before** | **int64**| ⚠️ Deprecated pagination method, please use `paginationKey` instead. Timestamp (in milliseconds since epoch) used to paginate results.  | 

### Return type

[**Response**](Response.md)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

