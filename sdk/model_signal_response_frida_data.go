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

type SignalResponseFridaData struct {
	// iOS specific [Frida](https://frida.re/docs/ios/) detection. There are 2 values: • `true` - Frida detected • `false` - No signs of Frida or the client is not iOS.
	Result bool `json:"result,omitempty"`
}
