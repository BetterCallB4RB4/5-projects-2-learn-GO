/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName, _ := cmd.Flags().GetString("task")
		if taskName != "" {
			// Call the function to check a task by name
			checkTaskByString(taskName)
			return
		}

		// If no --task flag, expect an integer argument
		if len(args) == 0 {
			fmt.Println("Error: You must provide a task ID or use the --task flag.")
			os.Exit(1)
		}

		// Parse the task ID from the argument
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: Invalid task ID '%s'. Please provide a valid integer.\n", args[0])
			os.Exit(1)
		}

		// Call the function to check a task by ID
		checkTaskByID(taskID)
		fmt.Printf("Task with ID %d marked as done.\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// checkCmd.PersistentFlags().String("task", "", "select a task by the name")
	checkCmd.Flags().StringP("task", "t", "", "check a task by his name")
}

func checkTaskByID(taskId int) {
	records := getCsvData()
	records[taskId-1].done = true
	writeCsv(records)
}

func checkTaskByString(taskString string) {
	records := getCsvData()
	for i, record := range records {
		if record.task == taskString {
			records[i].done = true
		}
	}
	writeCsv(records)
}
