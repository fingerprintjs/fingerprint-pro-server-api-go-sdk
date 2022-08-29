package main

import (
	"encoding/json"
	"fmt"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/config"
	"log"
	"os"
	"os/exec"
	"strings"
)

var files = []string{"README.md", "docs", ".swagger-codegen"}
var pathPrefix = "sdk"

func main() {
	bumpConfigVersion()
	generateSwagger()
	moveFiles()
	getExamples()
	fixFingerPrintApiMdFile()
	formatCode()
}

func getVersion() string {
	var version string

	envVersion := os.Getenv("VERSION")

	if envVersion != "" {
		version = envVersion
	} else {
		configFile := config.ReadConfig("./config.json")
		version = configFile.PackageVersion
	}

	return version
}

func bumpConfigVersion() {
	version := getVersion()

	configFile := config.ReadConfig("./config.json")

	if configFile.PackageVersion == version {
		log.Println("Version is up to date")
		return
	}

	configFile.PackageVersion = version

	configContents, err := json.MarshalIndent(configFile, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err = os.WriteFile("./config.json", configContents, 0644); err != nil {
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
	cmd := exec.Command(
		"java",
		"-jar",
		"./bin/swagger-codegen-cli.jar",
		"generate",
		"-t",
		"./template",
		"-l",
		"go",
		"-i",
		"https://fingerprintjs.github.io/fingerprint-pro-server-api-openapi/schemes/fingerprint-server-api.yaml",
		"-o",
		"./sdk",
		"-c",
		"config.json")

	out, cmdErr := cmd.Output()

	if cmdErr != nil {
		log.Fatal(cmdErr)
	}

	fmt.Println(string(out))
}

func getExamples() {
	list := []string{"visits_limit_1.json", "visits_limit_500.json", "webhook.json"}

	for _, file := range list {
		cmd := exec.Command("curl", "-o", fmt.Sprintf("./test/mocks/%s", file), fmt.Sprintf("https://fingerprintjs.github.io/fingerprint-pro-server-api-openapi/examples/%s", file))
		_, err := cmd.Output()

		if err != nil {
			log.Fatal(err)
		}
	}

}

/**
 * Fixes a bug with generated file in "docs/FingerprintApi.md" which contains invalid title generated by swagger
 */
func fixFingerPrintApiMdFile() {
	token := "{{classname}}"
	target := "FingerprintApi"
	targetsToRemove := []string{"**optional** | ***FingerprintApiGetVisitsOpts** | optional parameters | nil if no parameters"}
	filePath := "docs/FingerprintApi.md"

	fileContents, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fileContents = []byte(strings.Replace(string(fileContents), token, target, -1))
	fileContentsArray := strings.Split(string(fileContents), "\n")
	var fileContentsArrayResult []string

	// Fixes markdown table for optional parameters, by default swagger-codegen applies new line there which breaks the table.
	for _, line := range fileContentsArray {
		for _, targetToRemove := range targetsToRemove {
			if line != targetToRemove {
				fileContentsArrayResult = append(fileContentsArrayResult, strings.Replace(line, targetToRemove, "", -1))
			}
		}
	}

	err = os.WriteFile(filePath, []byte(strings.Join(fileContentsArrayResult, "\n")), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func formatCode() {
	cmd := exec.Command("go", "fmt")
	cmd.Dir = "./sdk"

	_, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
}
