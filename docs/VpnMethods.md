# VpnMethods

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimezoneMismatch** | **bool** | The browser timezone doesn't match the timezone inferred from the request IP address. | [default to null]
**PublicVPN** | **bool** | Request IP address is owned and used by a public VPN service provider. | [default to null]
**AuxiliaryMobile** | **bool** | This method applies to mobile devices only. Indicates the result of additional methods used to detect a VPN in mobile devices. | [default to null]
**OsMismatch** | **bool** | The browser runs on a different operating system than the operating system inferred from the request network signature. | [default to null]
**Relay** | **bool** | Request IP address belongs to a relay service provider, indicating the use of relay services like [Apple Private relay](https://support.apple.com/en-us/102602) or [Cloudflare Warp](https://developers.cloudflare.com/warp-client/).   * Like VPNs, relay services anonymize the visitor's true IP address. * Unlike traditional VPNs, relay services don't let visitors spoof their location by choosing an exit node in a different country.  This field allows you to differentiate VPN users and relay service users in your fraud prevention logic.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

