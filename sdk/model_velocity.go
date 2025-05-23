/*
 * Fingerprint Server API
 *
 * Fingerprint Server API allows you to search, update, and delete identification events in a server environment. It can be used for data exports, decision-making, and data analysis scenarios. Server API is intended for server-side usage, it's not intended to be used from the client side, whether it's a browser or a mobile device.
 *
 * API version: 3
 * Contact: support@fingerprint.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk

// Sums key data points for a specific `visitorId`, `ipAddress` and `linkedId` at three distinct time intervals: 5 minutes, 1 hour, and 24 hours as follows:   - Number of distinct IP addresses associated to the visitor ID. - Number of distinct linked IDs associated with the visitor ID. - Number of distinct countries associated with the visitor ID. - Number of identification events associated with the visitor ID. - Number of identification events associated with the detected IP address. - Number of distinct IP addresses associated with the provided linked ID. - Number of distinct visitor IDs associated with the provided linked ID.  The `24h` interval of `distinctIp`, `distinctLinkedId`, `distinctCountry`, `distinctIpByLinkedId` and `distinctVisitorIdByLinkedId` will be omitted  if the number of `events` for the visitor ID in the last 24 hours (`events.intervals.['24h']`) is higher than 20.000.
type Velocity struct {
	DistinctIp                  *VelocityData `json:"distinctIp"`
	DistinctLinkedId            *VelocityData `json:"distinctLinkedId"`
	DistinctCountry             *VelocityData `json:"distinctCountry"`
	Events                      *VelocityData `json:"events"`
	IpEvents                    *VelocityData `json:"ipEvents"`
	DistinctIpByLinkedId        *VelocityData `json:"distinctIpByLinkedId"`
	DistinctVisitorIdByLinkedId *VelocityData `json:"distinctVisitorIdByLinkedId"`
}
