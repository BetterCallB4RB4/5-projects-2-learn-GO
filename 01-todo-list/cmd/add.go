package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// global variable
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adding task")
		addTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(task string) {
	var record [4]string

	record[0] = strconv.Itoa(getLastIndex())
	record[1] = task
	record[2] = time.Now().String()
	record[3] = strconv.FormatBool(false)

	addCsvRecord(record[:])
}

func getLastIndex() int {
	return 0
}
