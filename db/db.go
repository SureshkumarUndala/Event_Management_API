package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10) // controlling howmany connections can be open simultaneously atmost
	DB.SetMaxIdleConns(5)  // how many connections will be open if no using these connections
	createtables()
}

func createtables() {
	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		 id INTEGER PRIMARY KEY AUTOINCREMENT,
		 name TEXT NOT NULL,
		 description TEXT NOT NULL,
		 dateTime DATETIME NOT NULL, 
		 user_id INTEGER

	)`
	_, err := DB.Exec(createEventTable)

	if err != nil {

		panic("could not creates event table")
	}
}
