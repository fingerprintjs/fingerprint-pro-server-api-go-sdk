# Proxy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **bool** | IP address was used by a public proxy provider or belonged to a known recent residential proxy  | [default to null]
**Confidence** | [***ProxyConfidence**](ProxyConfidence.md) |  | [default to null]
**Details** | [***ProxyDetails**](ProxyDetails.md) |  | [optional] [default to null]
**MlScore** | **float64** | Machine learning-based proxy score, represented as a floating-point value between 0 and 1 (inclusive), with up to three decimal places of precision. A higher score means a higher confidence in the positive `proxy` detection result. This Smart Signal is currently in beta and only available to select customers. If you are interested, please [contact our support team](https://fingerprint.com/support/).  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

