package version

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func getCurrentVersion() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		return "", fmt.Errorf("version not found in environment")
	}

	return version, nil
}

func Get(cmd *cobra.Command) {
	currentVersion, err := getCurrentVersion()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("version: %s \n", currentVersion)
}
