/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		printFormatterTaskList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printFormatterTaskList() {
	var maxLen int
	for _, task := range getTaskList() {
		lenght := len(task)
		if lenght > maxLen {
			maxLen = lenght
		}
	}

	tabFilter := tabwriter.NewWriter(os.Stdout, maxLen+10, 8, 1, ' ', 0)

	for i, record := range getCsvData() {
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
