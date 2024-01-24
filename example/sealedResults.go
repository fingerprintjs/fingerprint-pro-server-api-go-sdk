package main

import (
	"encoding/base64"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/sdk/sealed"
	"os"
)

func base64Decode(input string) []byte {
	output, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	sealedResult := base64Decode(os.Getenv("BASE64_SEALED_RESULT"))
	key := base64Decode(os.Getenv("BASE64_KEY"))

	keys := []sealed.DecryptionKey{
		{
			Key:       key,
			Algorithm: sealed.AlgorithmAES256GCM,
		},
	}
	unsealedResponse, err := sealed.UnsealEventsResponse(sealedResult, keys)
	if err != nil {
		fmt.Println("Unseal error:", err)
		return
	}

	PrintEventResponse(*unsealedResponse)
}
