package db

import (
	"database/sql"
	"log"
	"taskmanager/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func CreateUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL 
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create users database", err)
	}
}

func InsertUser(db *sql.DB, user models.User) error {
	query := `INSERT into users(username, password) VALUES (?, ?)`
	_, err := db.Exec(query, user.Username, user.Password)
	if err != nil {
		log.Println("Failed to add user", err)
		return err
	}
	return nil
}

func GetUserByName(db *sql.DB, username string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, password FROM users WHERE username=?`
	row := db.QueryRow(query, username)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}
