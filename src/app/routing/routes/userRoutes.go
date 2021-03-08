package routes

import (
	controller "fbhc.com/api/main/controllers"
	"fbhc.com/api/main/routing"
	"gorm.io/gorm"
)

// GetUserRoutes ...
func GetUserRoutes(database *gorm.DB) []routing.Route {

	userRoutes := []routing.Route{
		{
			Path: "/user",
			// this is a get for testing only
			Method: "GET",
			// Method:  "POST",
			Handler: controller.AddNewUser(database),
		},
		{
			Path:    "/getusers",
			Method:  "GET",
			Handler: controller.GetUsers(database),
		},
	}

	return userRoutes
}
