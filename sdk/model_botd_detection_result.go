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

// Stores bot detection result
type BotdDetectionResult struct {
	// Bot detection result:  * `notDetected` - the visitor is not a bot  * `good` - good bot detected, such as Google bot, Baidu Spider, AlexaBot and so on  * `bad` - bad bot detected, such as Selenium, Puppeteer, Playwright, headless browsers, and so on
	Result string `json:"result"`
}
