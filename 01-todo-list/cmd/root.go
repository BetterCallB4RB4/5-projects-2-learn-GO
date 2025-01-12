package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	fileName string = "tasks.csv"
	rootCmd         = &cobra.Command{
		Use:   "todo-list",
		Short: "A CLI tool for managing your to-do list tasks",
		Long: `todo-list is a command-line application for managing your to-do list tasks.
You can use it to create, list, and manage tasks, mark them as completed, and delete tasks. 
Tasks are stored in a CSV file.

Examples:
  todo-list add "Buy groceries"
  todo-list list
  todo-list check 2 
  todo-list check -t "Buy groceries"

This application uses Cobra to handle subcommands and flags to customize the behavior.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCsvFile(fileName)
}
