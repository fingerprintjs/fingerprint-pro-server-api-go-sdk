# ErrorUpdateEvent400ResponseError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Error code: * `RequestCannotBeParsed` - the JSON content of the request contains some errors that prevented us from parsing it (wrong type/surpassed limits) * `Failed` - the event is more than 10 days old and cannot be updated  | [default to null]
**Message** | **string** | Details about the underlying issue with the input payload | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

