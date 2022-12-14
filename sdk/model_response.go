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

// Fields `lastTimestamp` and `paginationKey` added when `limit` or `before` parameter provided and there is more data to show
type Response struct {
	VisitorId string           `json:"visitorId"`
	Visits    []ResponseVisits `json:"visits"`
	// When more results are available (e.g., you scanned 200 results using `limit` parameter, but a total of 600 results are available), a special `lastTimestamp` top-level attribute is added to the response. If you want to paginate the results further in the past, you should use the value of this attribute.
	LastTimestamp int64 `json:"lastTimestamp,omitempty"`
	// Visit's `requestId` of the last visit in the current page.
	PaginationKey string `json:"paginationKey,omitempty"`
}
