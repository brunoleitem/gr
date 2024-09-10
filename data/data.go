package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDb() error {
	var err error

	DB, err = sql.Open("sqlite3", "./data.db3")
	if err != nil {
		return err
	}

	return DB.Ping()
}

func CreateTables() {
	createSQL := `CREATE TABLE IF NOT EXISTS config (
		"key" TEXT
	);`

	stmt, err := DB.Prepare(createSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
