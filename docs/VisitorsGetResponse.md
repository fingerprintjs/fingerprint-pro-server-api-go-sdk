# VisitorsGetResponse
Deprecated response shape for `GET /visitors/{visitor_id}`. The `visits` array currently contains at most one item. Use `GET /events/search` for multi-event history and filtering.


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorId** | **string** |  | [default to null]
**Visits** | [**[]Visit**](Visit.md) |  | [default to null]
**LastTimestamp** | **int64** | ⚠️ Deprecated paging attribute, please use `paginationKey` instead. Timestamp of the last visit in the current page of results.  | [optional] [default to null]
**PaginationKey** | **string** | Use this value in the following request as the `paginationKey` parameter to get the next result. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

