package routes

import (
	controller "fbhc.com/api/main/controllers"
	"fbhc.com/api/main/routing"
	"gorm.io/gorm"
)

// GetPlayerRoutes ...
func GetPlayerRoutes(database *gorm.DB) []routing.Route {

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
