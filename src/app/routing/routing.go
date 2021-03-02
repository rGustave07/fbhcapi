package routing

import (
	"net/http"
)

// Route contains information on a specific route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}
