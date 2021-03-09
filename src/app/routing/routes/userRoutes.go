package routes

import (
	controller "fbhc.com/api/main/controllers"
	model "fbhc.com/api/main/models/users"
	"fbhc.com/api/main/routing"
	"gorm.io/gorm"
)

// GetUserRoutes ...
func GetUserRoutes(database *gorm.DB) []routing.Route {
	model.CreateUserTable(database)

	userRoutes := []routing.Route{
		{
			Path:    "/user/addUser",
			Method:  "POST",
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
