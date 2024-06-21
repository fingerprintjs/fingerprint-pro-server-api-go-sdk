package test

import (
	"encoding/json"
	"log"
	"os"
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

func GetMockResponse[T any](path string) T {
	var mockResponse T
	readFromFileAndUnmarshal(path, &mockResponse)
	return mockResponse
}
