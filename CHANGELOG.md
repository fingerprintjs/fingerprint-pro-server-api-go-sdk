## [4.1.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v4.0.0...v4.1.0) (2024-01-31)


### Features

* add method for decoding sealed results ([5ed5c5b](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/5ed5c5bb7222727f4816e4e7a4a7cc62b8a055de))


### Bug Fixes

* update module to v4 ([be9c14e](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/be9c14ecfa6ed869f03cac0c4a4de2d641af5217))

## [4.0.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.5.0...v4.0.0) (2024-01-12)


### ⚠ BREAKING CHANGES

* `IpInfo` field `DataCenter` renamed to `Datacenter`

### Features

* deprecate `IPLocation` ([3d142eb](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/3d142eb82f9bbd9267e5b068fbd30f69e8606dd0))
* use `datacenter` instead of the wrong `dataCenter` ([c1d0c01](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/c1d0c0134b984242bb6c026f9c3b10e8582a7a2f))

## [3.5.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.4.2...v3.5.0) (2023-11-27)


### Features

* add `highActivity` and `locationSpoofing` signals, support `originTimezone` for `vpn` signal ([81cc2ab](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/81cc2ab49b7910dc6752b468fa56224f0fd810ab))


### Documentation

* **README:** mention license ([61d5a6a](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/61d5a6a189d20a882acbecbcc8e20b07d39cc464))

## [3.4.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.4.1...v3.4.2) (2023-09-20)


### Bug Fixes

* update OpenAPI Schema with `asn` and `dataCenter` signals ([0164fe0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/0164fe009898afb42068fc28f4f7084a72dc27de))
* update OpenAPI Schema with `auxiliaryMobile` method for VPN signal ([193b787](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/193b787ae6378c71bc6da82842afdd53af972894))

## [3.4.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.4.0...v3.4.1) (2023-08-25)


### Build System

* **deps:** bump golang.org/x/net ([4b21e0b](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/4b21e0bae4ceb181310024f510463d4b4c2c0339))

## [3.4.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.3.0...v3.4.0) (2023-07-31)


### Features

* add raw device attributes ([17cac0f](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/17cac0fd1fa3bd08ebe472bd31a143d814f4e046))

## [3.3.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.2.0...v3.3.0) (2023-07-14)


### Features

* add smart signals support ([17e5854](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/17e5854d90a40641379b0b77839f2d3f47fbc763))

## [3.2.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.1.0...v3.2.0) (2023-06-06)


### Features

* update schema with correct `IpLocation` format and doc updates ([e3b5f78](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e3b5f789b85863bcc81d342878331c870b58f44d))


### Bug Fixes

* fix backtick problem in comments and documentation ([0063c75](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/0063c751b61c8d3990e2e6fbe5c27fd13d3c299f))

## [3.1.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.0.1...v3.1.0) (2023-05-11)


### Features

* update schema and add more signals ([8a7b0c3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/8a7b0c3705bd3ae310b2278048868699e3137b99))


### Bug Fixes

* update schema with correct Webhook Signals description ([54f2085](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/54f2085449eb172ca8db511f2ad62051640101fd))
* update schema, add test for undescribed fields case ([2d071a9](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/2d071a9adfc46dd71381368e8f1e554e4f5e9e94))

## [3.0.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v3.0.0...v3.0.1) (2023-01-30)


### Bug Fixes

* bump version in module name to v3 ([3988bf6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/3988bf62a50aaa48acaa27b58e12997a021a48d0))

## [3.0.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v2.0.0...v3.0.0) (2023-01-30)


### ⚠ BREAKING CHANGES

* changed `before` parameter type from `int32` to `int64`

### Features

* change `before` parameter type in /visits endpoint ([436f3bf](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/436f3bfa14bbd6a4f31eaecc477324b5c0023352))


### Documentation

* **README:** fix invalid install command ([fbb1769](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/fbb1769287dc83b89848455df35943fde8567b70))

## [2.0.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.2.0...v2.0.0) (2023-01-23)


### ⚠ BREAKING CHANGES

* `StSeenAt` type renamed to `SeenAt`

### Features

* generate new source file with updated swagger ([1d94e69](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1d94e698850a2f753b7f1398cdd667a6ae5aea10))
* introduce identification error into EventsResponse ([925334e](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/925334e52b4046b97d6b814f734b81ff2086fee7))
* store RetryAfter in ManyRequestsResponse ([8239e3c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/8239e3cae5e6246440f4ab76fcd605fb78aa50ab))
* Update list of examples in generate.go (new errors) ([a328ad6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a328ad6defbeff6cedc6769b87595fb18f56ba9e))
* update module name to github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v2 ([aec4af5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/aec4af5bd7ec823dabbbf10ef77203c7881079a0))


### Documentation

* **README:** update referenced module name ([78f5dac](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/78f5dac76507a993ce4f553ece7b4ceb5c39d67f))

## [2.0.0-test.3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v2.0.0-test.2...v2.0.0-test.3) (2023-01-23)


