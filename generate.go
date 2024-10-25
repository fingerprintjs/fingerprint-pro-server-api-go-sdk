package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/config"
)

var files = []string{"README.md", "docs", ".swagger-codegen"}
var filesToKeep = []string{"docs/DecryptionKey.md", "docs/SealedResults.md", "docs/Webhook.md"}
var pathPrefix = "sdk"

func main() {
	moveFilesToKeepToTmpDir()
	handlePotentialMajorRelease()
	bumpConfigVersion()
	generateSwagger()
	moveFiles()
	fixFingerPrintApiMdFile()
	fixErrorCodemodel()
	moveFilesToKeepFromTmpDir()
	formatCode()
}

func ensureTmpDir(paths ...string) {
	fullPath := fmt.Sprintf("tmp/%s", strings.Join(paths, "/"))

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err := os.Mkdir(fullPath, 0755)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func moveFilesToKeepToTmpDir() {
	ensureTmpDir()
	ensureTmpDir("docs")

	for _, file := range filesToKeep {
		filePath := fmt.Sprintf("%s", file)
		newFilePath := fmt.Sprintf("tmp/%s", file)

		err := os.Rename(filePath, newFilePath)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func moveFilesToKeepFromTmpDir() {
	for _, file := range filesToKeep {
		filePath := fmt.Sprintf("./tmp/%s", file)
		newFilePath := fmt.Sprintf("%s", file)

		err := os.Rename(filePath, newFilePath)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func getModuleVersion() string {
	cmd := exec.Command("go", "mod", "edit", "-json")

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var module struct {
		Module struct {
			Path string
		}
	}
	if err := json.Unmarshal(output, &module); err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(module.Module.Path, "/")
	version := parts[len(parts)-1]

	// Return version without "v" prefix
	return version[1:]
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

func replaceMajorVersionMentions(newMajor string, oldMajor string) {
	newMajor = fmt.Sprintf("github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v%s", newMajor)
	oldMajor = fmt.Sprintf("github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v%s", oldMajor)

	log.Println("Replacing major version mentions in files", oldMajor, "->", newMajor)

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() || strings.Contains(path, ".git") || strings.Contains(path, "node_modules") {
			log.Printf("Skipping %s", path)
			return nil
		}

		fileContents, err := os.ReadFile(path)

		if err != nil {
			return err
		}

		log.Printf("Processing %s", path)

		newContents := strings.ReplaceAll(string(fileContents), oldMajor, newMajor)

		err = os.WriteFile(path, []byte(newContents), 0644)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func handlePotentialMajorRelease() {
	version := getVersion()
	newMajorVersion := strings.Split(version, ".")[0]

	moduleVersion := getModuleVersion()

	if newMajorVersion != moduleVersion {
		log.Println("Major update detected, bumping version usage in all files")

		replaceMajorVersionMentions(newMajorVersion, moduleVersion)
	}
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
		"res/fingerprint-server-api.yaml",
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

// fixErrorCodemodel fixes a bug in the generated model_error_code.go file.
// The TOOMANYREQUESTS error code has a wrong name, it is generated as 429TOOMANYREQUESTS_ instead of TOOMANYREQUESTS429 ErrorCode.
// This function reads the file, replaces the wrong name with the correct one and saves the changes.
func fixErrorCodemodel() {
	path := "sdk/model_error_code.go"
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	contents = []byte(strings.Replace(string(contents), "429TOOMANYREQUESTS_ ErrorCode", "TOOMANYREQUESTS429 ErrorCode", -1))
	err = os.WriteFile(path, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// fixFingerPrintApiMdFile fixes a bug with generated file in "docs/FingerprintApi.md" which contains invalid title generated by swagger
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
