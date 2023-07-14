/*
 * Fingerprint Pro Server API
 *
 * Fingerprint Pro Server API allows you to get information about visitors and about individual events in a server environment. It can be used for data exports, decision-making, and data analysis scenarios. Server API is intended for server-side usage, it's not intended to be used from the client side, whether it's a browser or a mobile device.
 *
 * API version: 3
 * Contact: support@fingerprint.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk

// Contains all the information from each activated product - Fingerprint Pro or Bot Detection
type ProductsResponse struct {
	Identification  *ProductsResponseIdentification `json:"identification,omitempty"`
	Botd            *ProductsResponseBotd           `json:"botd,omitempty"`
	IpInfo          *SignalResponseIpInfo           `json:"ipInfo,omitempty"`
	Incognito       *SignalResponseIncognito        `json:"incognito,omitempty"`
	RootApps        *SignalResponseRootApps         `json:"rootApps,omitempty"`
	Emulator        *SignalResponseEmulator         `json:"emulator,omitempty"`
	ClonedApp       *SignalResponseClonedApp        `json:"clonedApp,omitempty"`
	FactoryReset    *SignalResponseFactoryReset     `json:"factoryReset,omitempty"`
	Jailbroken      *SignalResponseJailbroken       `json:"jailbroken,omitempty"`
	Frida           *SignalResponseFrida            `json:"frida,omitempty"`
	IpBlocklist     *SignalResponseIpBlocklist      `json:"ipBlocklist,omitempty"`
	Tor             *SignalResponseTor              `json:"tor,omitempty"`
	PrivacySettings *SignalResponsePrivacySettings  `json:"privacySettings,omitempty"`
	VirtualMachine  *SignalResponseVirtualMachine   `json:"virtualMachine,omitempty"`
	Vpn             *SignalResponseVpn              `json:"vpn,omitempty"`
	Proxy           *SignalResponseProxy            `json:"proxy,omitempty"`
	Tampering       *SignalResponseTampering        `json:"tampering,omitempty"`
}
