package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	PackageVersion string `json:"packageVersion"`
	GitRepoId      string `json:"gitRepoId"`
	GitUserId      string `json:"gitUserId"`
	PackageName    string `json:"packageName"`
	PackageUrl     string `json:"packageUrl"`
}

func ReadConfig(fileName string) Config {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var conf Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
