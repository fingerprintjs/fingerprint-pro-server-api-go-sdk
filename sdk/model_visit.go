/*
 * Fingerprint Pro Server API
 *
 * Fingerprint Pro Server API provides a way for validating visitors’ data issued by Fingerprint Pro.
 *
 * API version: 3
 * Contact: support@fingerprint.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk

import (
	"time"
)

type Visit struct {
	// Unique identifier of the user's identification request.
	RequestId      string          `json:"requestId"`
	BrowserDetails *BrowserDetails `json:"browserDetails"`
	// Flag if user used incognito session.
	Incognito  bool        `json:"incognito"`
	Ip         string      `json:"ip"`
	IpLocation *IpLocation `json:"ipLocation"`
	// Timestamp of the event with millisecond precision in Unix time.
	Timestamp int64 `json:"timestamp"`
	// Time expressed according to ISO 8601 in UTC format.
	Time time.Time `json:"time"`
	// Page URL from which identification request was sent.
	Url string `json:"url"`
	// A customer-provided value or an object that was sent with identification request.
	Tag ModelMap `json:"tag"`
	// A customer-provided id that was sent with identification request.
	LinkedId   string      `json:"linkedId,omitempty"`
	Confidence *Confidence `json:"confidence"`
	// Attribute represents if a visitor had been identified before.
	VisitorFound bool      `json:"visitorFound"`
	FirstSeenAt  *StSeenAt `json:"firstSeenAt"`
	LastSeenAt   *StSeenAt `json:"lastSeenAt"`
}