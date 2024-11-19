# Fingerprint Pro Server Go SDK

## 7.0.0

### Major Changes

- make `tag` field optional for Webhook ([406a373](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/406a3738e7f280a29bb691257c5b76ef936b387a))
- Provide error message from received response rather than status code text in `ApiError.Error()` ([76a47e1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/76a47e132a9fb85d271ae541c68a6727e8a20914))
- Change `ModelMap` to contain any possible property, not just strings ([f4a0749](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/f4a07499a69e758690313380a0138052ead5f327))
- Always throw `sdk.Error` from all Fingerprint API methods ([c02c7d9](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c02c7d9420f494cc080ada5c10b193c5bfc9ee1c))
- - Remove the `BrowserDetails` field `botProbability`.
  - Update the `IdentificationConfidence` field `score` type format: `float` -> `double`.
  - Make the `RawDeviceAttributeError` field `name` **optional** .
  - Make the `RawDeviceAttributeError` field `message` **optional** .
  - **events**: Remove the `EventsResponse` field `error`.
    - [note]: The errors are represented by `ErrorResponse` model.
  - **events**: Update the `HighActivity` field `dailyRequests` type format: `number` -> `int64`.
  - **events**: Specify the `Tampering` field `anomalyScore` type format: `double`.
  - **webhook**: Make the `Webhook` fields **optional**: `visitorId`, `visitorFound`, `firstSeenAt`, `lastSeenAt`, `browserDetails`, `incognito`.
  - **webhook**: Make the `WebhookClonedApp` field `result` **optional**.
  - **webhook**: Make the `WebhookDeveloperTools` field `result` **optional**.
  - **webhook**: Make the `WebhookEmulator` field `result` **optional**.
  - **webhook**: Make the `WebhookFactoryReset` fields `time` and `timestamp` **optional**.
  - **webhook**: Make the `WebhookFrida` field `result` **optional**.
  - **webhook**: Update the `WebhookHighActivity` field `dailyRequests` type format: `number` -> `int64`.
  - **webhook**: Make the `WebhookIPBlocklist` fields `result` and `details` **optional**.
  - **webhook**: Make the `WebhookJailbroken` field `result` **optional**.
  - **webhook**: Make the `WebhookLocationSpoofing` field `result` **optional**.
  - **webhook**: Make the `WebhookPrivacySettings` field `result` **optional**.
  - **webhook**: Make the `WebhookProxy` field `result` **optional**.
  - **webhook**: Make the `WebhookRemoteControl` field `result` **optional**.
  - **webhook**: Make the `WebhookRootApps` field `result` **optional**.
  - **webhook**: Make the `WebhookSuspectScore` field `result` **optional**.
  - **webhook**: Make the `WebhookTampering` fields `result`, `anomalyScore` and `antiDetectBrowser` **optional**.
  - **webhook**: Specify the `WebhookTampering` field `anomalyScore` type format: `double`.
  - **webhook**: Make the `WebhookTor` field `result` **optional**.
  - **webhook**: Make the `WebhookVelocity` fields **optional**: `distinctIp`, `distinctLinkedId`, `distinctCountry`, `events`, `ipEvents`, `distinctIpByLinkedId`, `distinctVisitorIdByLinkedId`.
  - **webhook**: Make the `WebhookVirtualMachine` field `result` **optional**.
  - **webhook**: Make the `WebhookVPN` fields **optional**: `result`, `confidence`, `originTimezone`, `methods`. ([c92322a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c92322a7254db73a80e7500bcadd1f66d2ec7b5c))
