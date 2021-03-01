package routing

import (
	"fmt"
	"html"
	"net/http"

	"fbhc.com/api/main/routing/players"
	"github.com/gorilla/mux"
)

// InitializeRoutes is where routes are registered with their event handler
func InitializeRoutes() *mux.Router {
	fmt.Println("Initializing Routes")

	r := mux.NewRouter()

	r.HandleFunc("/", index)

	players.InitializeEndpoints(r)
	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}
