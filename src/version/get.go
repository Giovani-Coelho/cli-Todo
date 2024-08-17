package version

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func getCurrentVersion() (string, error) {
	if Version == "" {
		return "", fmt.Errorf("version not found in environment")
	}

	return Version, nil
}

func Get(cmd *cobra.Command) {
	currentVersion, err := getCurrentVersion()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("version: %s \n", currentVersion)
}