- - Rename `BotdResult` -> `Botd`.
  - Rename `BotdDetectionResult` -> `BotdBot`:
    - Extract `result` type as `BotdBotResult`.
  - Rename `ClonedAppResult` -> `ClonedApp`.
  - Rename `DeveloperToolsResult` -> `DeveloperTools`.
  - Rename `EmulatorResult` -> `Emulator`.
  - Refactor error models:
    - Remove `ErrorCommon403Response`, `ErrorCommon429Response`, `ErrorEvent404Response`, `TooManyRequestsResponse`, `ErrorVisits403`, `ErrorUpdateEvent400Response`, `ErrorUpdateEvent409Response`, `ErrorVisitor400Response`, `ErrorVisitor404Response`, `IdentificationError`, `ProductError`.
    - Introduce `ErrorResponse` and `ErrorPlainResponse`.
      - [note]: `ErrorPlainResponse` has a different format `{ "error": string }` and it is used only in `GET /visitors`.
    - Extract `error` type as `Error`.
    - Extract `error.code` type as `ErrorCode`.
  - Rename `EventResponse` -> `EventsGetResponse`.
  - Rename `EventUpdateRequest` -> `EventsUpdateRequest`.
  - Rename `FactoryResetResult` -> `FactoryReset`.
  - Rename `FridaResult` -> `Frida`.
  - Rename `IPLocation` -> `Geolocation`:
    - Rename `IPLocationCity` -> `GeolocationCity`.
    - Extract `subdivisions` type as `GeolocationSubdivisions`.
    - Rename `Location` -> `GeolocationContinent`:
    - Introduce a dedicated type `GeolocationCountry`.
    - Rename `Subdivision` -> `GeolocationSubdivision`.
  - Rename `HighActivityResult` -> `HighActivity`.
  - Rename `Confidence` -> `IdentificationConfidence`.
  - Rename `SeenAt` -> `IdentificationSeenAt`.
  - Rename `IncognitoResult` -> `Incognito`.
  - Rename `IpBlockListResult` -> `IPBlocklist`:
    - Extract `details` type as `IPBlocklistDetails`.
  - Rename `IpInfoResult` -> `IPInfo`:
    - Rename `IpInfoResultV4` -> `IPInfoV4`.
    - Rename `IpInfoResultV6` -> `IPInfoV6`.
    - Rename `ASN` -> `IPInfoASN`.
    - Rename `DataCenter` -> `IPInfoDataCenter`.
  - Rename `JailbrokenResult` -> `Jailbroken`.
  - Rename `LocationSpoofingResult` -> `LocationSpoofing`.
  - Rename `PrivacySettingsResult` -> `PrivacySettings`.
  - Rename `ProductsResponse` -> `Products`:
    - Rename inner types: `ProductsResponseIdentification` -> `ProductIdentification`, `ProductsResponseIdentificationData` -> `Identification`, `ProductsResponseBotd` -> `ProductBotd`, `SignalResponseRootApps` -> `ProductRootApps`, `SignalResponseEmulator` -> `ProductEmulator`, `SignalResponseIpInfo` -> `ProductIPInfo`, `SignalResponseIpBlocklist` -> `ProductIPBlocklist`, `SignalResponseTor` -> `ProductTor`, `SignalResponseVpn` -> `ProductVPN`, `SignalResponseProxy` -> `ProductProxy`, `ProxyResult` -> `Proxy`, `SignalResponseIncognito` -> `ProductIncognito`, `SignalResponseTampering` -> `ProductTampering`, `SignalResponseClonedApp` -> `ProductClonedApp`, `SignalResponseFactoryReset` -> `ProductFactoryReset`, `SignalResponseJailbroken` -> `ProductJailbroken`, `SignalResponseFrida` -> `ProductFrida`, `SignalResponsePrivacySettings` -> `ProductPrivacySettings`, `SignalResponseVirtualMachine` -> `ProductVirtualMachine`, `SignalResponseRawDeviceAttributes` -> `ProductRawDeviceAttributes`, `RawDeviceAttributesResultValue` -> `RawDeviceAttributes`, `SignalResponseHighActivity` -> `ProductHighActivity`, `SignalResponseLocationSpoofing` -> `ProductLocationSpoofing`, `SignalResponseSuspectScore` -> `ProductSuspectScore`, `SignalResponseRemoteControl` -> `ProductRemoteControl`, `SignalResponseVelocity` -> `ProductVelocity`, `SignalResponseDeveloperTools` -> `ProductDeveloperTools`.
    - Extract `identification.data` type as `Identification`.
  - Rename `RawDeviceAttributesResult` -> `RawDeviceAttributes`:
    - Extract item type as `RawDeviceAttribute`.
    - Extract `error` type as `RawDeviceAttributeError`.
  - Rename `RemoteControlResult` -> `RemoteControl`.
  - Rename `RootAppsResult` -> `RootApps`.
  - Rename `SuspectScoreResult` -> `SuspectScore`.
  - Extract new model `Tag`.
  - Rename `TamperingResult` -> `Tampering`.
  - Rename `TorResult` -> `Tor`.
  - Rename `VelocityResult` -> `Velocity`:
    - Rename `VelocityIntervals` -> `VelocityData`.
    - Rename `VelocityIntervalResult` -> `VelocityIntervals`.
  - Rename `VirtualMachineResult` -> `VirtualMachine`.
  - Rename the `Visit` field `ipLocation` type `DeprecatedIPLocation` -> `DeprecatedGeolocation`.
    - Instead of `DeprecatedIPLocationCity` use common `GeolocationCity`
  - Rename `Response` -> `VisitorsGetResponse`.
    - Omit extra inner type `ResponseVisits`
  - Rename `VpnResult` -> `VPN`.
    - Extract `confidence` type as `VPNConfidence`.
    - Extract `methods` type as `VPNMethods`.
  - Rename `WebhookVisit` -> `Webhook`.
    - Introduce new inner types: `WebhookRootApps`, `WebhookEmulator`, `WebhookIPInfo`, `WebhookIPBlocklist`, `WebhookTor`, `WebhookVPN`, `WebhookProxy`, `WebhookTampering`, `WebhookClonedApp`, `WebhookFactoryReset`, `WebhookJailbroken`, `WebhookFrida`, `WebhookPrivacySettings`, `WebhookVirtualMachine`, `WebhookRawDeviceAttributes`, `WebhookHighActivity`, `WebhookLocationSpoofing`, `WebhookSuspectScore`, `WebhookRemoteControl`, `WebhookVelocity`, `WebhookDeveloperTools`. ([c92322a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c92322a7254db73a80e7500bcadd1f66d2ec7b5c))
- Rename errors models related to visits:
  - rename `ErrorVisitsDelete400Response` to `ErrorVisitor400Response`
  - rename `ErrorVisitsDelete404ResponseError` to `ErrorVisitor404ResponseError`
  - rename `ErrorVisitsDelete404Response` to `ErrorVisitor404Response` ([fe7e9f7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/fe7e9f74c82347043e41e7755a97790e77ed35e2))

