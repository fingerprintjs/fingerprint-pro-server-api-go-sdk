package test

import (
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk/webhook"
	"github.com/stretchr/testify/assert"
	"testing"
)

const secret = "secret"
const data = "data"

func TestWebhookModel(t *testing.T) {
	t.Run("Check if webhook response can be unmarshalled", func(t *testing.T) {
		var mockResponse sdk.Webhook
		err := readFromFileAndUnmarshalWithError("../test/mocks/webhook.json", &mockResponse)
		assert.Nil(t, err)
	})
}

func TestIsValidWebhookSignature(t *testing.T) {
	const validHeader = "v1=1b2c16b75bd2a870c114153ccda5bcfca63314bc722fa160d690de133ccbb9db"

	t.Run("With valid signature", func(t *testing.T) {
		isEqual := webhook.IsValidWebhookSignature(validHeader, []byte(data), secret)

		assert.Equal(t, true, isEqual)
	})

	t.Run("With invalid header", func(t *testing.T) {
		isEqual := webhook.IsValidWebhookSignature("v2=invalid", []byte(data), secret)

		assert.Equal(t, false, isEqual)
	})

	t.Run("With header without version", func(t *testing.T) {
		isEqual := webhook.IsValidWebhookSignature("invalid", []byte(data), secret)

		assert.Equal(t, false, isEqual)
	})

	t.Run("With empty header", func(t *testing.T) {
		isEqual := webhook.IsValidWebhookSignature("", []byte(data), secret)

		assert.Equal(t, false, isEqual)
	})

	t.Run("With empty secret", func(t *testing.T) {
		isEqual := webhook.IsValidWebhookSignature(validHeader, []byte(data), "")

		assert.Equal(t, false, isEqual)
	})

	t.Run("With empty data", func(t *testing.T) {
		isEqual := webhook.IsValidWebhookSignature(validHeader, []byte(""), secret)

		assert.Equal(t, false, isEqual)
	})
}
