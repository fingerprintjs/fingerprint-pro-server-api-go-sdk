# WebhookVisit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorId** | **string** |  | [default to null]
**ClientReferrer** | **string** |  | [optional] [default to null]
**UserAgent** | **string** |  | [optional] [default to null]
**Bot** | [***BotdDetectionResult**](BotdDetectionResult.md) |  | [optional] [default to null]
**IpInfo** | [***SignalResponseIpInfo**](SignalResponseIpInfo.md) |  | [optional] [default to null]
**Incognito** | **bool** | Flag if user used incognito session. | [default to null]
**RootApps** | [***SignalResponseRootApps**](SignalResponseRootApps.md) |  | [optional] [default to null]
**Emulator** | [***SignalResponseEmulator**](SignalResponseEmulator.md) |  | [optional] [default to null]
**IpBlocklist** | [***SignalResponseIpBlocklist**](SignalResponseIpBlocklist.md) |  | [optional] [default to null]
**Tor** | [***SignalResponseTor**](SignalResponseTor.md) |  | [optional] [default to null]
**Vpn** | [***SignalResponseVpn**](SignalResponseVpn.md) |  | [optional] [default to null]
**Proxy** | [***SignalResponseProxy**](SignalResponseProxy.md) |  | [optional] [default to null]
**Tampering** | [***SignalResponseTampering**](SignalResponseTampering.md) |  | [optional] [default to null]
**RequestId** | **string** | Unique identifier of the user&#x27;s identification request. | [default to null]
**BrowserDetails** | [***BrowserDetails**](BrowserDetails.md) |  | [default to null]
**Ip** | **string** |  | [default to null]
**IpLocation** | [***IpLocation**](IPLocation.md) |  | [default to null]
**Timestamp** | **int64** | Timestamp of the event with millisecond precision in Unix time. | [default to null]
**Time** | [**time.Time**](time.Time.md) | Time expressed according to ISO 8601 in UTC format. | [default to null]
**Url** | **string** | Page URL from which identification request was sent. | [default to null]
**Tag** | [**ModelMap**](interface{}.md) | A customer-provided value or an object that was sent with identification request. | [optional] [default to null]
**LinkedId** | **string** | A customer-provided id that was sent with identification request. | [optional] [default to null]
**Confidence** | [***Confidence**](Confidence.md) |  | [default to null]
**VisitorFound** | **bool** | Attribute represents if a visitor had been identified before. | [default to null]
**FirstSeenAt** | [***SeenAt**](SeenAt.md) |  | [default to null]
**LastSeenAt** | [***SeenAt**](SeenAt.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

