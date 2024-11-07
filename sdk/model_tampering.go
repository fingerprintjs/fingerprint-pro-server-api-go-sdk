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

type Tampering struct {
	// Flag indicating browser tampering was detected. This happens when either of these conditions is true:   * There are inconsistencies in the browser configuration that cross our internal tampering thresholds (indicated by `anomalyScore`).   * The browser signature resembles one of \"anti-detect\" browsers specifically designed to evade identification and fingerprinting, for example, Incognition (indicated by `antiDetectBrowser`).
	Result bool `json:"result"`
	// Confidence score (`0.0 - 1.0`) for tampering detection:   * Values above `0.5` indicate that there was a tampering attempt.    * Values below `0.5` indicate genuine browsers.
	AnomalyScore float64 `json:"anomalyScore"`
	// Is `true` if the identified browser resembles one of \"anti-detect\" browsers, for example, Incognition.  Anti-detect browsers try to evade identification by masking or manipulating their fingerprint to imitate legitimate browser configurations.
	AntiDetectBrowser bool `json:"antiDetectBrowser"`
}