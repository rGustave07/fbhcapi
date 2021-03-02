package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"fbhc.com/api/main/players"
	"github.com/gorilla/mux"
)

// App is an instance of the server application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize upon app startup, this function should be called
func (a *App) Initialize() *App {
	a.Router = mux.NewRouter()

	for _, route := range players.GetRoutes() {
		a.Router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	return a
}

// Run application
func (a *App) Run(addr string) {
	fmt.Println("Application is now being served on :8080")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
