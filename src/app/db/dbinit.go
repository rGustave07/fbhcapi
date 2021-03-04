package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// DatabaseInitialization initializes the database
func DatabaseInitialization() (db *sql.DB) {
	os.Remove("sqlite-database.db")
	log.Println("Creating sqlite-database.db")

	file, err := os.Create("sqlite-database.db")

	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	log.Println("sqlite database created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	return sqliteDatabase
}

// CreateTable Creates table in Sqlite Sig (db *sql.DB, tableName string, fieldInfo ...string)
func CreateTable(db *sql.DB, tableName string, fieldInfo ...string) {
	var fieldCreationString string = ""

	for idx, field := range fieldInfo {
		if idx != len(fieldInfo)-1 {
			fieldCreationString = fieldCreationString + field + ", "
		} else {
			fieldCreationString = fieldCreationString + field
		}
	}

	var creationString string = fmt.Sprintf(`
		CREATE TABLE %s (%s)
	`, tableName, fieldCreationString)

	log.Println("Creating player table")

	statement, err := db.Prepare(creationString)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("player table created!")
}

func insertPlayer(db *sql.DB, id int, firstname string, lastname string, number int) {
	log.Println("Inserting Player")
	insertPlayerQuery := `INSERT INTO player(id, firstname, lastname, number) VALUES (?, ?, ?, ?)`

	statement, err := db.Prepare(insertPlayerQuery)

	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(id, firstname, lastname, number)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