### Minor Changes

- **related-visitors**: Add GET `/related-visitors` endpoint ([e069c8f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e069c8fc35f6dcae9be924b69c42c2409f50c7d9))
- Added new `ipEvents`, `distinctIpByLinkedId`, and `distinctVisitorIdByLinkedId` fields to the `velocity` Smart Signal. ([c92322a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c92322a7254db73a80e7500bcadd1f66d2ec7b5c))
- - Make the `GeolocationCity` field `name` **required**.
  - Make the `GeolocationSubdivision` field `isoCode` **required**.
  - Make the `GeolocationSubdivision` field `name` **required**.
  - Make the `IPInfoASN` field `name` **required** .
  - Make the `IPInfoDataCenter` field `name` **required**.
  - Add **optional** `IdentificationConfidence` field `comment`.
  - **events**: Add **optional** `Botd` field `meta`.
  - **events**: Add **optional** `Identification` field `components`.
  - **events**: Make the `VPN` field `originCountry` **required**.
  - **visitors**: Add **optional** `Visit` field `components`.
  - **webhook**: Add **optional** `Webhook` field `components`. ([c92322a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c92322a7254db73a80e7500bcadd1f66d2ec7b5c))
- **visitors**: Add the confidence field to the VPN Detection Smart Signal ([782dc59](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/782dc59fa7c87835071c5b37e3d120e06a1591ef))
- Remove `ipv4` format from `ip` field in `Botd`, `Identification`, `Visit` and `Webhook` models. ([1bda1e3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1bda1e3027c51a8b27ecc9f87bccb0c5e09bcc39))
- **events**: Add `antiDetectBrowser` detection method to the `tampering` Smart Signal. ([e069c8f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e069c8fc35f6dcae9be924b69c42c2409f50c7d9))
- Provide ErrorCode in `Code()` method in `ApiError` ([76a47e1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/76a47e132a9fb85d271ae541c68a6727e8a20914))
- **events**: Introduce `PUT` endpoint for `/events` API ([e9df386](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e9df386d101da407e6bf16b59bca4f47d5690246))

### Patch Changes

- **related-visitors**: Add mention that the API is billable ([e069c8f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e069c8fc35f6dcae9be924b69c42c2409f50c7d9))
- Remove unused `Model` struct ([d398848](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/d39884830b68ebe5eb8a600b14bb862ac36a949d))

## 7.0.0-test.2

### Major Changes

- Provide error message from received response rather than status code text in `ApiError.Error()` ([76a47e1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/76a47e132a9fb85d271ae541c68a6727e8a20914))
- Always throw `sdk.Error` from all Fingerprint API methods ([c02c7d9](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c02c7d9420f494cc080ada5c10b193c5bfc9ee1c))

### Minor Changes

- Provide ErrorCode in `Code()` method in `ApiError` ([76a47e1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/76a47e132a9fb85d271ae541c68a6727e8a20914))

## 7.0.0-test.1

### Major Changes

- - Remove the `BrowserDetails` field `botProbability`.
  - Update the `IdentificationConfidence` field `score` type format: `float` -> `double`.
  - Make the `RawDeviceAttributeError` field `name` **optional** .
  - Make the `RawDeviceAttributeError` field `message` **optional** .
  - **events**: Remove the `EventsResponse` field `error`.
    - [note]: The errors are represented by `ErrorResponse` model.
  - **events**: Update the `HighActivity` field `dailyRequests` type format: `number` -> `int64`.
  - **events**: Specify the `Tampering` field `anomalyScore` type format: `double`.
  - **webhook**: Make the `Webhook` fields **optional**: `visitorId`, `visitorFound`, `firstSeenAt`, `lastSeenAt`, `browserDetails`, `incognito`.
  - **webhook**: Make the `WebhookClonedApp` field `result` **optional**.
  - **webhook**: Make the `WebhookDeveloperTools` field `result` **optional**.
  - **webhook**: Make the `WebhookEmulator` field `result` **optional**.
  - **webhook**: Make the `WebhookFactoryReset` fields `time` and `timestamp` **optional**.
  - **webhook**: Make the `WebhookFrida` field `result` **optional**.
  - **webhook**: Update the `WebhookHighActivity` field `dailyRequests` type format: `number` -> `int64`.
  - **webhook**: Make the `WebhookIPBlocklist` fields `result` and `details` **optional**.
  - **webhook**: Make the `WebhookJailbroken` field `result` **optional**.
  - **webhook**: Make the `WebhookLocationSpoofing` field `result` **optional**.
  - **webhook**: Make the `WebhookPrivacySettings` field `result` **optional**.
  - **webhook**: Make the `WebhookProxy` field `result` **optional**.
  - **webhook**: Make the `WebhookRemoteControl` field `result` **optional**.
  - **webhook**: Make the `WebhookRootApps` field `result` **optional**.
  - **webhook**: Make the `WebhookSuspectScore` field `result` **optional**.
  - **webhook**: Make the `WebhookTampering` fields `result`, `anomalyScore` and `antiDetectBrowser` **optional**.
  - **webhook**: Specify the `WebhookTampering` field `anomalyScore` type format: `double`.
  - **webhook**: Make the `WebhookTor` field `result` **optional**.
  - **webhook**: Make the `WebhookVelocity` fields **optional**: `distinctIp`, `distinctLinkedId`, `distinctCountry`, `events`, `ipEvents`, `distinctIpByLinkedId`, `distinctVisitorIdByLinkedId`.
  - **webhook**: Make the `WebhookVirtualMachine` field `result` **optional**.
  - **webhook**: Make the `WebhookVPN` fields **optional**: `result`, `confidence`, `originTimezone`, `methods`. ([c92322a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c92322a7254db73a80e7500bcadd1f66d2ec7b5c))
