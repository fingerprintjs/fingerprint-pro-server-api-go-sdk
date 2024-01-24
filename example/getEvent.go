package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk"
	"github.com/joho/godotenv"
)

func main() {
	cfg := sdk.NewConfiguration()
	client := sdk.NewAPIClient(cfg)

	// Load environment variables
	godotenv.Load()

	// Default region is sdk.RegionUS
	if os.Getenv("REGION") == "eu" {
		cfg.ChangeRegion(sdk.RegionEU)
	}
	if os.Getenv("REGION") == "ap" {
		cfg.ChangeRegion(sdk.RegionAsia)
	}

	// Configure authorization, in our case with API Key
	auth := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
		Key: os.Getenv("FINGERPRINT_API_KEY"),
	})

	// Usually this data will come from your frontend app
	requestId := os.Getenv("REQUEST_ID")

	response, httpRes, err := client.FingerprintApi.GetEvent(auth, requestId)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatal(err)
	}

	PrintEventResponse(response)
}
