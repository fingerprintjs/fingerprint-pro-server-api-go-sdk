package main

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk"
	"log"
)

func main() {
	cfg := sdk.NewConfiguration()
	client := sdk.NewAPIClient(cfg)

	// You can also use sdk.RegionEU or sdk.RegionAsia. Default one is RegionUS
	cfg.ChangeRegion(sdk.RegionUS)

	// Configure authorization, in our case with API Key
	auth := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
		Key: "F6gQ8H8vQLc7mVsVKaFx",
	})

	// Usually this data will come from your frontend app
	visitorId := "un01pJ7BKgQgqO6omZCw"
	opts := sdk.FingerprintApiGetVisitsOpts{
		RequestId: optional.NewString("1660296228280.hBBze5"),
	}

	response, _, err := client.FingerprintApi.GetVisits(auth, visitorId, &opts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got response with visitorId: %s", response.VisitorId)
}