- - Rename `BotdResult` -> `Botd`.
  - Rename `BotdDetectionResult` -> `BotdBot`:
    - Extract `result` type as `BotdBotResult`.
  - Rename `ClonedAppResult` -> `ClonedApp`.
  - Rename `DeveloperToolsResult` -> `DeveloperTools`.
  - Rename `EmulatorResult` -> `Emulator`.
  - Refactor error models:
    - Remove `ErrorCommon403Response`, `ErrorCommon429Response`, `ErrorEvent404Response`, `TooManyRequestsResponse`, `ErrorVisits403`, `ErrorUpdateEvent400Response`, `ErrorUpdateEvent409Response`, `ErrorVisitor400Response`, `ErrorVisitor404Response`, `IdentificationError`, `ProductError`.
    - Introduce `ErrorResponse` and `ErrorPlainResponse`.
      - [note]: `ErrorPlainResponse` has a different format `{ "error": string }` and it is used only in `GET /visitors`.
    - Extract `error` type as `Error`.
    - Extract `error.code` type as `ErrorCode`.
  - Rename `EventResponse` -> `EventsGetResponse`.
  - Rename `EventUpdateRequest` -> `EventsUpdateRequest`.
  - Rename `FactoryResetResult` -> `FactoryReset`.
  - Rename `FridaResult` -> `Frida`.
  - Rename `IPLocation` -> `Geolocation`:
    - Rename `IPLocationCity` -> `GeolocationCity`.
    - Extract `subdivisions` type as `GeolocationSubdivisions`.
    - Rename `Location` -> `GeolocationContinent`:
    - Introduce a dedicated type `GeolocationCountry`.
    - Rename `Subdivision` -> `GeolocationSubdivision`.
  - Rename `HighActivityResult` -> `HighActivity`.
  - Rename `Confidence` -> `IdentificationConfidence`.
  - Rename `SeenAt` -> `IdentificationSeenAt`.
  - Rename `IncognitoResult` -> `Incognito`.
  - Rename `IpBlockListResult` -> `IPBlocklist`:
    - Extract `details` type as `IPBlocklistDetails`.
  - Rename `IpInfoResult` -> `IPInfo`:
    - Rename `IpInfoResultV4` -> `IPInfoV4`.
    - Rename `IpInfoResultV6` -> `IPInfoV6`.
    - Rename `ASN` -> `IPInfoASN`.
    - Rename `DataCenter` -> `IPInfoDataCenter`.
  - Rename `JailbrokenResult` -> `Jailbroken`.
  - Rename `LocationSpoofingResult` -> `LocationSpoofing`.
  - Rename `PrivacySettingsResult` -> `PrivacySettings`.
  - Rename `ProductsResponse` -> `Products`:
    - Rename inner types: `ProductsResponseIdentification` -> `ProductIdentification`, `ProductsResponseIdentificationData` -> `Identification`, `ProductsResponseBotd` -> `ProductBotd`, `SignalResponseRootApps` -> `ProductRootApps`, `SignalResponseEmulator` -> `ProductEmulator`, `SignalResponseIpInfo` -> `ProductIPInfo`, `SignalResponseIpBlocklist` -> `ProductIPBlocklist`, `SignalResponseTor` -> `ProductTor`, `SignalResponseVpn` -> `ProductVPN`, `SignalResponseProxy` -> `ProductProxy`, `ProxyResult` -> `Proxy`, `SignalResponseIncognito` -> `ProductIncognito`, `SignalResponseTampering` -> `ProductTampering`, `SignalResponseClonedApp` -> `ProductClonedApp`, `SignalResponseFactoryReset` -> `ProductFactoryReset`, `SignalResponseJailbroken` -> `ProductJailbroken`, `SignalResponseFrida` -> `ProductFrida`, `SignalResponsePrivacySettings` -> `ProductPrivacySettings`, `SignalResponseVirtualMachine` -> `ProductVirtualMachine`, `SignalResponseRawDeviceAttributes` -> `ProductRawDeviceAttributes`, `RawDeviceAttributesResultValue` -> `RawDeviceAttributes`, `SignalResponseHighActivity` -> `ProductHighActivity`, `SignalResponseLocationSpoofing` -> `ProductLocationSpoofing`, `SignalResponseSuspectScore` -> `ProductSuspectScore`, `SignalResponseRemoteControl` -> `ProductRemoteControl`, `SignalResponseVelocity` -> `ProductVelocity`, `SignalResponseDeveloperTools` -> `ProductDeveloperTools`.
    - Extract `identification.data` type as `Identification`.
  - Rename `RawDeviceAttributesResult` -> `RawDeviceAttributes`:
    - Extract item type as `RawDeviceAttribute`.
    - Extract `error` type as `RawDeviceAttributeError`.
  - Rename `RemoteControlResult` -> `RemoteControl`.
  - Rename `RootAppsResult` -> `RootApps`.
  - Rename `SuspectScoreResult` -> `SuspectScore`.
  - Extract new model `Tag`.
  - Rename `TamperingResult` -> `Tampering`.
  - Rename `TorResult` -> `Tor`.
  - Rename `VelocityResult` -> `Velocity`:
    - Rename `VelocityIntervals` -> `VelocityData`.
    - Rename `VelocityIntervalResult` -> `VelocityIntervals`.
  - Rename `VirtualMachineResult` -> `VirtualMachine`.
  - Rename the `Visit` field `ipLocation` type `DeprecatedIPLocation` -> `DeprecatedGeolocation`.
    - Instead of `DeprecatedIPLocationCity` use common `GeolocationCity`
  - Rename `Response` -> `VisitorsGetResponse`.
    - Omit extra inner type `ResponseVisits`
  - Rename `VpnResult` -> `VPN`.
    - Extract `confidence` type as `VPNConfidence`.
    - Extract `methods` type as `VPNMethods`.
  - Rename `WebhookVisit` -> `Webhook`.
    - Introduce new inner types: `WebhookRootApps`, `WebhookEmulator`, `WebhookIPInfo`, `WebhookIPBlocklist`, `WebhookTor`, `WebhookVPN`, `WebhookProxy`, `WebhookTampering`, `WebhookClonedApp`, `WebhookFactoryReset`, `WebhookJailbroken`, `WebhookFrida`, `WebhookPrivacySettings`, `WebhookVirtualMachine`, `WebhookRawDeviceAttributes`, `WebhookHighActivity`, `WebhookLocationSpoofing`, `WebhookSuspectScore`, `WebhookRemoteControl`, `WebhookVelocity`, `WebhookDeveloperTools`. ([c92322a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c92322a7254db73a80e7500bcadd1f66d2ec7b5c))

