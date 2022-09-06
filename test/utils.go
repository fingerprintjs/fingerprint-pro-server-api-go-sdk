package test

import (
	"encoding/json"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk"
	"log"
	"os"
)

func GetMockResponse(path string) sdk.Response {
	var mockResponse sdk.Response

	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &mockResponse)

	if err != nil {
		log.Fatal(err)
	}

	return mockResponse
}

func GetMockEventResponse(path string) sdk.EventResponse {
	var mockResponse sdk.EventResponse

	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &mockResponse)

	if err != nil {
		log.Fatal(err)
	}

	return mockResponse
}
