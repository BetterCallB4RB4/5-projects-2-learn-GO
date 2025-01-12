package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove completed tasks from the to-do list",
	Long: `The remove command allows you to remove all tasks that have been marked as completed 
from your to-do list. This command is useful to clean up completed tasks after you've finished them.

Examples:
  todo-list remove
  This command will remove all tasks marked as completed from your list.

It works by reading the tasks stored in the CSV file, filtering out the tasks that are marked as 'done', 
and then saving the updated list back to the CSV file.`,
	Run: func(cmd *cobra.Command, args []string) {
		remeveCheckedTask()
		fmt.Println("")
		printFormatterTaskList()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func remeveCheckedTask() {
	var uncompleteRecord []Record
	for _, record := range getCsvData() {
		if !record.done {
			uncompleteRecord = append(uncompleteRecord, record)
		}
	}
	for i := range uncompleteRecord {
		uncompleteRecord[i].id = i + 1
	}
	writeCsv(uncompleteRecord)
}