### Minor Changes

- **related-visitors**: Add GET `/related-visitors` endpoint ([e069c8f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e069c8fc35f6dcae9be924b69c42c2409f50c7d9))
- Added new `ipEvents`, `distinctIpByLinkedId`, and `distinctVisitorIdByLinkedId` fields to the `velocity` Smart Signal. ([c92322a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c92322a7254db73a80e7500bcadd1f66d2ec7b5c))
- - Make the `GeolocationCity` field `name` **required**.
  - Make the `GeolocationSubdivision` field `isoCode` **required**.
  - Make the `GeolocationSubdivision` field `name` **required**.
  - Make the `IPInfoASN` field `name` **required** .
  - Make the `IPInfoDataCenter` field `name` **required**.
  - Add **optional** `IdentificationConfidence` field `comment`.
  - **events**: Add **optional** `Botd` field `meta`.
  - **events**: Add **optional** `Identification` field `components`.
  - **events**: Make the `VPN` field `originCountry` **required**.
  - **visitors**: Add **optional** `Visit` field `components`.
  - **webhook**: Add **optional** `Webhook` field `components`. ([c92322a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c92322a7254db73a80e7500bcadd1f66d2ec7b5c))
- **visitors**: Add the confidence field to the VPN Detection Smart Signal ([782dc59](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/782dc59fa7c87835071c5b37e3d120e06a1591ef))
- Remove `ipv4` format from `ip` field in `Botd`, `Identification`, `Visit` and `Webhook` models. ([1bda1e3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1bda1e3027c51a8b27ecc9f87bccb0c5e09bcc39))
- **events**: Add `antiDetectBrowser` detection method to the `tampering` Smart Signal. ([e069c8f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e069c8fc35f6dcae9be924b69c42c2409f50c7d9))

### Patch Changes

- **related-visitors**: Add mention that the API is billable ([e069c8f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e069c8fc35f6dcae9be924b69c42c2409f50c7d9))

## 7.0.0-test.0

### Major Changes

- make `tag` field optional for Webhook ([406a373](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/406a3738e7f280a29bb691257c5b76ef936b387a))
- Change `ModelMap` to contain any possible property, not just strings ([f4a0749](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/f4a07499a69e758690313380a0138052ead5f327))
- Rename errors models related to visits:
  - rename `ErrorVisitsDelete400Response` to `ErrorVisitor400Response`
  - rename `ErrorVisitsDelete404ResponseError` to `ErrorVisitor404ResponseError`
  - rename `ErrorVisitsDelete404Response` to `ErrorVisitor404Response` ([fe7e9f7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/fe7e9f74c82347043e41e7755a97790e77ed35e2))

### Minor Changes

- **events**: Introduce `PUT` endpoint for `/events` API ([e9df386](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e9df386d101da407e6bf16b59bca4f47d5690246))

### Patch Changes

- Remove unused `Model` struct ([d398848](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/d39884830b68ebe5eb8a600b14bb862ac36a949d))

## [6.1.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v6.0.0...v6.1.0) (2024-07-30)

### Features

- add velocity, remote control and developer tools smart signals ([a66f05c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a66f05cc3b743f3ddf3467bc79afaa9311cf2073))

## [6.0.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v5.0.2...v6.0.0) (2024-06-27)

