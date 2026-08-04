---
"fingerprint-pro-server-api-go-sdk": minor
---

Add the `RareDevice` Smart Signal. The `Products` response now exposes a `RareDevice` field (`ProductRareDevice`) whose `data` is a `bool` that is true when the device is considered rare based on its hardware and software attributes. This signal is currently in beta and only available to select customers.

Add `RareDevice` and `RareDevicePercentileBucket` filters to `FingerprintApiSearchEventsOpts`
