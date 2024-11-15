package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
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
	requestId := os.Getenv("REQUEST_ID")

	response, httpRes, err := client.FingerprintApi.GetEvent(auth, requestId)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		var apiError sdk.ApiError

		if errors.As(err, &apiError) {
			log.Fatalf("Error: %s, %s", apiError.Code(), apiError.Error())
		}

		log.Fatal(err)
	}

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
