package main

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk"
	"log"
	"os"
)

func main() {
	cfg := sdk.NewConfiguration()
	client := sdk.NewAPIClient(cfg)

	// You can also use sdk.RegionUS or sdk.RegionAsia. Default one is sdk.RegionUS
	cfg.ChangeRegion(sdk.RegionEU)

	// Configure authorization, in our case with API Key
	auth := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
		Key: os.Getenv("FINGERPRINT_API_KEY"),
	})

	// Usually this data will come from your frontend app
	visitorId := os.Getenv("VISITOR_ID")
	opts := sdk.FingerprintApiGetVisitsOpts{
		RequestId: optional.NewString(os.Getenv("REQUEST_ID")),
	}

	response, httpRes, err := client.FingerprintApi.GetVisits(auth, visitorId, &opts)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got response with visitorId: %s", response.VisitorId)
}
