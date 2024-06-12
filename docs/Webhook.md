# Webhook

## **IsValidWebhookSignature**

> bool IsValidWebhookSignature(header string, data []byte, secret string)

Verifies the HMAC signature extracted from the "fpjs-event-signature" header of the incoming request. This is a part of the webhook signing process, which is available only for enterprise customers.
If you wish to enable it, please [contact our support](https://fingerprint.com/support).

### Required Parameters

| Name       | Type       | Description                                                | Notes |
|------------|------------|------------------------------------------------------------|-------|
| **header** | **string** | Value of the "fpjs-event-signature" header.                |       |
| **data**   | **[]byte** | Body of the request from which above header was extracted. |       | 
| **secret** | **secret** | Your generated secret used to sign the request.            |       | 
