package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("sqlite3", "taskmanager.db")
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN NOT NULL CHECK (completed IN (0,1))
	);
	`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to create tasks table", err)

	}

	log.Println("Database connected and tasks table ready")

}
