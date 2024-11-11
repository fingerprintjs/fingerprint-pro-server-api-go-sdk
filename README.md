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
- Package version: 7.0.0-test.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Requirements

Go Lang 1.21 or higher

We keep the [Go support policy](https://go.dev/doc/devel/release) and support the last two major versions of Go.

## Installation & Usage

1. Get the package from GitHub:
```shell
go get github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk
```

2. Import and use the library:

```go
package main

import (
	"context"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
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
	visitorId := "<VISITOR_ID>"
	requestId := "<REQUEST_ID>"
	opts := sdk.FingerprintApiGetVisitsOpts{
		RequestId: requestId,
	}
	visits, httpRes, err := client.FingerprintApi.GetVisits(auth, visitorId, &opts)
	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		var tooManyRequestsError *sdk.TooManyRequestsError

		if errors.As(err, &tooManyRequestsError) {
			log.Fatalf("Too many requests, retry after %d seconds", tooManyRequestsError.RetryAfter())
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("Got response with visitorId: %s", visits.VisitorId)
	
	event, httpRes, err := client.FingerprintApi.GetEvent(auth, requestId)
	if err != nil {
		log.Fatal(err)
    }

	if event.Products.Identification != nil {
		fmt.Printf("Got response with Identification: %v", event.Products.Identification)

	}
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
    "github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
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
go get github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk/sealed
```
Then you can use below code to unseal results:
```go
package main

import (
	"encoding/base64"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk/sealed"
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

## Webhook signing

This SDK provides utility method for verifing the HMAC signature of the incoming [webhook](https://dev.fingerprint.com/docs/webhooks) request.
Install the webhook dependency as below:
```shell
go get github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk/webhook
```

Then you can use below code to verify signature:
```go
package main

import (
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk/webhook"
)

func main() {
    // Your webhook signing secret.
    secret := "secret"

    // Request data. In real life scenario this will be the body of incoming request
    data := []byte("data")

    // Value of the "fpjs-event-signature" header.
    header := "v1=1b2c16b75bd2a870c114153ccda5bcfca63314bc722fa160d690de133ccbb9db"

    isValid := webhook.IsValidWebhookSignature(header, data, secret)

    if !isValid {
        panic("Invalid signature")
    }
}
```


To learn more, refer to example located in [example/sealedResults.go](./example/sealedResults.go).

## Documentation for API Endpoints

All URIs are relative to *https://api.fpjs.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FingerprintApi* | [**DeleteVisitorData**](docs/FingerprintApi.md#deletevisitordata) | **Delete** /visitors/{visitor_id} | Delete data by visitor ID
*FingerprintApi* | [**GetEvent**](docs/FingerprintApi.md#getevent) | **Get** /events/{request_id} | Get event by request ID
*FingerprintApi* | [**GetRelatedVisitors**](docs/FingerprintApi.md#getrelatedvisitors) | **Get** /related-visitors | Get Related Visitors
*FingerprintApi* | [**GetVisits**](docs/FingerprintApi.md#getvisits) | **Get** /visitors/{visitor_id} | Get visits by visitor ID
*FingerprintApi* | [**UpdateEvent**](docs/FingerprintApi.md#updateevent) | **Put** /events/{request_id} | Update an event with a given request ID

## Documentation For Models

 - [Botd](docs/Botd.md)
 - [BotdBot](docs/BotdBot.md)
 - [BotdBotResult](docs/BotdBotResult.md)
 - [BrowserDetails](docs/BrowserDetails.md)
 - [ClonedApp](docs/ClonedApp.md)
 - [DeprecatedGeolocation](docs/DeprecatedGeolocation.md)
 - [DeveloperTools](docs/DeveloperTools.md)
 - [Emulator](docs/Emulator.md)
 - [ErrorCode](docs/ErrorCode.md)
 - [ErrorPlainResponse](docs/ErrorPlainResponse.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [EventsGetResponse](docs/EventsGetResponse.md)
 - [EventsUpdateRequest](docs/EventsUpdateRequest.md)
 - [FactoryReset](docs/FactoryReset.md)
 - [Frida](docs/Frida.md)
 - [Geolocation](docs/Geolocation.md)
 - [GeolocationCity](docs/GeolocationCity.md)
 - [GeolocationContinent](docs/GeolocationContinent.md)
 - [GeolocationCountry](docs/GeolocationCountry.md)
 - [GeolocationSubdivision](docs/GeolocationSubdivision.md)
 - [HighActivity](docs/HighActivity.md)
 - [Identification](docs/Identification.md)
 - [IdentificationConfidence](docs/IdentificationConfidence.md)
 - [IdentificationSeenAt](docs/IdentificationSeenAt.md)
 - [Incognito](docs/Incognito.md)
 - [IpBlocklist](docs/IpBlocklist.md)
 - [IpBlocklistDetails](docs/IpBlocklistDetails.md)
 - [IpInfo](docs/IpInfo.md)
 - [IpInfoAsn](docs/IpInfoAsn.md)
 - [IpInfoDataCenter](docs/IpInfoDataCenter.md)
 - [IpInfoV4](docs/IpInfoV4.md)
 - [IpInfoV6](docs/IpInfoV6.md)
 - [Jailbroken](docs/Jailbroken.md)
 - [LocationSpoofing](docs/LocationSpoofing.md)
 - [ModelError](docs/ModelError.md)
 - [PrivacySettings](docs/PrivacySettings.md)
 - [ProductBotd](docs/ProductBotd.md)
 - [ProductClonedApp](docs/ProductClonedApp.md)
 - [ProductDeveloperTools](docs/ProductDeveloperTools.md)
 - [ProductEmulator](docs/ProductEmulator.md)
 - [ProductFactoryReset](docs/ProductFactoryReset.md)
 - [ProductFrida](docs/ProductFrida.md)
 - [ProductHighActivity](docs/ProductHighActivity.md)
 - [ProductIdentification](docs/ProductIdentification.md)
 - [ProductIncognito](docs/ProductIncognito.md)
 - [ProductIpBlocklist](docs/ProductIpBlocklist.md)
 - [ProductIpInfo](docs/ProductIpInfo.md)
 - [ProductJailbroken](docs/ProductJailbroken.md)
 - [ProductLocationSpoofing](docs/ProductLocationSpoofing.md)
 - [ProductPrivacySettings](docs/ProductPrivacySettings.md)
 - [ProductProxy](docs/ProductProxy.md)
 - [ProductRawDeviceAttributes](docs/ProductRawDeviceAttributes.md)
 - [ProductRemoteControl](docs/ProductRemoteControl.md)
 - [ProductRootApps](docs/ProductRootApps.md)
 - [ProductSuspectScore](docs/ProductSuspectScore.md)
 - [ProductTampering](docs/ProductTampering.md)
 - [ProductTor](docs/ProductTor.md)
 - [ProductVelocity](docs/ProductVelocity.md)
 - [ProductVirtualMachine](docs/ProductVirtualMachine.md)
 - [ProductVpn](docs/ProductVpn.md)
 - [Products](docs/Products.md)
 - [Proxy](docs/Proxy.md)
 - [RawDeviceAttribute](docs/RawDeviceAttribute.md)
 - [RawDeviceAttributeError](docs/RawDeviceAttributeError.md)
 - [RelatedVisitor](docs/RelatedVisitor.md)
 - [RelatedVisitorsResponse](docs/RelatedVisitorsResponse.md)
 - [RemoteControl](docs/RemoteControl.md)
 - [RootApps](docs/RootApps.md)
 - [SuspectScore](docs/SuspectScore.md)
 - [Tampering](docs/Tampering.md)
 - [Tor](docs/Tor.md)
 - [Velocity](docs/Velocity.md)
 - [VelocityData](docs/VelocityData.md)
 - [VelocityIntervals](docs/VelocityIntervals.md)
 - [VirtualMachine](docs/VirtualMachine.md)
 - [Visit](docs/Visit.md)
 - [VisitorsGetResponse](docs/VisitorsGetResponse.md)
 - [Vpn](docs/Vpn.md)
 - [VpnConfidence](docs/VpnConfidence.md)
 - [VpnMethods](docs/VpnMethods.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookClonedApp](docs/WebhookClonedApp.md)
 - [WebhookDeveloperTools](docs/WebhookDeveloperTools.md)
 - [WebhookEmulator](docs/WebhookEmulator.md)
 - [WebhookFactoryReset](docs/WebhookFactoryReset.md)
 - [WebhookFrida](docs/WebhookFrida.md)
 - [WebhookHighActivity](docs/WebhookHighActivity.md)
 - [WebhookIpBlocklist](docs/WebhookIpBlocklist.md)
 - [WebhookIpInfo](docs/WebhookIpInfo.md)
 - [WebhookJailbroken](docs/WebhookJailbroken.md)
 - [WebhookLocationSpoofing](docs/WebhookLocationSpoofing.md)
 - [WebhookPrivacySettings](docs/WebhookPrivacySettings.md)
 - [WebhookProxy](docs/WebhookProxy.md)
 - [WebhookRemoteControl](docs/WebhookRemoteControl.md)
 - [WebhookRootApps](docs/WebhookRootApps.md)
 - [WebhookSuspectScore](docs/WebhookSuspectScore.md)
 - [WebhookTampering](docs/WebhookTampering.md)
 - [WebhookTor](docs/WebhookTor.md)
 - [WebhookVelocity](docs/WebhookVelocity.md)
 - [WebhookVirtualMachine](docs/WebhookVirtualMachine.md)
 - [WebhookVpn](docs/WebhookVpn.md)

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

## Documentation for webhooks

- [DecryptionKey](docs/Webhook.md)

## Author

support@fingerprint.com

## Support and feedback

To report problems, ask questions, or provide feedback, please use [Issues](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/issues). If you need private support, you can email us at [oss-support@fingerprint.com](mailto:oss-support@fingerprint.com).

## License

This project is licensed under the [MIT license](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/blob/main/LICENSE).

