/*
Copyright © 2024 Giovani Coelho
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Giovani-Coelho/Todo-CLI/src/infra/database"
	taskUseCase "github.com/Giovani-Coelho/Todo-CLI/src/useCase/task"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "CLI app for managing tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Esse Run só será executado se nenhum comando específico for passado
		fmt.Println("Use a valid command.")
	},
}

func Execute() {
	database.InitDB()
	if database.SQL == nil {
		log.Fatal("Failed to initialize database connection")
	}
	defer database.SQL.Close()

	taskUseCase.InitDependencies()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
