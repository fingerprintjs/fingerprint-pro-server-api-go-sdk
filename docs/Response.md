# Response
Fields `lastTimestamp` and `paginationKey` added when `limit` or `before` parameter provided and there is more data to show


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorId** | **string** |  | [default to null]
**Visits** | [**[]ResponseVisits**](ResponseVisits.md) |  | [default to null]
**LastTimestamp** | **int64** | ⚠️ Deprecated paging attribute, please use `paginationKey` instead. Timestamp of the last visit in the current page of results.  | [optional] [default to null]
**PaginationKey** | **string** | Request ID of the last visit in the current page of results. Use this value in the following request as the `paginationKey` parameter to get the next page of results. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

