# Common403ErrorResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Error code:  * `TokenRequired` - `Auth-API-Key` header is missing or empty  * `TokenNotFound` - No Fingerprint application found for specified secret key  * `SubscriptionNotActive` - Fingerprint application is not active  * `WrongRegion` - server and application region differ  * `FeatureNotEnabled` - this feature (for example, Delete API) is not enabled for your application   | [default to null]
**Message** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

