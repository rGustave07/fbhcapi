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
