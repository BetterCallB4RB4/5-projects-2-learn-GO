package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile string = "./data/data.db"

func testDatabaseConnection() {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		logger.Error("Cannot establish connection to Database")
		return
	}
	defer db.Close()

	// Create a table (if it doesn't exist)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	if err != nil {
		logger.Error("Cannot create a table inside the DB")
	}

	insertDummyData(db, "Mario", 30)
	insertDummyData(db, "Alessio", 40)
	insertDummyData(db, "Zebe", 33)

	printData(db)
}

func insertDummyData(db *sql.DB, name string, age int) {
	_, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		fmt.Printf("Failed to insert data: %v", err)
	} else {
		fmt.Printf("Inserted user: %s, age: %d", name, age)
	}
}

func printData(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		logger.Error("Cannot print data from the DB")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			logger.Error("Cannot read row number: %d in the database", id)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}
