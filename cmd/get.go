/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	taskUseCase "github.com/Giovani-Coelho/Todo-CLI/src/useCase/task"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := taskUseCase.Task.ListTasks()
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
