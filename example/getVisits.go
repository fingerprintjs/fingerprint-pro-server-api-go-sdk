package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v6/sdk"
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
	opts := sdk.FingerprintApiGetVisitsOpts{
		RequestId: os.Getenv("REQUEST_ID"),
	}

	response, httpRes, err := client.FingerprintApi.GetVisits(auth, visitorId, &opts)
	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		var tooManyRequestsError *sdk.TooManyRequestsError

		if errors.As(err, &tooManyRequestsError) {
			log.Fatalf("Too many requests, retry after %d seconds", tooManyRequestsError.RetryAfter())
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("Got response with visitorId: %s", response.VisitorId)

	// Print full response as JSON
	responseJsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(responseJsonData))
}
