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
Fingerprint Pro Server API allows you to get information about visitors and about individual events in a server environment. It can be used for data exports, decision-making, and data analysis scenarios. Server API is intended for server-side usage, it's not intended to be used from the client side, whether it's a browser or a mobile device. 

This Go package is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: 3
- Package version: 3.3.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Requirements.

Go Lang 1.17 or higher

## Installation & Usage

1. Get the package from GitHub:
```shell
go get github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v3
```

2. Import and use the library:

```go
package main

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk/v3"
	"log"
)

func main() {
	cfg := sdk.NewConfiguration()
	client := sdk.NewAPIClient(cfg)

	// You can also use sdk.RegionUS or sdk.RegionAsia. Default one is sdk.RegionUS
	//cfg.ChangeRegion(sdk.RegionEU)

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
		switch err.(type) {
		case *sdk.GenericSwaggerError:
			switch model := err.(sdk.GenericSwaggerError).Model().(type) {
			case sdk.ManyRequestsResponse:
				log.Printf("Too many requests, retry after %d seconds", model.RetryAfter)
			}

		default:
			log.Fatal(err)
		}

	}
	fmt.Printf("Got response with visitorId: %s", response.VisitorId)
}
```

> **Note**
> You can also check examples located in [example](./example) directory.
> To run the examples:
> ```shell
> cd example && FINGERPRINT_API_KEY=SECRET_API_KEY VISITOR_ID=VISITOR_ID_EXAMPLE go run getVisits.go
> ```
> Alternatively, you can define your environment variables inside `example/.env` file and run the examples without passing them as arguments. 
> If your subscription region is not the “Global/US” region, use `REGION=eu` or `REGION=ap` in the line above or in your local `.env` file.

### Region
If your subscription is in region other than US, you need to change the region in the configuration:
```go

import (
    "github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk/v3"
)

func main() {
    cfg := sdk.NewConfiguration()

    cfg.ChangeRegion(sdk.RegionEU) // or sdk.RegionAsia
}
```

## Documentation for API Endpoints

