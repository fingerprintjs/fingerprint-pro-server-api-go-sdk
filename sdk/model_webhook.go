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

import (
	"time"
)

type Webhook struct {
	// Unique identifier of the user's request.
	RequestId string `json:"requestId"`
	// Page URL from which the request was sent.
	Url string `json:"url"`
	// IP address of the requesting browser or bot.
	Ip string `json:"ip"`
	// Environment ID of the event.
	EnvironmentId string    `json:"environmentId,omitempty"`
	Tag           *ModelMap `json:"tag,omitempty"`
	// Time expressed according to ISO 8601 in UTC format, when the request from the JS agent was made. We recommend to treat requests that are older than 2 minutes as malicious. Otherwise, request replay attacks are possible.
	Time *time.Time `json:"time"`
	// Timestamp of the event with millisecond precision in Unix time.
	Timestamp  int64                  `json:"timestamp"`
	IpLocation *DeprecatedGeolocation `json:"ipLocation,omitempty"`
	// A customer-provided id that was sent with the request.
	LinkedId string `json:"linkedId,omitempty"`
	// String of 20 characters that uniquely identifies the visitor's browser.
	VisitorId string `json:"visitorId,omitempty"`
	// Attribute represents if a visitor had been identified before.
	VisitorFound   bool                      `json:"visitorFound,omitempty"`
	Confidence     *IdentificationConfidence `json:"confidence,omitempty"`
	FirstSeenAt    *IdentificationSeenAt     `json:"firstSeenAt,omitempty"`
	LastSeenAt     *IdentificationSeenAt     `json:"lastSeenAt,omitempty"`
	BrowserDetails *BrowserDetails           `json:"browserDetails,omitempty"`
	// Flag if user used incognito session.
	Incognito           bool                           `json:"incognito,omitempty"`
	ClientReferrer      string                         `json:"clientReferrer,omitempty"`
	Components          *map[string]RawDeviceAttribute `json:"components,omitempty"`
	Bot                 *BotdBot                       `json:"bot,omitempty"`
	UserAgent           string                         `json:"userAgent,omitempty"`
	RootApps            *WebhookRootApps               `json:"rootApps,omitempty"`
	Emulator            *WebhookEmulator               `json:"emulator,omitempty"`
	IpInfo              *WebhookIpInfo                 `json:"ipInfo,omitempty"`
	IpBlocklist         *WebhookIpBlocklist            `json:"ipBlocklist,omitempty"`
	Tor                 *WebhookTor                    `json:"tor,omitempty"`
	Vpn                 *WebhookVpn                    `json:"vpn,omitempty"`
	Proxy               *WebhookProxy                  `json:"proxy,omitempty"`
	Tampering           *WebhookTampering              `json:"tampering,omitempty"`
	ClonedApp           *WebhookClonedApp              `json:"clonedApp,omitempty"`
	FactoryReset        *WebhookFactoryReset           `json:"factoryReset,omitempty"`
	Jailbroken          *WebhookJailbroken             `json:"jailbroken,omitempty"`
	Frida               *WebhookFrida                  `json:"frida,omitempty"`
	PrivacySettings     *WebhookPrivacySettings        `json:"privacySettings,omitempty"`
	VirtualMachine      *WebhookVirtualMachine         `json:"virtualMachine,omitempty"`
	RawDeviceAttributes *map[string]RawDeviceAttribute `json:"rawDeviceAttributes,omitempty"`
	HighActivity        *WebhookHighActivity           `json:"highActivity,omitempty"`
	LocationSpoofing    *WebhookLocationSpoofing       `json:"locationSpoofing,omitempty"`
	SuspectScore        *WebhookSuspectScore           `json:"suspectScore,omitempty"`
	RemoteControl       *WebhookRemoteControl          `json:"remoteControl,omitempty"`
	Velocity            *WebhookVelocity               `json:"velocity,omitempty"`
	DeveloperTools      *WebhookDeveloperTools         `json:"developerTools,omitempty"`
	MitmAttack          *WebhookMitMAttack             `json:"mitmAttack,omitempty"`
}
