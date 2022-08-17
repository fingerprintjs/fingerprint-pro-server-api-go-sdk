package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	version := os.Getenv("VERSION")

	if version == "" {
		log.Fatal("VERSION environment variable not set")
	}

	fileName := "./config.json"
	configContents, err := os.ReadFile(fileName)

	var config map[string]interface{}

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(configContents, &config); err != nil {
		log.Fatal(err)
	}

	config["version"] = version

	configContents, err = json.MarshalIndent(config, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fileName, configContents, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
