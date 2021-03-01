package db

import (
	"database/sql"
	"log"
	"os"
)

// DatabaseInitialization initializes the database
func DatabaseInitialization() {
	os.Remove("sqlite-database.db")
	log.Println("Creating sqlite-database.db")

	file, err := os.Create("sqlite-database.db")

	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	log.Println("sqlite database created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer sqliteDatabase.Close()

	createTable(sqliteDatabase)

	// Test Insertion
	insertPlayer(sqliteDatabase, 1, "Michael", "Franco", 13)
	insertPlayer(sqliteDatabase, 2, "Jordan", "Schmidt", 31)
	insertPlayer(sqliteDatabase, 3, "Troy", "Thompson", 28)
}

func createTable(db *sql.DB) {
	createSeededPlayers := `CREATE TABLE player (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"firstname" TEXT,
		"lastname" TEXT,
		"number" integer
	)`

	log.Println("Creating player table")

	statement, err := db.Prepare(createSeededPlayers)
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
