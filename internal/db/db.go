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

	tasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN NOT NULL CHECK (completed IN (0,1))
	);
	`
	_, err = DB.Exec(tasksTable)
	if err != nil {
		log.Fatal("Failed to create tasks table", err)
	}

	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	_, err = DB.Exec(usersTable)
	if err != nil {
		log.Fatal("Failed to create users table", err)
	}

	log.Println("Database connected and tables ready")
}