### ⚠ BREAKING CHANGES

- it is now easier to check for too many requests error (429):

```go
	response, httpRes, err := client.FingerprintApi.GetVisits(auth, visitorId, &opts)
	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		var tooManyRequestsError *sdk.TooManyRequestsError

		if errors.As(err, &tooManyRequestsError) {
			log.Printf("Too many requests, retry after %d seconds", tooManyRequestsError.RetryAfter())
		} else {
			log.Print(err)
		}
	}
```

- rename `GenericSwaggerError` to `ApiError`
- rename `ManyRequestsResponse` to `TooManyRequestsResponse`
- go 1.20 has reached EOL. Minimal supported version of go is now 1.21
- right now we use native `errors` package for joining errors, meaning that multiple error messages are now joined by new line rather than colon (:)
- optional pkg is no longer used in this SDK. Please pass native GO types instead.

### Features

- add `IsValidWebhookSignature` function for validating webhook signature ([a5bf13d](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a5bf13d62171ce06ec031e26a33d27d3f0b851bb))
- add delete API ([0e077c3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/0e077c3546c4a29d4ca8ae42da2eff6c587fee6f))
- add os Mismatch ([30b0215](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/30b0215128f488db116ce29e8c531cbb8718eafb))
- add revision string field to confidence object ([8a2f270](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/8a2f270a3cc057dec78bf7b4aaa36522ca960d9c))
- drop support for go 1.20 ([46953bc](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/46953bc345e5d6c8acdb215c2b665b54c91fe5a8))
- drop usage of `github.com/pkg/errors` ([186d30a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/186d30a6ae0135a02f6abd20e0746ee59bee024a))
- introduce `TooManyRequestsError` ([85f3307](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/85f3307c19991a0cb716a6c8f313ed7fa83bccf9))
- provide `HttpResponse()` in `ApiError` ([acd1274](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/acd1274a1e7b157f0b3dad1ff1d9b53a2730b4e6))
- re-write request handling logic ([14b7e7f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/14b7e7f2f26280e5ffe59269273995b17ab19126))
- remove usage of github.com/antihax/optional package ([62db97f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/62db97f9373b7bf929cee3f8b5fccb50d8b82bd8))
- rename `GenericSwaggerError` to `ApiError` ([259b7b4](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/259b7b417c7dcab31021e67aaa34e8daa3e41d82))
- rename `ManyRequestsResponse` to `TooManyRequestsResponse` ([3f66641](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/3f66641c610d34db4cd0657833ff2696948f2f3b))

### Bug Fixes

- allow passing `nil` configuration to `NewAPIClient` ([8234fbe](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/8234fbef2bd91cbc4b62ac7061b5c907759e9527))
- move test related dependencies to test module ([298275d](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/298275d75d13298745304916425539f905b0225f))
- use correct error type for `incognito`, `rawDeviceAttributes` and `tampering` in the `GetEvent` method ([c29aea9](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c29aea98fb7f814f19b46225a5656318fc4f81fd))

### Reverts

- Revert "chore(release): 6.0.0-test.1 [skip ci]" ([9916b45](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/9916b459c455951a6247636a9bbccc7cecc3b285))
- "chore(release): 6.0.0-test.1 [skip ci]" ([84ec138](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/84ec138b42fe10a7ffe649c031692da68e2cfaff))

## [5.0.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v5.0.1...v5.0.2) (2024-03-28)

### Build System

- **deps:** bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([b749ff0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/b749ff0c678b761c9994290cdf1937033a81cabb))
- **deps:** bump google.golang.org/protobuf in /example ([2fd6964](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/2fd6964c4712d5ade810f96ad91ed0666df784d3))

## [5.0.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v5.0.0...v5.0.1) (2024-02-27)

### Bug Fixes

- fix version references after a major release ([df759d6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/df759d6f10b10903535dcca4b2ebdf20210d8bf7))

### Documentation

- **README:** update readme requirements section ([4c4776b](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/4c4776b0c85bd4cf6392b72058439f47a6925598))

### Build System

- **deps:** bump go version to 1.20 ([469f36a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/469f36a6c25653f6139ca4db06296a1e5dc31980))
- **deps:** dump go to 1.20 in examples ([f88ae65](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/f88ae65521c451d5d76b108a553480e7072bd91b))
- **deps:** update dependencies to latest versions ([ae0189b](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/ae0189b4b0fa7b9d3a8250b744b5cc55f05fdc7a))
- **deps:** update example project dependencies ([6944d46](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/6944d4662b61a506325c3c5fe1ab7673545d9959))
- **deps:** update project dependencies to latest versions ([b2332a6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/b2332a6d92e1e4fdf0192e6b5160c80462dcb27c))

## [5.0.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v4.1.0...v5.0.0) (2024-02-27)

### ⚠ BREAKING CHANGES

- change models for the most smart signals
- make identification field `confidence` optional
- deprecated `ipLocation` field uses `DeprecatedIpLocation` model

### Features

