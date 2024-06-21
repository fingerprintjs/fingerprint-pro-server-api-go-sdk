package main

import (
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v6/sdk/webhook"
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
