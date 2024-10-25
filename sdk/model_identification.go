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

type Identification struct {
	// String of 20 characters that uniquely identifies the visitor's browser.
	VisitorId string `json:"visitorId"`
	// Unique identifier of the user's request.
	RequestId      string          `json:"requestId"`
	BrowserDetails *BrowserDetails `json:"browserDetails"`
	// Flag if user used incognito session.
	Incognito bool `json:"incognito"`
	// IP address of the requesting browser or bot.
	Ip         string                 `json:"ip"`
	IpLocation *DeprecatedGeolocation `json:"ipLocation,omitempty"`
	// A customer-provided id that was sent with the request.
	LinkedId string `json:"linkedId,omitempty"`
	// Timestamp of the event with millisecond precision in Unix time.
	Timestamp int64 `json:"timestamp"`
	// Time expressed according to ISO 8601 in UTC format, when the request from the JS agent was made. We recommend to treat requests that are older than 2 minutes as malicious. Otherwise, request replay attacks are possible.
	Time *time.Time `json:"time"`
	// Page URL from which the request was sent.
	Url        string                    `json:"url"`
	Tag        *ModelMap                 `json:"tag"`
	Confidence *IdentificationConfidence `json:"confidence,omitempty"`
	// Attribute represents if a visitor had been identified before.
	VisitorFound bool                           `json:"visitorFound"`
	FirstSeenAt  *IdentificationSeenAt          `json:"firstSeenAt"`
	LastSeenAt   *IdentificationSeenAt          `json:"lastSeenAt"`
	Components   *map[string]RawDeviceAttribute `json:"components,omitempty"`
}
