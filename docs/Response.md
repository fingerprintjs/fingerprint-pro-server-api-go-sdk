# Response

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorId** | **string** |  | [default to null]
**Visits** | [**[]ResponseVisits**](ResponseVisits.md) |  | [default to null]
**LastTimestamp** | **int64** | ⚠️ Deprecated paging attribute, please use &#x60;paginationKey&#x60; instead. Timestamp of the last visit in the current page of results.  | [optional] [default to null]
**PaginationKey** | **string** | Request ID of the last visit in the current page of results. Use this value in the following request as the &#x60;paginationKey&#x60; parameter to get the next page of results. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

