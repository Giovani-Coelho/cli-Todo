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

// rootCmd represents the base command when called without any subcommands
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Todo-CLI.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
