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

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

type Region string

const (
	RegionEU   Region = "eu"
	RegionUS   Region = "us"
	RegionAsia Region = "asia"
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	region        Region
	basePath      string
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

// ChangeBasePath Change base path to allow switching to mocks
func (c *Configuration) ChangeBasePath(path string) {
	c.basePath = path
}

// ChangeRegion Changes region and sets basePath for it
func (c *Configuration) ChangeRegion(region Region) {
	c.region = region
	c.basePath = resolveBasePath(region)
}

func (c *Configuration) GetBasePath() string {
	return c.basePath
}

func (c *Configuration) GetRegion() Region {
	return c.region
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/7.7.0/go",
		region:        RegionUS,
	}
	cfg.basePath = resolveBasePath(cfg.region)

	return cfg
}

func resolveBasePath(region Region) string {
	switch region {
	case RegionEU:
		return "https://eu.api.fpjs.io"

	case RegionAsia:
		return "https://ap.api.fpjs.io"

	default:
		return "https://api.fpjs.io"
	}
}
