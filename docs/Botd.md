# Botd
Contains all the information from Bot Detection product


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bot** | [***BotdBot**](BotdBot.md) |  | [default to null]
**Meta** | [***ModelMap**](map.md) |  | [optional] [default to null]
**LinkedId** | **string** | A customer-provided id that was sent with the request. | [optional] [default to null]
**Url** | **string** | Page URL from which the request was sent. | [default to null]
**Ip** | **string** | IP address of the requesting browser or bot. | [default to null]
**Time** | [**time.Time**](time.Time.md) | Time in UTC when the request from the JS agent was made. We recommend to treat requests that are older than 2 minutes as malicious. Otherwise, request replay attacks are possible. | [default to null]
**UserAgent** | **string** |  | [default to null]
**RequestId** | **string** | Unique identifier of the user's request. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

