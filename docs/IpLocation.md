# IpLocation
This field is **deprecated** and will not return a result for **accounts created after December 18th, 2023**. Please use the [`ipInfo` Smart signal](https://dev.fingerprint.com/docs/smart-signals-overview#ip-geolocation) for geolocation information.


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccuracyRadius** | **int32** | The IP address is likely to be within this radius (in km) of the specified location. | [optional] [default to null]
**Latitude** | **float64** |  | [optional] [default to null]
**Longitude** | **float64** |  | [optional] [default to null]
**PostalCode** | **string** |  | [optional] [default to null]
**Timezone** | **string** |  | [optional] [default to null]
**City** | [***IpLocationCity**](IPLocationCity.md) |  | [optional] [default to null]
**Country** | [***Location**](Location.md) |  | [optional] [default to null]
**Continent** | [***Location**](Location.md) |  | [optional] [default to null]
**Subdivisions** | [**[]Subdivision**](Subdivision.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

