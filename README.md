<p align="center">
  <a href="https://fingerprint.com">
    <picture>
     <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/main/res/logo_light.svg" />
     <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/main/res/logo_dark.svg" />
     <img src="https://raw.githubusercontent.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/main/res/logo_dark.svg" alt="Fingerprint logo" width="312px" />
   </picture>
  </a>
</p>
<p align="center">
  <a href="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/release.yml">
    <img src="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/release.yml/badge.svg" alt="CI badge" />
  </a>
  <a href="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/tests.yml">
    <img src="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/tests.yml/badge.svg" alt="CI badge" />
  </a>
  <a href="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/functional_tests.yml">
    <img src="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/functional_tests.yml/badge.svg" alt="CI badge" />
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/:license-mit-blue.svg?style=flat"/>
  </a>
  <a href="https://discord.gg/39EpE2neBg">
    <img src="https://img.shields.io/discord/852099967190433792?style=logo&label=Discord&logo=Discord&logoColor=white" alt="Discord server">
  </a>
</p>

# Fingerprint Pro Server Go SDK
Fingerprint Pro Server API provides a way for validating visitors’ data issued by Fingerprint Pro.

This Go package is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: 3
- Package version: 1.0.1
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Requirements.

Go Lang 1.17 or higher

## Installation & Usage

1. Get the package from GitHub:
```shell
go get github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk
```

2. Import and use the library:

```go
package main

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk"
	"log"
)

func main() {
	cfg := sdk.NewConfiguration()
	client := sdk.NewAPIClient(cfg)
	// Configure authorization, in our case with API Key
	auth := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
		Key: "SECRET_API_KEY",
	})
	// Usually this data will come from your frontend app
	visitorId := "VISITOR_ID"
	opts := sdk.FingerprintApiGetVisitsOpts{
		RequestId: optional.NewString("REQUEST_ID"),
	}
	response, httpRes, err := client.FingerprintApi.GetVisits(auth, visitorId, &opts)
	fmt.Printf("%+v\n", httpRes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got response with visitorId: %s", response.VisitorId)
}
```

> **Note**
> You can also check examples located in [example](./example) directory.
> To run the examples:
> ```shell
> cd example && FINGERPRINT_API_KEY=SECRET_API_KEY VISITOR_ID=VISITOR_ID_EXAMPLE go run GetVisits_APIKey.go
> ```

## Documentation for API Endpoints

All URIs are relative to *https://api.fpjs.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FingerprintApi* | [**GetVisits**](docs/FingerprintApi.md#getvisits) | **Get** /visitors/{visitor_id} | 

## Documentation For Models

 - [BrowserDetails](docs/BrowserDetails.md)
 - [Confidence](docs/Confidence.md)
 - [IpLocation](docs/IpLocation.md)
 - [IpLocationCity](docs/IpLocationCity.md)
 - [Location](docs/Location.md)
 - [ManyRequestsResponse](docs/ManyRequestsResponse.md)
 - [Response](docs/Response.md)
 - [StSeenAt](docs/StSeenAt.md)
 - [Subdivision](docs/Subdivision.md)
 - [Visit](docs/Visit.md)
 - [WebhookVisit](docs/WebhookVisit.md)

## Documentation For Authorization


## ApiKeyHeader

- **Type**: API key
- **API key parameter name**: Auth-API-Key
- **Location**: HTTP header

## ApiKeyQuery

- **Type**: API key
- **API key parameter name**: api_key
- **Location**: URL query string


## Author

support@fingerprint.com