- add `linkedId` field to the `BotdResult` type ([f3dec04](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/f3dec04387eb9f8c1efd889cb2ab28c7822479b2))
- add `SuspectScore` smart signal support ([a6fe1a5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a6fe1a5090109f97514ad2d3db7263080b19ad9b))
- add missed errors structures ([903bf6b](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/903bf6b8fd4a931c80ebae1998ad78604343be06))
- fix `ipLocation` deprecation ([ec59bc6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/ec59bc6e581f623c6675403d9bf215a811f07a73))
- make identification field `tag` required ([b6e841e](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/b6e841e7ef77830f281e2b73e53b36711d86c4d4))
- update `originCountry` field to the`vpn` signal ([6ce55a7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/6ce55a7a6c65c7ec02a7dbb832694393cf7e8b57))
- use shared structures for webhooks and event ([01c1132](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/01c11328af7d5cc10dfa0c01047d6318baf0a24c))

### Bug Fixes

- make fields required according to real API response ([a1c7578](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a1c757859632107bec1e5bd6212185f1d462c417))

## [4.1.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v4.0.0...v4.1.0) (2024-01-31)

### Features

- add method for decoding sealed results ([5ed5c5b](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/5ed5c5bb7222727f4816e4e7a4a7cc62b8a055de))

### Bug Fixes

- update module to v4 ([be9c14e](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/be9c14ecfa6ed869f03cac0c4a4de2d641af5217))

## [4.0.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.5.0...v4.0.0) (2024-01-12)

### ⚠ BREAKING CHANGES

- `IpInfo` field `DataCenter` renamed to `Datacenter`

### Features

- deprecate `IPLocation` ([3d142eb](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/3d142eb82f9bbd9267e5b068fbd30f69e8606dd0))
- use `datacenter` instead of the wrong `dataCenter` ([c1d0c01](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c1d0c0134b984242bb6c026f9c3b10e8582a7a2f))

## [3.5.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.4.2...v3.5.0) (2023-11-27)

### Features

- add `highActivity` and `locationSpoofing` signals, support `originTimezone` for `vpn` signal ([81cc2ab](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/81cc2ab49b7910dc6752b468fa56224f0fd810ab))

### Documentation

- **README:** mention license ([61d5a6a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/61d5a6a189d20a882acbecbcc8e20b07d39cc464))

## [3.4.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.4.1...v3.4.2) (2023-09-20)

### Bug Fixes

- update OpenAPI Schema with `asn` and `dataCenter` signals ([0164fe0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/0164fe009898afb42068fc28f4f7084a72dc27de))
- update OpenAPI Schema with `auxiliaryMobile` method for VPN signal ([193b787](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/193b787ae6378c71bc6da82842afdd53af972894))

## [3.4.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.4.0...v3.4.1) (2023-08-25)

### Build System

- **deps:** bump golang.org/x/net ([4b21e0b](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/4b21e0bae4ceb181310024f510463d4b4c2c0339))

## [3.4.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.3.0...v3.4.0) (2023-07-31)

### Features

- add raw device attributes ([17cac0f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/17cac0fd1fa3bd08ebe472bd31a143d814f4e046))

## [3.3.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.2.0...v3.3.0) (2023-07-14)

### Features

- add smart signals support ([17e5854](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/17e5854d90a40641379b0b77839f2d3f47fbc763))

## [3.2.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.1.0...v3.2.0) (2023-06-06)

### Features

- update schema with correct `IpLocation` format and doc updates ([e3b5f78](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e3b5f789b85863bcc81d342878331c870b58f44d))

### Bug Fixes

- fix backtick problem in comments and documentation ([0063c75](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/0063c751b61c8d3990e2e6fbe5c27fd13d3c299f))

## [3.1.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.0.1...v3.1.0) (2023-05-11)

### Features

- update schema and add more signals ([8a7b0c3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/8a7b0c3705bd3ae310b2278048868699e3137b99))

### Bug Fixes

- update schema with correct Webhook Signals description ([54f2085](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/54f2085449eb172ca8db511f2ad62051640101fd))
- update schema, add test for undescribed fields case ([2d071a9](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/2d071a9adfc46dd71381368e8f1e554e4f5e9e94))

## [3.0.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.0.0...v3.0.1) (2023-01-30)

### Bug Fixes

- bump version in module name to v3 ([3988bf6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/3988bf62a50aaa48acaa27b58e12997a021a48d0))

## [3.0.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v2.0.0...v3.0.0) (2023-01-30)

### ⚠ BREAKING CHANGES

- changed `before` parameter type from `int32` to `int64`

### Features

- change `before` parameter type in /visits endpoint ([436f3bf](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/436f3bfa14bbd6a4f31eaecc477324b5c0023352))

### Documentation

- **README:** fix invalid install command ([fbb1769](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/fbb1769287dc83b89848455df35943fde8567b70))

## [2.0.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.2.0...v2.0.0) (2023-01-23)

### ⚠ BREAKING CHANGES

- `StSeenAt` type renamed to `SeenAt`

### Features

- generate new source file with updated swagger ([1d94e69](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1d94e698850a2f753b7f1398cdd667a6ae5aea10))
- introduce identification error into EventsResponse ([925334e](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/925334e52b4046b97d6b814f734b81ff2086fee7))
- store RetryAfter in TooManyRequestsResponse ([8239e3c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/8239e3cae5e6246440f4ab76fcd605fb78aa50ab))
- Update list of examples in generate.go (new errors) ([a328ad6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a328ad6defbeff6cedc6769b87595fb18f56ba9e))
- update module name to github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v2 ([aec4af5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/aec4af5bd7ec823dabbbf10ef77203c7881079a0))

