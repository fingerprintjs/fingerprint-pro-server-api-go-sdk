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

// Contains all the information from each activated product - BOTD and Identification
type ProductsResponse struct {
	Identification *ProductsResponseIdentification `json:"identification,omitempty"`
	Botd           *ProductsResponseBotd           `json:"botd,omitempty"`
}
