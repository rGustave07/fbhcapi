package routes

import (
	controller "fbhc.com/api/main/controllers"
	"fbhc.com/api/main/routing"
	"gorm.io/gorm"
)

// GetRefereeRoutes ...
func GetRefereeRoutes(database *gorm.DB) []routing.Route {
	// Create a table that's needed for the data

	refereeRoutes := []routing.Route{
		{
			Path:    "/referee",
			Method:  "GET",
			Handler: controller.GetReferees(),
		},
		{
			Path:    "/referee",
			Method:  "POST",
			Handler: controller.AddReferee(),
		},
	}

	return refereeRoutes
}
