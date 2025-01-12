package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in your to-do list with details",
	Long: `The list command displays all tasks stored in your to-do list along with their details,
including the task ID, description, creation time, and completion status.

Tasks that are marked as completed will be shown with a strike-through style, making it easy to visually distinguish between completed and pending tasks.

Examples:
  todo-list list
  This command will list all tasks with their details (ID, Task, Created, Done). Completed tasks will be marked with a strike-through.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		printFormatterTaskList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func printFormatterTaskList() {
	var maxLen int
	records := getCsvData()

	for _, record := range records {
		lenght := len(record.task)
		if lenght > maxLen {
			maxLen = lenght
		}
	}

	tabFilter := tabwriter.NewWriter(os.Stdout, maxLen+15, 10, 4, ' ', 0)

	for i, record := range records {
		// print the header
		if i == 0 {
			fmt.Fprintln(tabFilter, "ID"+"\t"+"TASK"+"\t"+"CREATED"+"\t"+"DONE")
		}

		// compute the printable format for:
		// ID
		stringID := strconv.Itoa(record.id)

		// time
		timeDiff := timediff.TimeDiff(record.age)

		// done field
		stringDone := strconv.FormatBool(record.done)

		if record.done {
			// this allow for strikethrow with ascee code and some random space for formatting issue
			fmt.Fprintln(tabFilter, "\033[9m"+stringID+"\t"+"    "+record.task+"\t"+"    "+timeDiff+"\t"+"    "+stringDone+"\033[0m")
		} else {
			fmt.Fprintln(tabFilter, stringID+"\t"+record.task+"\t"+timeDiff+"\t"+stringDone)
		}
	}
	tabFilter.Flush()
}
