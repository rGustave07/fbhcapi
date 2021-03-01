package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"fbhc.com/api/main/db"
	"fbhc.com/api/main/routing"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	routingInfo := routing.InitializeRoutes()
	db.DatabaseInitialization()

	fmt.Println("Server has started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", routingInfo))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}
