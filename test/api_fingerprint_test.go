package test

import (
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v6/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApiFingerprint(t *testing.T) {
	t.Run("Create with empty config", func(t *testing.T) {
		client := sdk.NewAPIClient(nil)

		assert.NotNil(t, client)
	})
}
