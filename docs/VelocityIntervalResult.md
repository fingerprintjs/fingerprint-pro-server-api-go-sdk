# VelocityIntervalResult
Is absent if the velocity data could not be generated for the visitor ID. 


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5m** | **int32** |  | [default to null]
**Var1h** | **int32** |  | [default to null]
**Var24h** | **int32** | The `24h` interval of `distinctIp`, `distinctLinkedId`, and `distinctCountry` will be omitted if the number of `events`` for the visitor ID in the last 24 hours (`events.intervals.['24h']`) is higher than 20.000.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

