package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Could not connect to the database")
	}

	// Verify connection
	if err = DB.Ping(); err != nil {
		panic("Database is not reachable")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsertable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsertable)
	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS createRegistrationTable (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
		FOREIGN KEY(event_id) REFERENCES evetns(id)
	);
	`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("Could not create events table")
	}

}
