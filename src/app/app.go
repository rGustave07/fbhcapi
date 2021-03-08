package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"fbhc.com/api/main/db"
	"fbhc.com/api/main/routing"
	"fbhc.com/api/main/routing/routes"
)

// App is an instance of the server application
type App struct {
	Router *mux.Router
	DB     *gorm.DB
	routes []routing.Route
}

func (a *App) registerRoutes(r []routing.Route) {
	for _, route := range r {
		a.Router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}

// Initialize upon app startup, this function should be called
func (a *App) Initialize() *App {
	a.Router = mux.NewRouter().StrictSlash(false)
	a.DB = db.DatabaseInitialization()

	routes := routing.RouteCombiner(
		routes.GetPlayerRoutes(a.DB),
		routes.GetUserRoutes(a.DB),
		routes.GetRefereeRoutes(a.DB),
	)

	a.registerRoutes(routes)

	return a
}

// Run application
func (a *App) Run(addr string) {
	fmt.Println("Application is now being served on :8080")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
