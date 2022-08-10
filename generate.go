package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// TODO Move Readme.md and "docs" and "api" dir to root
func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cwd)

	cmd := exec.Command("sh", "generate.sh")
	cmd.Dir = cwd

	out, cmdErr := cmd.Output()

	if cmdErr != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
