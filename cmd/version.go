package cmd

import (
	"github.com/Giovani-Coelho/Todo-CLI/src/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Long:  `This software is under development`,
	Run: func(cmd *cobra.Command, args []string) {
		version.Get(cmd)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