All URIs are relative to *https://api.fpjs.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FingerprintApi* | [**GetEvent**](docs/FingerprintApi.md#getevent) | **Get** /events/{request_id} | Get event by requestId
*FingerprintApi* | [**GetVisits**](docs/FingerprintApi.md#getvisits) | **Get** /visitors/{visitor_id} | Get visits by visitorId

## Documentation For Models

 - [BotdDetectionResult](docs/BotdDetectionResult.md)
 - [BotdResult](docs/BotdResult.md)
 - [BrowserDetails](docs/BrowserDetails.md)
 - [Confidence](docs/Confidence.md)
 - [ErrorEvent403Response](docs/ErrorEvent403Response.md)
 - [ErrorEvent403ResponseError](docs/ErrorEvent403ResponseError.md)
 - [ErrorEvent404Response](docs/ErrorEvent404Response.md)
 - [ErrorEvent404ResponseError](docs/ErrorEvent404ResponseError.md)
 - [ErrorVisits403](docs/ErrorVisits403.md)
 - [EventResponse](docs/EventResponse.md)
 - [IdentificationError](docs/IdentificationError.md)
 - [IpBlockListResult](docs/IpBlockListResult.md)
 - [IpBlockListResultDetails](docs/IpBlockListResultDetails.md)
 - [IpInfoResult](docs/IpInfoResult.md)
 - [IpInfoResultV4](docs/IpInfoResultV4.md)
 - [IpInfoResultV6](docs/IpInfoResultV6.md)
 - [IpLocation](docs/IpLocation.md)
 - [IpLocationCity](docs/IpLocationCity.md)
 - [Location](docs/Location.md)
 - [ManyRequestsResponse](docs/ManyRequestsResponse.md)
 - [ProductError](docs/ProductError.md)
 - [ProductsResponse](docs/ProductsResponse.md)
 - [ProductsResponseBotd](docs/ProductsResponseBotd.md)
 - [ProductsResponseIdentification](docs/ProductsResponseIdentification.md)
 - [ProductsResponseIdentificationData](docs/ProductsResponseIdentificationData.md)
 - [Response](docs/Response.md)
 - [ResponseVisits](docs/ResponseVisits.md)
 - [SeenAt](docs/SeenAt.md)
 - [SignalResponseClonedApp](docs/SignalResponseClonedApp.md)
 - [SignalResponseClonedAppData](docs/SignalResponseClonedAppData.md)
 - [SignalResponseEmulator](docs/SignalResponseEmulator.md)
 - [SignalResponseEmulatorData](docs/SignalResponseEmulatorData.md)
 - [SignalResponseFactoryReset](docs/SignalResponseFactoryReset.md)
 - [SignalResponseFactoryResetData](docs/SignalResponseFactoryResetData.md)
 - [SignalResponseFrida](docs/SignalResponseFrida.md)
 - [SignalResponseFridaData](docs/SignalResponseFridaData.md)
 - [SignalResponseIncognito](docs/SignalResponseIncognito.md)
 - [SignalResponseIncognitoData](docs/SignalResponseIncognitoData.md)
 - [SignalResponseIpBlocklist](docs/SignalResponseIpBlocklist.md)
 - [SignalResponseIpInfo](docs/SignalResponseIpInfo.md)
 - [SignalResponseJailbroken](docs/SignalResponseJailbroken.md)
 - [SignalResponseJailbrokenData](docs/SignalResponseJailbrokenData.md)
 - [SignalResponsePrivacySettings](docs/SignalResponsePrivacySettings.md)
 - [SignalResponsePrivacySettingsData](docs/SignalResponsePrivacySettingsData.md)
 - [SignalResponseProxy](docs/SignalResponseProxy.md)
 - [SignalResponseProxyData](docs/SignalResponseProxyData.md)
 - [SignalResponseRawDeviceAttributes](docs/SignalResponseRawDeviceAttributes.md)
 - [SignalResponseRootApps](docs/SignalResponseRootApps.md)
 - [SignalResponseRootAppsData](docs/SignalResponseRootAppsData.md)
 - [SignalResponseTampering](docs/SignalResponseTampering.md)
 - [SignalResponseTor](docs/SignalResponseTor.md)
 - [SignalResponseTorData](docs/SignalResponseTorData.md)
 - [SignalResponseVirtualMachine](docs/SignalResponseVirtualMachine.md)
 - [SignalResponseVirtualMachineData](docs/SignalResponseVirtualMachineData.md)
 - [SignalResponseVpn](docs/SignalResponseVpn.md)
 - [Subdivision](docs/Subdivision.md)
 - [TamperingResult](docs/TamperingResult.md)
 - [Visit](docs/Visit.md)
 - [VpnResult](docs/VpnResult.md)
 - [VpnResultMethods](docs/VpnResultMethods.md)
 - [WebhookSignalResponseClonedApp](docs/WebhookSignalResponseClonedApp.md)
 - [WebhookSignalResponseEmulator](docs/WebhookSignalResponseEmulator.md)
 - [WebhookSignalResponseFactoryReset](docs/WebhookSignalResponseFactoryReset.md)
 - [WebhookSignalResponseFrida](docs/WebhookSignalResponseFrida.md)
 - [WebhookSignalResponseJailbroken](docs/WebhookSignalResponseJailbroken.md)
 - [WebhookSignalResponsePrivacySettings](docs/WebhookSignalResponsePrivacySettings.md)
 - [WebhookSignalResponseProxy](docs/WebhookSignalResponseProxy.md)
 - [WebhookSignalResponseRootApps](docs/WebhookSignalResponseRootApps.md)
 - [WebhookSignalResponseTor](docs/WebhookSignalResponseTor.md)
 - [WebhookSignalResponseVirtualMachine](docs/WebhookSignalResponseVirtualMachine.md)
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
