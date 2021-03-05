package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"fbhc.com/api/main/db"
	"fbhc.com/api/main/models/players"
	"fbhc.com/api/main/models/referees"
	"fbhc.com/api/main/routing"
)

// App is an instance of the server application
type App struct {
	Router *mux.Router
	DB     *sql.DB
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

	routes := append(
		players.GetRoutes(a.DB),
		referees.GetRoutes()...,
	)

	a.registerRoutes(routes)

	return a
}

// Run application
func (a *App) Run(addr string) {
	fmt.Println("Application is now being served on :8080")

	defer func() {
		log.Fatal(a.DB.Close())
	}()

	log.Fatal(http.ListenAndServe(addr, a.Router))
}
