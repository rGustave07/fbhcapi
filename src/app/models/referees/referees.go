package referees

import (
	"fmt"
	"net/http"

	"fbhc.com/api/main/routing"
)

// Perhaps make a referee controller that contains a route struct

// Referee is not a player and has no number
type Referee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func refHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Big Fat Dicks")
}

func addRef(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding Ref")
}

func getSingleRef(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Single Ref")
}

// GetRoutes is a route generator
func GetRoutes() []routing.Route {
	routes := []routing.Route{
		// Get all refs
		{
			Path:    "/referees",
			Method:  "GET",
			Handler: refHandler,
		},
		// Add a ref
		{
			Path:    "/referees",
			Method:  "POST",
			Handler: addRef,
		},
		// Get a single ref
		{
			Path:    "/referee/{name}",
			Method:  "GET",
			Handler: getSingleRef,
		},
	}

	return routes
}
