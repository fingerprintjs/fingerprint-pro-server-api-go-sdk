package test

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk"
)

func readFromFileAndUnmarshal(path string, i interface{}) {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &i)

	if err != nil {
		log.Fatal(err)
	}
}

func GetMockResponse(path string) sdk.Response {
	var mockResponse sdk.Response
	readFromFileAndUnmarshal(path, &mockResponse)
	return mockResponse
}

func GetMockEventResponse(path string) sdk.EventResponse {
	var mockResponse sdk.EventResponse
	readFromFileAndUnmarshal(path, &mockResponse)
	return mockResponse
}

func GetEvent403ErrorMockResponse(path string) sdk.ErrorCommon403Response {
	var mockResponse sdk.ErrorCommon403Response
	readFromFileAndUnmarshal(path, &mockResponse)
	return mockResponse
}
