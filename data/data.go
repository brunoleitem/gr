package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDb() error {
	var err error

	db, err = sql.Open("sqlite3", "./data.db3")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTables() {
	createSQL := `CREATE TABLE IF NOT EXISTS config (
		"key" TEXT
	);`

	stmt, err := db.Prepare(createSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
