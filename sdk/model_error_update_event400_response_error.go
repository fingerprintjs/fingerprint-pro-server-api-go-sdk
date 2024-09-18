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

type ErrorUpdateEvent400ResponseError struct {
	// Error code: * `RequestCannotBeParsed` - the JSON content of the request contains some errors that prevented us from parsing it (wrong type/surpassed limits) * `Failed` - the event is more than 10 days old and cannot be updated
	Code string `json:"code"`
	// Details about the underlying issue with the input payload
	Message string `json:"message"`
}
