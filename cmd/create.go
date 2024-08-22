package cmd

import (
	"fmt"

	"github.com/Giovani-Coelho/Todo-CLI/src/di"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		res := di.TaskController.NewTask(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
