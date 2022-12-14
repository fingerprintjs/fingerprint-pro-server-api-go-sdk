# FingerprintApi

All URIs are relative to *https://api.fpjs.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](FingerprintApi.md#GetEvent) | **Get** /events/{request_id} | Get event by request ID
[**GetVisits**](FingerprintApi.md#GetVisits) | **Get** /visitors/{visitor_id} | 

# **GetEvent**
> EventResponse GetEvent(ctx, requestId)
Get event by request ID

This endpoint allows you to get events with all the information from each activated product - BOTD and Fingerprinting. Use the requestId as a URL path :request_id parameter. This API method is scoped to a request, i.e. all returned information is by requestId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Request ID | 

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


This endpoint allows you to get a history of visits with all available information. Use the visitorId as a URL path parameter. This API method is scoped to a visitor, i.e. all returned information is by visitorId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **visitorId** | **string**|  | 
 

### Optional Parameters
Optional parameters are passed through a pointer to a FingerprintApiGetVisitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **optional.String**| Filter events by requestId | 
 **linkedId** | **optional.String**| Filter events by custom identifier | 
 **limit** | **optional.Int32**| Limit scanned results | 
 **before** | **optional.Int32**| Used to paginate results | 

### Return type

[**Response**](Response.md)

### Authorization

[ApiKeyHeader](../README.md#ApiKeyHeader), [ApiKeyQuery](../README.md#ApiKeyQuery)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

