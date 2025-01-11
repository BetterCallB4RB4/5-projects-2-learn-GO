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
		// taskFlagOn, _ := cmd.Flags().GetString("task")
		// if taskFlagOn != "" {
		// 	taskId, err := strconv.Atoi(args[0])
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		os.Exit(1)
		// 	}
		// 	checkTaskByID(taskId)
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		os.Exit(1)
		// 	}
		// } else {
		// }

		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		checkTaskByID(taskId)
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

	checkCmd.PersistentFlags().String("task", "t", "select a task by the name")
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
