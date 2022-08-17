package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var files = []string{"README.md", "docs", ".swagger-codegen"}
var pathPrefix = "sdk"

func main() {
	bumpConfigVersion()
	generateSwagger()
	moveFiles()
}

func getVersion() string {
	var version string

	envVersion := os.Getenv("VERSION")

	if envVersion != "" {
		version = envVersion
	} else {
		version = "1.0.0"
	}

	return version
}

func bumpConfigVersion() {
	version := getVersion()
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

func removeFileOrDirIfExists(path string) {
	if stat, err := os.Stat(path); err == nil {
		var err error

		if stat.IsDir() {
			err = os.RemoveAll(path)
		} else {
			err = os.Remove(path)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}

func cleanupOldFiles() {
	for _, filePath := range files {

		removeFileOrDirIfExists(filePath)

	}
}

func moveFiles() {
	cleanupOldFiles()

	for _, file := range files {

		filePath := fmt.Sprintf("%s/%s", pathPrefix, file)
		newFilePath := fmt.Sprintf("./%s", file)

		err := os.Rename(filePath, newFilePath)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func generateSwagger() {
	cmd := exec.Command("sh", "generate.sh")

	out, cmdErr := cmd.Output()

	if cmdErr != nil {
		log.Fatal(cmdErr)
	}

	fmt.Println(string(out))
}