### Documentation

- **README:** update referenced module name ([78f5dac](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/78f5dac76507a993ce4f553ece7b4ceb5c39d67f))

## [2.0.0-test.3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v2.0.0-test.2...v2.0.0-test.3) (2023-01-23)

### Features

- update module name to github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v2 ([aec4af5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/aec4af5bd7ec823dabbbf10ef77203c7881079a0))

## [2.0.0-test.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v2.0.0-test.1...v2.0.0-test.2) (2023-01-23)

### Features

- introduce identification error into EventsResponse ([925334e](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/925334e52b4046b97d6b814f734b81ff2086fee7))
- store RetryAfter in TooManyRequestsResponse ([8239e3c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/8239e3cae5e6246440f4ab76fcd605fb78aa50ab))

## [2.0.0-test.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.2.0...v2.0.0-test.1) (2023-01-18)

### ⚠ BREAKING CHANGES

- `StSeenAt` type renamed to `SeenAt`

### Features

- generate new source file with updated swagger ([1d94e69](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1d94e698850a2f753b7f1398cdd667a6ae5aea10))
- Update list of examples in generate.go (new errors) ([a328ad6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a328ad6defbeff6cedc6769b87595fb18f56ba9e))

## [1.2.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.1.0...v1.2.0) (2022-10-24)

### Features

- update schema to support url field for botd result ([5e0ec6c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/5e0ec6c9c65ec79e20dfbb062c6a7471215852cd))

### Documentation

- **README:** add different region to code example ([3986d6d](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/3986d6df1306666bb03812be05c408ed91ecf0d9))
- **README:** add region section ([a2342cd](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a2342cdc0451982ee5c33bd46704d193a263ddd1))

## [1.1.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.2...v1.1.0) (2022-09-19)

### Features

- introduce /event/{request_id} endpoint ([74a39b6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/74a39b609b64ef2f9b7eae76972d7e4532b1867b))

## [1.0.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.1...v1.0.2) (2022-09-01)

### Documentation

- **README:** update template ([0bb3917](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/0bb391711ec3625af7c8ffb2de6bdc525758fbf1))

## [1.0.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0...v1.0.1) (2022-09-01)

### Documentation

- **README:** remove WIP label ([5d910ae](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/5d910ae9a0d43e19647d5982eefec536502f616f))

## 1.0.0 (2022-09-01)

### Features

- add "integrationsInfo" query param ([b326815](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/b326815f69b92c3c1d2d691a99c8483753ec6e49))
- create Go SDK ([a5e03b5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a5e03b5b1ad5e58441d88faf992f5f6e08033d55))
- support passing region ([1ba2e94](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1ba2e941ae8fe65abd706f7e5506953b03cde9ab))

### Bug Fixes

- send API key only in headers ([92a4f88](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/92a4f886b6876d878d9c7ca61f6b4e3af34445d6))
- support nil values for time.Time ([459ba4c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/459ba4c8dde6c6e1428fdeb9b0c2975de1a2f1d6))
- use config.json as single source of truth ([519f0d7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/519f0d7b0c7c84fc164c4cf71440a83c87ab6239))

### Documentation

- **README:** fix installation cmd typo ([2017b4c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/2017b4c890cf7d0cb6b9dd1df5a374a8af2e96a4))
- **README:** remove unnecessary import from example ([e6759e7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e6759e71e712cff1508fbfc88a941e68244bbd66))
- **README:** update readme ([ae4e0ea](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/ae4e0ea67c95598f3771cd1e7c89189bab17793e))

## [1.0.0-test.5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0-test.4...v1.0.0-test.5) (2022-08-29)

### Bug Fixes

- send API key only in headers ([92a4f88](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/92a4f886b6876d878d9c7ca61f6b4e3af34445d6))
- support nil values for time.Time ([459ba4c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/459ba4c8dde6c6e1428fdeb9b0c2975de1a2f1d6))

### Documentation

- **README:** fix installation cmd typo ([2017b4c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/2017b4c890cf7d0cb6b9dd1df5a374a8af2e96a4))
- **README:** remove unnecessary import from example ([e6759e7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e6759e71e712cff1508fbfc88a941e68244bbd66))

## [1.0.0-test.4](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0-test.3...v1.0.0-test.4) (2022-08-25)

### Documentation

- **README:** update readme ([ae4e0ea](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/ae4e0ea67c95598f3771cd1e7c89189bab17793e))

## [1.0.0-test.3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0-test.2...v1.0.0-test.3) (2022-08-24)

### Bug Fixes

- use config.json as single source of truth ([519f0d7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/519f0d7b0c7c84fc164c4cf71440a83c87ab6239))

## [1.0.0-test.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0-test.1...v1.0.0-test.2) (2022-08-19)

### Features

- add "integrationsInfo" query param ([b326815](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/b326815f69b92c3c1d2d691a99c8483753ec6e49))

## 1.0.0-test.1 (2022-08-19)

### Features

- create Go SDK ([a5e03b5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a5e03b5b1ad5e58441d88faf992f5f6e08033d55))
- support passing region ([1ba2e94](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1ba2e941ae8fe65abd706f7e5506953b03cde9ab))
