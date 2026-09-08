---
"fingerprint-pro-server-api-go-sdk": patch
---

Removed `omitempty` from nonpointer bool fields. Affected fields are `Identification.Suspect`, `LabelsInner.Prediction`, `RareDevice.Result`, `SupplementaryId.VisitorFound`, `VpnMethods.MlPrediction`, `Webhook.VisitorFound`, `Webhook.Incognito`, `Webhook.Replayed`, and `Result` on all webhook smart signal models (`WebhookClonedApp`, `WebhookDeveloperTools`, `WebhookEmulator`, `WebhookFrida`, `WebhookIpBlocklist`, `WebhookJailbroken`, `WebhookLocationSpoofing`, `WebhookMitMAttack`, `WebhookPrivacySettings`, `WebhookProxy`, `WebhookRareDevice`, `WebhookRemoteControl`, `WebhookRootApps`, `WebhookTampering`, `WebhookTor`, `WebhookVirtualMachine`, `WebhookVpn`) and also `WebhookTampering.AntiDetectBrowser`
