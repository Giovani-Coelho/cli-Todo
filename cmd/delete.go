/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	taskUseCase "github.com/Giovani-Coelho/Todo-CLI/src/useCase/task"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskUseCase.Task.DeleteTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
