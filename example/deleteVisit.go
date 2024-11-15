package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"log"
	"os"

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
	visitorId := os.Getenv("VISITOR_ID")

	// Delete visitor data. If you are interested in using this API, please contact our support team (https://fingerprint.com/support/) to activate it for you
	httpRes, err := client.FingerprintApi.DeleteVisitorData(auth, visitorId)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		var apiError sdk.ApiError

		if errors.As(err, &apiError) {
			log.Fatalf("Error: %s, %s", apiError.Code(), apiError.Error())
		}

		log.Fatal(err)
	}

}
