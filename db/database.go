package db

import (
	"database/sql"
	"log"

	// Import and initialize the SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

// DB, holds the database connection
var DB *sql.DB

// InitDB, initializes the database and creates the table
func InitDB() {
	var err error
	// Connect to the SQLite database (creates the file if it doesn't exist)
	DB, err = sql.Open("sqlite3", "./accounts.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Create the table
	createTable := `
	CREATE TABLE IF NOT EXISTS accounts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		platform TEXT NOT NULL,
		url TEXT NOT NULL,
		identity TEXT NOT NULL,
		passphrase TEXT NOT NULL,
		notes TEXT NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Failed to create the table:", err)
	}

	log.Println("Database initialized.")
}
