# ProxyDetails
Proxy detection details (present if proxy is detected)


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProxyType** | **string** | Residential proxies use real user IP addresses to appear as legitimate traffic,  while data center proxies are public proxies hosted in data centers  | [default to null]
**LastSeenAt** | [**time.Time**](time.Time.md) | ISO 8601 formatted timestamp in UTC with hourly resolution of when this IP was last seen as a proxy when available.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

