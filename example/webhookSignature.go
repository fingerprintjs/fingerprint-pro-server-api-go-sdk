package main

import (
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk/webhook"
)

func main() {
	// Your webhook signing secret.
	secret := "secret"

	// Request data. In real life scenerio this will be the body of incoming request
	data := []byte("data")

	// Value of the "fpjs-event-signature" header.
	header := "v1=1b2c16b75bd2a870c114153ccda5bcfca63314bc722fa160d690de133ccbb9db"

	isValid := webhook.CheckHeader(header, data, secret)

	if !isValid {
		panic("Invalid signature")
	}
}
