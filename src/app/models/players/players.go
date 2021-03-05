package players

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"fbhc.com/api/main/db"
	"fbhc.com/api/main/routing"
)

// Player is a player
type Player struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Number    int    `json:"number"`
}

var players = []Player{
	{
		FirstName: "Michael",
		LastName:  "Franco",
		Number:    13,
	},
	{
		FirstName: "Ritter",
		LastName:  "Gustave",
		Number:    77,
	},
}

func unmarshalJSON(incData []byte) Player {
	var data Player
	if err := json.Unmarshal(incData, &data); err != nil {
		panic(err)
	}

	return data
}

// GetPlayers returns all registered players
func getAllPlayers(db *sql.DB) http.HandlerFunc {

	// Middleware can go here

	return func(w http.ResponseWriter, r *http.Request) {
		jsonMarshalledPlayers, err := json.Marshal(players)
		fmt.Println("Route hit")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMarshalledPlayers)
	}
}

func addPlayer(db *sql.DB) http.HandlerFunc {

	// Middleware can go here

	return func(w http.ResponseWriter, r *http.Request) {
		request, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		playerData := unmarshalJSON(request)
		players = append(players, playerData)

		log.Println("Inserting Player")
		insertPlayerQuery := `INSERT INTO player(firstname, lastname, number) VALUES (?, ?, ?)`

		statement, statementErr := db.Prepare(insertPlayerQuery)

		if statementErr != nil {
			log.Fatalln(statementErr.Error())
		}

		_, queryErr := statement.Exec(playerData.FirstName, playerData.LastName, playerData.Number)

		if queryErr != nil {
			log.Fatalln(queryErr.Error())
		}

		fmt.Println("Player successfully added")
		fmt.Printf("%+v\n", playerData)
	}
}

// GetRoutes returns the routes metainfo and controllers
func GetRoutes(database *sql.DB) []routing.Route {

	db.CreateTable(
		database,
		"player",
		`"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT`,
		`"firstname" TEXT`,
		`"lastname" TEXT`,
		`"number" integer`,
	)

	playerRoutes := []routing.Route{
		{
			Path:    "/players",
			Method:  "GET",
			Handler: getAllPlayers(database),
		},
		{
			Path:    "/player/addPlayer",
			Method:  "POST",
			Handler: addPlayer(database),
		},
	}

	return playerRoutes
}
