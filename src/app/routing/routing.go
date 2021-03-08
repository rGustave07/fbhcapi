package routing

import (
	"net/http"
)

// Route contains information on a specific route
type Route struct {
	Path   string
	Method string
	// A function that returns a function to handle the http request
	Handler http.HandlerFunc
}

// RouteCombiner will take a slice of routes and flatten them
func RouteCombiner(routes ...[]Route) []Route {
	var flattenedRoutes []Route

	for _, routeSlice := range routes {
		flattenedRoutes = append(flattenedRoutes, routeSlice...)
	}

	return flattenedRoutes
}
