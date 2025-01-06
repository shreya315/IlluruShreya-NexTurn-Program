package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Change to mattn/go-sqlite3
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./inventory.db") // Use sqlite3 as the driver name
	if err != nil {
		log.Fatal(err)
	}
	createTable := `CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT,
        price REAL NOT NULL,
        stock INTEGER NOT NULL,
        category_id INTEGER
    );`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
