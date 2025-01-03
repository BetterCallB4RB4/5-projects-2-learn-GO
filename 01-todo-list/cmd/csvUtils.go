package cmd

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func createCsvReader(filename string) *csv.Reader {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// start reading the file and
	reader := csv.NewReader(bytes.NewReader(data))
	return reader
}

func getLastID(filename string) (lastID int) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// start reading the file and
	reader := csv.NewReader(bytes.NewReader(data))

	var lastRecord []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		lastRecord = record
	}

	value, err := strconv.Atoi(lastRecord[0])
	if err != nil {
		return 0
	}
	return value
}

func initCsvFile(filename string) {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		// init new csv file
		header := []string{"ID", "Task", "Created", "Done"}
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0o644)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		writer := csv.NewWriter(file)
		err = writer.Write(header)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		writer.Flush()
	}
}

func addCsvRecord(filename string, record []string) {
	// create file in appens mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	// create a csv writer and write the input record
	writer := csv.NewWriter(file)
	err = writer.Write(record)
	writer.Flush()
}
