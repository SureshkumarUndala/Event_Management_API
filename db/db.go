package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "Suresh@123",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "suresh",
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", cfg.FormatDSN()) // creates a database object
	if err != nil {

		panic(err.Error())
	}

	DB.SetMaxOpenConns(10) // controlling howmany connections can be open simultaneously atmost
	DB.SetMaxIdleConns(5)  // how many connections will be open if no using these connections
	createtables()
}

func createtables() {
	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		 id INTEGER PRIMARY KEY AUTO_INCREMENT,
		 name TEXT NOT NULL,
		 description TEXT NOT NULL,
		 dateTime DATETIME NOT NULL, 
		 user_id INTEGER

	)`
	_, err := DB.Exec(createEventTable)

	if err != nil {

		panic(err.Error())
	}
}
