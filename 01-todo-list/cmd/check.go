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
	Short: "Mark a task as completed by ID or name",
	Long: `The check command allows you to mark a task as completed either by providing the task ID 
or by using the task name with the --task flag.

If you provide a task ID as a positional argument, the task corresponding to that ID will be marked as done.
Alternatively, you can use the --task flag to specify the task by its name, and it will be marked as done.

Examples:
  todo-list check 1
  This command will mark the task with ID 1 as completed.

  todo-list check --task "Buy groceries"
  This command will mark the task with the name "Buy groceries" as completed.

This is helpful for tracking and managing your tasks in the to-do list.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName, _ := cmd.Flags().GetString("task")
		if taskName != "" {
			// Call the function to check a task by name
			checkTaskByString(taskName)

			fmt.Println("")
			printFormatterTaskList()

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
		fmt.Println("")
		printFormatterTaskList()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// here you init flags, look at the cobra cli doc to check how this work
	checkCmd.Flags().StringP("task", "t", "", "check a task by his name")
}

func checkTaskByID(taskId int) {
	records := getCsvData()
	if len(records) == 0 || taskId > len(records) {
		return
	}
	records[taskId-1].done = true
	writeCsv(records)
}

func checkTaskByString(taskString string) {
	records := getCsvData()
	if len(records) == 0 {
		return
	}
	for i, record := range records {
		if record.task == taskString {
			records[i].done = true
		}
	}
	writeCsv(records)
}
