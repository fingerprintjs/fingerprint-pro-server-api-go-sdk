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

type IpBlocklist struct {
	// `true` if request IP address is part of any database that we use to search for known malicious actors, `false` otherwise.
	Result  bool                `json:"result"`
	Details *IpBlocklistDetails `json:"details"`
}
