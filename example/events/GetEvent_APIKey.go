package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk"
	"log"
	"os"
)

func main() {
	cfg := sdk.NewConfiguration()
	client := sdk.NewAPIClient(cfg)

	// You can also use sdk.RegionUS or sdk.RegionAsia. Default one is sdk.RegionUS
	//cfg.ChangeRegion(sdk.RegionEU)

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

	if response.Products.Botd != nil {
		fmt.Printf("Got response with Botd: %s", response.Products.Botd)
	}

	if response.Products.Identification != nil {
		stringResponse, _ := json.Marshal(response.Products.Identification)
		fmt.Printf("Got response with Identification: %s", string(stringResponse))
	}
}
