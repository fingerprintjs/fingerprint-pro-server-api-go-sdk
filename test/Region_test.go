package test

import (
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v2/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsesCorrectEndpointForRegion(t *testing.T) {
	regionsMap := make(map[sdk.Region]string)
	regionsMap[sdk.RegionUS] = "https://api.fpjs.io"
	regionsMap[sdk.RegionAsia] = "https://ap.api.fpjs.io"
	regionsMap[sdk.RegionEU] = "https://eu.api.fpjs.io"

	for region, endpoint := range regionsMap {
		cfg := sdk.NewConfiguration()
		cfg.ChangeRegion(region)

		assert.Equal(t, cfg.GetBasePath(), endpoint)
	}
}
