# Identification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorId** | **string** | String of 20 characters that uniquely identifies the visitor's browser. | [default to null]
**RequestId** | **string** | Unique identifier of the user's request. | [default to null]
**BrowserDetails** | [***BrowserDetails**](BrowserDetails.md) |  | [default to null]
**Incognito** | **bool** | Flag if user used incognito session. | [default to null]
**Ip** | **string** | IP address of the requesting browser or bot. | [default to null]
**IpLocation** | [***DeprecatedGeolocation**](DeprecatedGeolocation.md) |  | [optional] [default to null]
**LinkedId** | **string** | A customer-provided id that was sent with the request. | [optional] [default to null]
**Timestamp** | **int64** | Timestamp of the event with millisecond precision in Unix time. | [default to null]
**Time** | [**time.Time**](time.Time.md) | Time expressed according to ISO 8601 in UTC format, when the request from the JS agent was made. We recommend to treat requests that are older than 2 minutes as malicious. Otherwise, request replay attacks are possible. | [default to null]
**Url** | **string** | Page URL from which the request was sent. | [default to null]
**Tag** | [***ModelMap**](map.md) |  | [default to null]
**Confidence** | [***IdentificationConfidence**](IdentificationConfidence.md) |  | [optional] [default to null]
**VisitorFound** | **bool** | Attribute represents if a visitor had been identified before. | [default to null]
**FirstSeenAt** | [***IdentificationSeenAt**](IdentificationSeenAt.md) |  | [default to null]
**LastSeenAt** | [***IdentificationSeenAt**](IdentificationSeenAt.md) |  | [default to null]
**Components** | [***map[string]RawDeviceAttribute**](map.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

