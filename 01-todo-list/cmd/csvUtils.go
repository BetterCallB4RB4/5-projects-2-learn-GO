package cmd

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

// ##### READER #####

func readCsvFile(filename string) ([]byte, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func csvParser(data []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader, nil
}

func processCsv(reader *csv.Reader) {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error caused reading csv line", err)
			break
		}
		fmt.Println(record)
	}
}

// ##### WRITER #####

func createCsvWriter(filename string) (*csv.Writer, *os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	return writer, file, nil
}

func createCsvRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Println("Error on writing record", err)
	}
	writer.Flush()
}

func addCsvRecord(record []string) error {
	filename := "task.csv"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	defer file.Close()

	err = writer.Write(record)
	return err
}
