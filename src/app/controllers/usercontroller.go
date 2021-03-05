package controller

import (
	"database/sql"
	"net/http"

	"fbhc.com/api/main/db"
	"fbhc.com/api/main/routing"
)

func GetRoutes(database *sql.DB) []routing.Route {

	db.CreateTable(
		database,
		"users",
		`"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT`,
		`"firstname" TEXT`,
		`"lastname" TEXT`,
		`"email" integer`,
		`"password" string`,
	)

	playerRoutes := []routing.Route{
		{
			Path:    "/users",
			Method:  "GET",
			Handler: addNewUser(database),
		},
	}

	return playerRoutes
}

func addNewUser(db *sql.DB) http.HandlerFunc {
	// Maybe middleware here?
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
