/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

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
	tabFilter := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	reader := createCsvReader(fileName)

	isHeader := false
	for {
		if !isHeader {
			header, err := reader.Read()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			isHeader = true
			fmt.Fprintln(tabFilter, header[0]+"\t"+header[1]+"\t"+header[2]+"\t"+header[3]+"\t")

		}

		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// dummy formatting, can be
		id := record[0]
		taskName := record[1]
		isDone := record[3]

		timeDate, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		timeDiff := timediff.TimeDiff(timeDate)

		fmt.Fprintln(tabFilter, id+"\t"+taskName+"\t"+timeDiff+"\t"+isDone+"\t")

		tabFilter.Flush()
	}
}
