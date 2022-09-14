# BotdResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP address of the requesting browser or bot. | [default to null]
**Time** | [**time.Time**](time.Time.md) | Time in UTC when the request from the JS agent was made. We recommend to treat requests that are older than 2 minutes as malicious. Otherwise, request replay attacks are possible | [default to null]
**Bot** | [***BotdDetectionResult**](BotdDetectionResult.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
