package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// global variable
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the to-do list",
	Long: `The add command allows you to add a new task to the to-do list.
To use this command, simply provide the task's name as a positional argument.

The task will be added with the current timestamp, and its 'done' status will be set to false.
Each task will be assigned a unique ID, which is incremented based on the current number of tasks in the list.

Examples:
  todo-list add "Buy groceries"
  This command will add a task named "Buy groceries" to the to-do list.

  todo-list add "Finish project"
  This command will add a task named "Finish project" to the to-do list.`,
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args[0])
		fmt.Println("")
		printFormatterTaskList()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(taskName string) {
	records := getCsvData()
	record := Record{
		id:   len(records) + 1,
		task: taskName,
		age:  time.Now(),
		done: false,
	}
	records = append(records, record)
	writeCsv(records)
}
