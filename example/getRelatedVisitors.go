package main

import (
	"context"
	"encoding/json"
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

	response, httpRes, err := client.FingerprintApi.GetRelatedVisitors(auth, visitorId)
	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		var tooManyRequestsError *sdk.TooManyRequestsError

		if errors.As(err, &tooManyRequestsError) {
			log.Printf("Too many requests, retry after %d seconds", tooManyRequestsError.RetryAfter())
		} else {
			log.Fatalf("Error: %s, %s", err.Code(), err.Error())
		}
	}

	// Print full response as JSON
	responseJsonData, jsonErr := json.MarshalIndent(response, "", "  ")
	if jsonErr != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(responseJsonData))
}
