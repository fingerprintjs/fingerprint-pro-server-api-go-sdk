package config

import (
	"encoding/json"
	"log"
	"os"
)

func ReadConfig(fileName string) map[string]interface{} {
	configContents, err := os.ReadFile(fileName)

	var config map[string]interface{}

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(configContents, &config); err != nil {
		log.Fatal(err)
	}

	configContents, err = json.MarshalIndent(config, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	return config
}
