package data

import (
	"os"
	"testing"
)

func TestOpenDb(t *testing.T) {
	_, err := os.Stat("./data.db3")
	if err != nil {
		os.Remove("./data.db3")
	}
	err = OpenDb()
	if err != nil {
		t.Fatalf("OpenDb failed: %v", err)
	}
	defer DB.Close()
	_, err = os.Stat("./data.db3")
	if os.IsNotExist(err) {
		t.Fatalf("Db creation failed: %v", err)
	}
}

func TestCreateTables(t *testing.T) {
	err := OpenDb()
	if err != nil {
		t.Fatalf("OpenDb failed: %v", err)
	}
	defer DB.Close()
	CreateTables()
	rows, err := DB.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='config';")
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		t.Fatalf("Table 'config' not created")
	}
}
