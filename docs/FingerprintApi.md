# FingerprintApi

All URIs are relative to *https://api.fpjs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVisitorData**](FingerprintApi.md#DeleteVisitorData) | **Delete** /visitors/{visitor_id} | Delete data by visitor ID
[**GetEvent**](FingerprintApi.md#GetEvent) | **Get** /events/{request_id} | Get event by request ID
[**GetRelatedVisitors**](FingerprintApi.md#GetRelatedVisitors) | **Get** /related-visitors | Get Related Visitors
[**GetVisits**](FingerprintApi.md#GetVisits) | **Get** /visitors/{visitor_id} | Get visits by visitor ID
[**SearchEvents**](FingerprintApi.md#SearchEvents) | **Get** /events/search | Get events via search
[**UpdateEvent**](FingerprintApi.md#UpdateEvent) | **Put** /events/{request_id} | Update an event with a given request ID

# **DeleteVisitorData**
> DeleteVisitorData(ctx, visitorId)
Delete data by visitor ID

Request deleting all data associated with the specified visitor ID. This API is useful for compliance with privacy regulations. ### Which data is deleted? - Browser (or device) properties - Identification requests made from this browser (or device)  #### Browser (or device) properties - Represents the data that Fingerprint collected from this specific browser (or device) and everything inferred and derived from it. - Upon request to delete, this data is deleted asynchronously (typically within a few minutes) and it will no longer be used to identify this browser (or device) for your [Fingerprint Application](https://dev.fingerprint.com/docs/glossary#fingerprint-application).  #### Identification requests made from this browser (or device) - Fingerprint stores the identification requests made from a browser (or device) for up to 30 (or 90) days depending on your plan. To learn more, see [Data Retention](https://dev.fingerprint.com/docs/regions#data-retention). - Upon request to delete, the identification requests that were made by this browser   - Within the past 10 days are deleted within 24 hrs.   - Outside of 10 days are allowed to purge as per your data retention period.  ### Corollary After requesting to delete a visitor ID, - If the same browser (or device) requests to identify, it will receive a different visitor ID. - If you request [`/events` API](https://dev.fingerprint.com/reference/getevent) with a `request_id` that was made outside of the 10 days, you will still receive a valid response. - If you request [`/visitors` API](https://dev.fingerprint.com/reference/getvisits) for the deleted visitor ID, the response will include identification requests that were made outside of those 10 days.  ### Interested? Please [contact our support team](https://fingerprint.com/support/) to enable it for you. Otherwise, you will receive a 403. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **visitorId** | **string**| The [visitor ID](https://dev.fingerprint.com/reference/get-function#visitorid) you want to delete. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvent**
> EventsGetResponse GetEvent(ctx, requestId)
Get event by request ID

Get a detailed analysis of an individual identification event, including Smart Signals.  Please note that the response includes mobile signals (e.g. `rootApps`) even if the request originated from a non-mobile platform. It is highly recommended that you **ignore** the mobile signals for such requests.   Use `requestId` as the URL path parameter. This API method is scoped to a request, i.e. all returned information is by `requestId`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| The unique [identifier](https://dev.fingerprint.com/reference/get-function#requestid) of each identification request. | 

### Return type

[**EventsGetResponse**](EventsGetResponse.md)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelatedVisitors**
> RelatedVisitorsResponse GetRelatedVisitors(ctx, visitorId)
Get Related Visitors

Related visitors API lets you link web visits and in-app browser visits that originated from the same mobile device. It searches the past 6 months of identification events to find the visitor IDs that belong to the same mobile device as the given visitor ID.  ⚠️ Please note that this API is not enabled by default and is billable separately. ⚠️  If you would like to use Related visitors API, please contact our [support team](https://fingerprint.com/support). To learn more, see [Related visitors API reference](https://dev.fingerprint.com/reference/related-visitors-api). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **visitorId** | **string**| The [visitor ID](https://dev.fingerprint.com/reference/get-function#visitorid) for which you want to find the other visitor IDs that originated from the same mobile device. | 

### Return type

[**RelatedVisitorsResponse**](RelatedVisitorsResponse.md)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVisits**
> VisitorsGetResponse GetVisits(ctx, visitorId, optional)
Get visits by visitor ID

Get a history of visits (identification events) for a specific `visitorId`. Use the `visitorId` as a URL path parameter. Only information from the _Identification_ product is returned.  #### Headers  * `Retry-After` — Present in case of `429 Too many requests`. Indicates how long you should wait before making a follow-up request. The value is non-negative decimal integer indicating the seconds to delay after the response is received. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **visitorId** | **string**| Unique [visitor identifier](https://dev.fingerprint.com/reference/get-function#visitorid) issued by Fingerprint Pro. | 
 

### Optional Parameters
Optional parameters are passed through a pointer to a FingerprintApiGetVisitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string**| Filter visits by `requestId`.   Every identification request has a unique identifier associated with it called `requestId`. This identifier is returned to the client in the identification [result](https://dev.fingerprint.com/reference/get-function#requestid). When you filter visits by `requestId`, only one visit will be returned.  | 
 **linkedId** | **string**| Filter visits by your custom identifier.   You can use [`linkedId`](https://dev.fingerprint.com/reference/get-function#linkedid) to associate identification requests with your own identifier, for example: session ID, purchase ID, or transaction ID. You can then use this `linked_id` parameter to retrieve all events associated with your custom identifier.  | 
 **limit** | **int32**| Limit scanned results.   For performance reasons, the API first scans some number of events before filtering them. Use `limit` to specify how many events are scanned before they are filtered by `requestId` or `linkedId`. Results are always returned sorted by the timestamp (most recent first). By default, the most recent 100 visits are scanned, the maximum is 500.  | 
 **paginationKey** | **string**| Use `paginationKey` to get the next page of results.   When more results are available (e.g., you requested 200 results using `limit` parameter, but a total of 600 results are available), the `paginationKey` top-level attribute is added to the response. The key corresponds to the `requestId` of the last returned event. In the following request, use that value in the `paginationKey` parameter to get the next page of results:  1. First request, returning most recent 200 events: `GET api-base-url/visitors/:visitorId?limit=200` 2. Use `response.paginationKey` to get the next page of results: `GET api-base-url/visitors/:visitorId?limit=200&paginationKey=1683900801733.Ogvu1j`  Pagination happens during scanning and before filtering, so you can get less visits than the `limit` you specified with more available on the next page. When there are no more results available for scanning, the `paginationKey` attribute is not returned.  | 
 **before** | **int64**| ⚠️ Deprecated pagination method, please use `paginationKey` instead. Timestamp (in milliseconds since epoch) used to paginate results.  | 

### Return type

[**VisitorsGetResponse**](VisitorsGetResponse.md)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchEvents**
> SearchEventsResponse SearchEvents(ctx, limit, optional)
Get events via search

Search for identification events, including Smart Signals, using multiple filtering criteria. If you don't provide `start` or `end` parameters, the default search range is the last 7 days.  Please note that events include mobile signals (e.g. `rootApps`) even if the request originated from a non-mobile platform. We recommend you **ignore** mobile signals for such requests. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **int32**| Limit the number of events returned.  | 
 **optional** | ***FingerprintApiSearchEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FingerprintApiSearchEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visitorId** | **string**| Unique [visitor identifier](https://dev.fingerprint.com/reference/get-function#visitorid) issued by Fingerprint Pro. Filter for events matching this `visitor_id`.  | 
 **bot** | **string**| Filter events by the bot detection result, specifically:    - events where any kind of bot was detected.   - events where a good bot was detected.   - events where a bad bot was detected.   - events where no bot was detected.  | 
 **ipAddress** | **string**| Filter events by IP address range. The range can be as specific as a single IP (/32 for IPv4 or /128 for IPv6)  All ip_address filters must use CIDR notation, for example, 10.0.0.0/24, 192.168.0.1/32  | 
 **linkedId** | **string**| Filter events by your custom identifier.   You can use [linked IDs](https://dev.fingerprint.com/reference/get-function#linkedid) to associate identification requests with your own identifier, for example, session ID, purchase ID, or transaction ID. You can then use this `linked_id` parameter to retrieve all events associated with your custom identifier.  | 
 **start** | **int64**| Filter events with a timestamp greater than the start time, in Unix time (milliseconds).  | 
 **end** | **int64**| Filter events with a timestamp smaller than the end time, in Unix time (milliseconds).  | 
 **reverse** | **bool**| Sort events in reverse timestamp order.  | 
 **suspect** | **bool**| Filter events previously tagged as suspicious via the [Update API](https://dev.fingerprint.com/reference/updateevent).  | 

### Return type

[**SearchEventsResponse**](SearchEventsResponse.md)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEvent**
> UpdateEvent(ctx, body, requestId)
Update an event with a given request ID

Change information in existing events specified by `requestId` or *flag suspicious events*.  When an event is created, it is assigned `linkedId` and `tag` submitted through the JS agent parameters. This information might not be available on the client so the Server API allows for updating the attributes after the fact.  **Warning** It's not possible to update events older than 10 days. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventsUpdateRequest**](EventsUpdateRequest.md)|  | 
  **requestId** | **string**| The unique event [identifier](https://dev.fingerprint.com/reference/get-function#requestid). | 

### Return type

 (empty response body)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

