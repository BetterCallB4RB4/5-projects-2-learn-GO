package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// using a global instance
const dbFile string = "./data/data.db"

var db *sql.DB

func initDataBaseConnection() {
	// initialize database with a global reference
	var err error
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		logger.Error("Cannot establish connection to Database")
		return
	}
	// defer db.Close()

	// Create a table (if it doesn't exist)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS operation (operation TEXT, number1 INTEGER, number2 INTEGER)")
	if err != nil {
		logger.Error("Cannot create a table inside the DB")
	}
}

func insertOperationData(numbers Operation, kind string) {
	_, err := db.Exec("INSERT INTO operation (operation, number1, number2) VALUES (?, ?, ?)", kind, numbers.Number1, numbers.Number2)
	if err != nil {
		logger.Error("Failed to insert data: %v", err)
	}
}