### Features

* update module name to github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v2 ([aec4af5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/aec4af5bd7ec823dabbbf10ef77203c7881079a0))

## [2.0.0-test.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v2.0.0-test.1...v2.0.0-test.2) (2023-01-23)


### Features

* introduce identification error into EventsResponse ([925334e](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/925334e52b4046b97d6b814f734b81ff2086fee7))
* store RetryAfter in ManyRequestsResponse ([8239e3c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/8239e3cae5e6246440f4ab76fcd605fb78aa50ab))

## [2.0.0-test.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.2.0...v2.0.0-test.1) (2023-01-18)


### ⚠ BREAKING CHANGES

* `StSeenAt` type renamed to `SeenAt`

### Features

* generate new source file with updated swagger ([1d94e69](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1d94e698850a2f753b7f1398cdd667a6ae5aea10))
* Update list of examples in generate.go (new errors) ([a328ad6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a328ad6defbeff6cedc6769b87595fb18f56ba9e))

## [1.2.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.1.0...v1.2.0) (2022-10-24)


### Features

* update schema to support url field for botd result ([5e0ec6c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/5e0ec6c9c65ec79e20dfbb062c6a7471215852cd))


### Documentation

* **README:** add different region to code example ([3986d6d](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/3986d6df1306666bb03812be05c408ed91ecf0d9))
* **README:** add region section ([a2342cd](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a2342cdc0451982ee5c33bd46704d193a263ddd1))

## [1.1.0](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.2...v1.1.0) (2022-09-19)


### Features

* introduce /event/{request_id} endpoint ([74a39b6](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/74a39b609b64ef2f9b7eae76972d7e4532b1867b))

## [1.0.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.1...v1.0.2) (2022-09-01)


### Documentation

* **README:** update template ([0bb3917](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/0bb391711ec3625af7c8ffb2de6bdc525758fbf1))

## [1.0.1](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0...v1.0.1) (2022-09-01)


### Documentation

* **README:** remove WIP label ([5d910ae](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/5d910ae9a0d43e19647d5982eefec536502f616f))

## 1.0.0 (2022-09-01)


### Features

* add "integrationsInfo" query param ([b326815](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/b326815f69b92c3c1d2d691a99c8483753ec6e49))
* create Go SDK ([a5e03b5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a5e03b5b1ad5e58441d88faf992f5f6e08033d55))
* support passing region ([1ba2e94](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1ba2e941ae8fe65abd706f7e5506953b03cde9ab))


### Bug Fixes

* send API key only in headers ([92a4f88](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/92a4f886b6876d878d9c7ca61f6b4e3af34445d6))
* support nil values for time.Time ([459ba4c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/459ba4c8dde6c6e1428fdeb9b0c2975de1a2f1d6))
* use config.json as single source of truth ([519f0d7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/519f0d7b0c7c84fc164c4cf71440a83c87ab6239))


### Documentation

* **README:** fix installation cmd typo ([2017b4c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/2017b4c890cf7d0cb6b9dd1df5a374a8af2e96a4))
* **README:** remove unnecessary import from example ([e6759e7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e6759e71e712cff1508fbfc88a941e68244bbd66))
* **README:** update readme ([ae4e0ea](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/ae4e0ea67c95598f3771cd1e7c89189bab17793e))

## [1.0.0-test.5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0-test.4...v1.0.0-test.5) (2022-08-29)


### Bug Fixes

* send API key only in headers ([92a4f88](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/92a4f886b6876d878d9c7ca61f6b4e3af34445d6))
* support nil values for time.Time ([459ba4c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/459ba4c8dde6c6e1428fdeb9b0c2975de1a2f1d6))


### Documentation

* **README:** fix installation cmd typo ([2017b4c](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/2017b4c890cf7d0cb6b9dd1df5a374a8af2e96a4))
* **README:** remove unnecessary import from example ([e6759e7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/e6759e71e712cff1508fbfc88a941e68244bbd66))

## [1.0.0-test.4](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0-test.3...v1.0.0-test.4) (2022-08-25)


### Documentation

* **README:** update readme ([ae4e0ea](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/ae4e0ea67c95598f3771cd1e7c89189bab17793e))

## [1.0.0-test.3](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0-test.2...v1.0.0-test.3) (2022-08-24)


### Bug Fixes

* use config.json as single source of truth ([519f0d7](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/519f0d7b0c7c84fc164c4cf71440a83c87ab6239))

## [1.0.0-test.2](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/compare/v1.0.0-test.1...v1.0.0-test.2) (2022-08-19)


### Features

* add "integrationsInfo" query param ([b326815](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/b326815f69b92c3c1d2d691a99c8483753ec6e49))

## 1.0.0-test.1 (2022-08-19)


### Features

* create Go SDK ([a5e03b5](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/a5e03b5b1ad5e58441d88faf992f5f6e08033d55))
* support passing region ([1ba2e94](https://github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/commit/1ba2e941ae8fe65abd706f7e5506953b03cde9ab))
