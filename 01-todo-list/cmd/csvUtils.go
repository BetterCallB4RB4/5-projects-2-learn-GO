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

type Record struct {
	id   int
	task string
	age  string
	done bool
}

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
	if err != nil {
		return
	}
	writer.Flush()
}

func morphRecordToString(records []Record) [][]string {
	var stringRecord [][]string
	for _, record := range records {
		record := []string{
			strconv.Itoa(record.id),
			record.task,
			record.age,
			strconv.FormatBool(record.done),
		}
		stringRecord = append(stringRecord, record)
	}
	return stringRecord
}

func writeCsv(records []Record) {
	// create a new empty file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	// prepare the writer
	writer := csv.NewWriter(file)

	// init the file by writing the header
	header := []string{"ID", "Task", "Created", "Done"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	stringRecords := morphRecordToString(records)
	writer.WriteAll(stringRecords)

	writer.Flush()
}

func getCsvData() []Record {
	reader := createCsvReader(fileName)
	records, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	var data []Record
	for i, record := range records {
		if i == 0 {
			continue
		}
		intId, err := strconv.Atoi(record[0])
		if err != nil {
			return nil
		}
		boolDone, err := strconv.ParseBool(record[3])
		if err != nil {
			return nil
		}
		data = append(data, Record{
			id:   intId,
			task: record[1],
			age:  record[1],
			done: boolDone,
		})
	}
	return data
}

func getIdList() (ids []int) {
	records := getCsvData()
	var idList []int
	for _, record := range records {
		idList = append(idList, record.id)
	}
	return idList
}

func getTaskList() (tasks []string) {
	records := getCsvData()
	var taskList []string
	for _, record := range records {
		taskList = append(taskList, record.task)
	}
	return taskList
}

func getAgeList() (ages []string) {
	records := getCsvData()
	var ageList []string
	for _, record := range records {
		ageList = append(ageList, record.age)
	}
	return ageList
}

func getDoneList() (done []bool) {
	records := getCsvData()
	var doneList []bool
	for _, record := range records {
		doneList = append(doneList, record.done)
	}
	return doneList
}
