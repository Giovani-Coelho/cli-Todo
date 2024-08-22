/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database/repository"
	"github.com/Giovani-Coelho/Todo-CLI/src/service"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskRepository := repository.NewMemoryRepository()
		taskService := service.NewTaskService(taskRepository)

		tasks, err := taskService.ListTasks()
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
