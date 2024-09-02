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

	const create string = `
	  CREATE TABLE IF NOT EXISTS colaboradores (
	  id INTEGER NOT NULL PRIMARY KEY,
	  time DATETIME NOT NULL,
	  description TEXT
  	);`
	if _, err := DB.Exec(create); err != nil {
		log.Fatal(err)
	}
}
