package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseInitialization initializes the database
func DatabaseInitialization() (db *gorm.DB) {
	builtConnString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_CONNSTRING"),
		os.Getenv("DB_CONN_OPTIONALS"),
	)

	db, err := gorm.Open(mysql.Open(builtConnString), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

// CreateTable Creates table in Sqlite Sig (db *sql.DB, tableName string, fieldInfo ...string)
func CreateTable(db *gorm.DB, tableName string, fieldInfo ...string) {
	// TODO: Rework this using GORM
	// var fieldCreationString string = ""

	// for idx, field := range fieldInfo {
	// 	if idx != len(fieldInfo)-1 {
	// 		fieldCreationString = fieldCreationString + field + ", "
	// 	} else {
	// 		fieldCreationString = fieldCreationString + field
	// 	}
	// }

	// var creationString string = fmt.Sprintf(`
	// 	CREATE TABLE %s (%s)
	// `, tableName, fieldCreationString)

	// log.Println("Creating player table")

	// statement, err := db.Prepare(creationString)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// statement.Exec()
	// log.Println("player table created!")
}
