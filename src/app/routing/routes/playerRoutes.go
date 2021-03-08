package routes

import (
	"database/sql"

	controller "fbhc.com/api/main/controllers"
	"fbhc.com/api/main/db"
	"fbhc.com/api/main/routing"
)

// GetPlayerRoutes ...
func GetPlayerRoutes(database *sql.DB) []routing.Route {
	db.CreateTable(
		database,
		"player",
		`"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT`,
		`"firstname" TEXT`,
		`"lastname" TEXT`,
		`"number" integer`,
	)

	playerRoutes := []routing.Route{
		{
			Path:    "/players",
			Method:  "GET",
			Handler: controller.GetAllPlayers(database),
		},
		{
			Path:    "/player/addPlayer",
			Method:  "POST",
			Handler: controller.AddPlayer(database),
		},
	}

	return playerRoutes
}
