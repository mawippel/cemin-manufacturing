package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	// Run migrations or initial setup
}
