package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	model "fbhc.com/api/main/models/player"
)

func unmarshalJSON(incData []byte) model.Player {
	var data model.Player
	if err := json.Unmarshal(incData, &data); err != nil {
		panic(err)
	}

	return data
}

// GetAllPlayers returns all registered players
func GetAllPlayers(db *sql.DB) http.HandlerFunc {

	// Middleware can go here

	return func(w http.ResponseWriter, r *http.Request) {
		// refactor to use database
		// jsonMarshalledPlayers, err := json.Marshal(players)
		// fmt.Println("Route hit")

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

		// w.Header().Set("Content-Type", "application/json")
		// w.Write(jsonMarshalledPlayers)
		fmt.Println("getAllPlayers controller")
	}
}

// AddPlayer adds a player
func AddPlayer(db *sql.DB) http.HandlerFunc {

	// Middleware can go here

	return func(w http.ResponseWriter, r *http.Request) {
		// Refactor to use database
		// request, err := ioutil.ReadAll(r.Body)

		// if err != nil {
		// 	panic(err)
		// }

		// playerData := unmarshalJSON(request)
		// players = append(players, playerData)

		// log.Println("Inserting Player")
		// insertPlayerQuery := `INSERT INTO player(firstname, lastname, number) VALUES (?, ?, ?)`

		// statement, statementErr := db.Prepare(insertPlayerQuery)

		// if statementErr != nil {
		// 	log.Fatalln(statementErr.Error())
		// }

		// _, queryErr := statement.Exec(playerData.FirstName, playerData.LastName, playerData.Number)

		// if queryErr != nil {
		// 	log.Fatalln(queryErr.Error())
		// }

		// fmt.Println("Player successfully added")
		// fmt.Printf("%+v\n", playerData)
		fmt.Println("AddPlayer controller")
	}
}
