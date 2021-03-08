package controller

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// AddNewUser does stuff
func AddNewUser(db *gorm.DB) http.HandlerFunc {
	// Maybe middleware here?
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("userController activated")
	}
}

// GetUsers ...
func GetUsers(db *gorm.DB) http.HandlerFunc {
	// Maybe middleware here?
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("userController activated")
	}
}

// GetRoutes retrieves routes
// func GetUserRoutes(database *sql.DB) []routing.Route {

// 	db.CreateTable(
// 		database,
// 		"users",
// 		`"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT`,
// 		`"firstname" TEXT`,
// 		`"lastname" TEXT`,
// 		`"email" integer`,
// 		`"password" string`,
// 	)

// 	playerRoutes := []routing.Route{
// 		{
// 			Path:    "/users",
// 			Method:  "GET",
// 			Handler: addNewUser(database),
// 		},
// 	}
