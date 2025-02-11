package main

import (
	"context"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	opts := sdk.FingerprintApiSearchEventsOpts{
		//Suspect: false,
	}

	response, httpRes, err := client.FingerprintApi.SearchEvents(auth, 10, &opts)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s, %s", err.Code(), err.Error())
	}

	if response.Events != nil {
		fmt.Printf("Got response with Events: %v \n", response.Events)
	}
}
