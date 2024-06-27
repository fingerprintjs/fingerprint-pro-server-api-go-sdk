package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v5/sdk/sealed"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func base64Decode(input string) []byte {
	output, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[sealedResults error] Error while loading env", err)
		return
	}

	sealedResult := base64Decode(os.Getenv("BASE64_SEALED_RESULT"))
	key := base64Decode(os.Getenv("BASE64_SEALED_RESULT_KEY"))

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

	var response = *unsealedResponse
	if response.Products.Botd != nil {
		fmt.Printf("Got response with Botd: %v \n", response.Products.Botd)
	}

	if response.Products.Identification != nil {
		stringResponse, _ := json.Marshal(response.Products.Identification)
		fmt.Printf("Got response with Identification: %s \n", string(stringResponse))

	}

	if response.Products.Emulator != nil {
		fmt.Printf("Got response with Emulator: %v \n", response.Products.Emulator.Data)
	}

	if response.Products.IpInfo != nil {
		fmt.Printf("Got response with IpInfo: %v \n", response.Products.IpInfo.Data)
	}

	if response.Products.Incognito != nil {
		fmt.Printf("Got response with Incognito: %v \n", response.Products.Incognito.Data)
	}

	if response.Products.RootApps != nil {
		fmt.Printf("Got response with RootApps: %v \n", response.Products.RootApps.Data)
	}

	if response.Products.ClonedApp != nil {
		fmt.Printf("Got response with ClonedApp: %v \n", response.Products.ClonedApp.Data)
	}

	if response.Products.FactoryReset != nil {
		fmt.Printf("Got response with FactoryReset: %v \n", response.Products.FactoryReset.Data)
	}

	if response.Products.Jailbroken != nil {
		fmt.Printf("Got response with Jailbroken: %v \n", response.Products.Jailbroken.Data)
	}

	if response.Products.Frida != nil {
		fmt.Printf("Got response with Frida: %v \n", response.Products.Frida.Data)
	}

	if response.Products.IpBlocklist != nil {
		fmt.Printf("Got response with IpBlocklist: %v \n", response.Products.IpBlocklist.Data)
	}

	if response.Products.Tor != nil {
		fmt.Printf("Got response with Tor: %v \n", response.Products.Tor.Data)
	}

	if response.Products.PrivacySettings != nil {
		fmt.Printf("Got response with PrivacySettings: %v \n", response.Products.PrivacySettings.Data)
	}

	if response.Products.VirtualMachine != nil {
		fmt.Printf("Got response with VirtualMachine: %v \n", response.Products.VirtualMachine.Data)
	}

	if response.Products.Vpn != nil {
		fmt.Printf("Got response with Vpn: %v \n", response.Products.Vpn.Data)
	}

	if response.Products.Proxy != nil {
		fmt.Printf("Got response with Proxy: %v \n", response.Products.Proxy.Data)
	}

	if response.Products.Tampering != nil {
		fmt.Printf("Got response with Tampering: %v \n", response.Products.Tampering.Data)
	}

	if response.Products.RawDeviceAttributes != nil {
		fmt.Printf("Got response with RawDeviceAttributes: %v \n", response.Products.RawDeviceAttributes.Data)
	}
}
