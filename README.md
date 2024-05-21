<p align="center">
  <a href="https://fingerprint.com">
    <picture>
     <source media="(prefers-color-scheme: dark)" srcset="https://fingerprintjs.github.io/home/resources/logo_light.svg" />
     <source media="(prefers-color-scheme: light)" srcset="https://fingerprintjs.github.io/home/resources/logo_dark.svg" />
     <img src="https://fingerprintjs.github.io/home/resources/logo_dark.svg" alt="Fingerprint logo" width="312px" />
   </picture>
  </a>
</p>
<p align="center">
<a href="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/release.yml"><img src="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/release.yml/badge.svg" alt="CI badge" /></a>
<a href="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/tests.yml"><img src="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/tests.yml/badge.svg" alt="CI badge" /></a>
<a href="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/functional_tests.yml"><img src="https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/actions/workflows/functional_tests.yml/badge.svg" alt="CI badge" /></a>
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/:license-mit-blue.svg?style=flat"/></a>
<a href="https://discord.gg/39EpE2neBg"><img src="https://img.shields.io/discord/852099967190433792?style=logo&label=Discord&logo=Discord&logoColor=white" alt="Discord server"></a>
</p>

# Fingerprint Pro Server Go SDK
[Fingerprint](https://fingerprint.com/) is a device intelligence platform offering 99.5% accurate visitor identification.
Fingerprint Pro Server API allows you to get information about visitors and about individual events in a server environment. It can be used for data exports, decision-making, and data analysis scenarios. Server API is intended for server-side usage, it's not intended to be used from the client side, whether it's a browser or a mobile device. 

This Go package is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: 3
- Package version: 5.0.2
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Requirements

Go Lang 1.20 or higher

We keep the [Go support policy](https://go.dev/doc/devel/release) and support the last two major versions of Go.

## Installation & Usage

1. Get the package from GitHub:
```shell
go get github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk
```

2. Import and use the library:

```go
package main

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk"
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
    "github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk"
)

func main() {
    cfg := sdk.NewConfiguration()

    cfg.ChangeRegion(sdk.RegionEU) // or sdk.RegionAsia
}
```

## Sealed results

This SDK provides utility methods for decoding [sealed results](https://dev.fingerprint.com/docs/sealed-client-results).
Install the sealed results dependency as below:
```shell
go get github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk/sealed
```
Then you can use below code to unseal results:
```go
package main

import (
	"encoding/base64"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk/sealed"
	"os"
)

// Utility function to decode base64 string
func base64Decode(input string) []byte {
	output, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	// Sealed result from the frontend.
	sealedResult := base64Decode(os.Getenv("BASE64_SEALED_RESULT"))
	// Base64 encoded key generated in the dashboard.
	key := base64Decode(os.Getenv("BASE64_SEALED_RESULT_KEY"))

	keys := []sealed.DecryptionKey{
		// You can provide more than one key to support key rotation. The SDK will try to decrypt the result with each key.
		{
			Key:       key,
			Algorithm: sealed.AlgorithmAES256GCM,
		},
	}
	unsealedResponse, err := sealed.UnsealEventsResponse(sealedResult, keys)

	if err != nil {
		panic(err)
	}

	// Do something with unsealed response, e.g: send it back to the frontend.
	fmt.Println(unsealedResponse)
}

```
To learn more, refer to example located in [example/sealedResults.go](./example/sealedResults.go).

## Documentation for API Endpoints

All URIs are relative to *https://api.fpjs.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FingerprintApi* | [**DeleteVisitorData**](docs/FingerprintApi.md#deletevisitordata) | **Delete** /visitors/{visitor_id} | Delete data by visitor ID
*FingerprintApi* | [**GetEvent**](docs/FingerprintApi.md#getevent) | **Get** /events/{request_id} | Get event by request ID
*FingerprintApi* | [**GetVisits**](docs/FingerprintApi.md#getvisits) | **Get** /visitors/{visitor_id} | Get visits by visitor ID

## Documentation For Models

 - [Asn](docs/Asn.md)
 - [BotdDetectionResult](docs/BotdDetectionResult.md)
 - [BotdResult](docs/BotdResult.md)
 - [BrowserDetails](docs/BrowserDetails.md)
 - [ClonedAppResult](docs/ClonedAppResult.md)
 - [Common403ErrorResponse](docs/Common403ErrorResponse.md)
 - [Confidence](docs/Confidence.md)
 - [DataCenter](docs/DataCenter.md)
 - [DeprecatedIpLocation](docs/DeprecatedIpLocation.md)
 - [DeprecatedIpLocationCity](docs/DeprecatedIpLocationCity.md)
 - [EmulatorResult](docs/EmulatorResult.md)
 - [ErrorCommon403Response](docs/ErrorCommon403Response.md)
 - [ErrorEvent404Response](docs/ErrorEvent404Response.md)
 - [ErrorEvent404ResponseError](docs/ErrorEvent404ResponseError.md)
 - [ErrorVisits403](docs/ErrorVisits403.md)
 - [ErrorVisitsDelete404Response](docs/ErrorVisitsDelete404Response.md)
 - [ErrorVisitsDelete404ResponseError](docs/ErrorVisitsDelete404ResponseError.md)
 - [EventResponse](docs/EventResponse.md)
 - [FactoryResetResult](docs/FactoryResetResult.md)
 - [FridaResult](docs/FridaResult.md)
 - [HighActivityResult](docs/HighActivityResult.md)
 - [IdentificationError](docs/IdentificationError.md)
 - [IncognitoResult](docs/IncognitoResult.md)
 - [IpBlockListResult](docs/IpBlockListResult.md)
 - [IpBlockListResultDetails](docs/IpBlockListResultDetails.md)
 - [IpInfoResult](docs/IpInfoResult.md)
 - [IpInfoResultV4](docs/IpInfoResultV4.md)
 - [IpInfoResultV6](docs/IpInfoResultV6.md)
 - [IpLocation](docs/IpLocation.md)
 - [IpLocationCity](docs/IpLocationCity.md)
 - [JailbrokenResult](docs/JailbrokenResult.md)
 - [Location](docs/Location.md)
 - [LocationSpoofingResult](docs/LocationSpoofingResult.md)
 - [ManyRequestsResponse](docs/ManyRequestsResponse.md)
 - [PrivacySettingsResult](docs/PrivacySettingsResult.md)
 - [ProductError](docs/ProductError.md)
 - [ProductsResponse](docs/ProductsResponse.md)
 - [ProductsResponseBotd](docs/ProductsResponseBotd.md)
 - [ProductsResponseIdentification](docs/ProductsResponseIdentification.md)
 - [ProductsResponseIdentificationData](docs/ProductsResponseIdentificationData.md)
 - [ProxyResult](docs/ProxyResult.md)
 - [Response](docs/Response.md)
 - [ResponseVisits](docs/ResponseVisits.md)
 - [RootAppsResult](docs/RootAppsResult.md)
 - [SeenAt](docs/SeenAt.md)
 - [SignalResponseClonedApp](docs/SignalResponseClonedApp.md)
 - [SignalResponseEmulator](docs/SignalResponseEmulator.md)
 - [SignalResponseFactoryReset](docs/SignalResponseFactoryReset.md)
 - [SignalResponseFrida](docs/SignalResponseFrida.md)
 - [SignalResponseHighActivity](docs/SignalResponseHighActivity.md)
 - [SignalResponseIncognito](docs/SignalResponseIncognito.md)
 - [SignalResponseIpBlocklist](docs/SignalResponseIpBlocklist.md)
 - [SignalResponseIpInfo](docs/SignalResponseIpInfo.md)
 - [SignalResponseJailbroken](docs/SignalResponseJailbroken.md)
 - [SignalResponseLocationSpoofing](docs/SignalResponseLocationSpoofing.md)
 - [SignalResponsePrivacySettings](docs/SignalResponsePrivacySettings.md)
 - [SignalResponseProxy](docs/SignalResponseProxy.md)
 - [SignalResponseRawDeviceAttributes](docs/SignalResponseRawDeviceAttributes.md)
 - [SignalResponseRootApps](docs/SignalResponseRootApps.md)
 - [SignalResponseSuspectScore](docs/SignalResponseSuspectScore.md)
 - [SignalResponseTampering](docs/SignalResponseTampering.md)
 - [SignalResponseTor](docs/SignalResponseTor.md)
 - [SignalResponseVirtualMachine](docs/SignalResponseVirtualMachine.md)
 - [SignalResponseVpn](docs/SignalResponseVpn.md)
 - [Subdivision](docs/Subdivision.md)
 - [SuspectScoreResult](docs/SuspectScoreResult.md)
 - [TamperingResult](docs/TamperingResult.md)
 - [TorResult](docs/TorResult.md)
 - [VirtualMachineResult](docs/VirtualMachineResult.md)
 - [Visit](docs/Visit.md)
 - [VpnResult](docs/VpnResult.md)
 - [VpnResultMethods](docs/VpnResultMethods.md)
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


## Documentation for sealed results

- [SealedResults](docs/SealedResults.md)
- [DecryptionKey](docs/DecryptionKey.md)

## Author

support@fingerprint.com

## Support and feedback

To report problems, ask questions or provide feedback, please use [Issues](https://github.com/fingerprintjs/fingerprintjs-pro-server-api-go-sdk/issues). If you need private support, you can email us at [oss-support@fingerprint.com](mailto:oss-support@fingerprint.com).

## License

This project is licensed under the [MIT license](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/blob/main/LICENSE).

