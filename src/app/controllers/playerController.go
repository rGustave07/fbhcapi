package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	model "fbhc.com/api/main/models/player"
	"gorm.io/gorm"
)

func unmarshalJSON(incData []byte) *model.Player {
	var data *model.Player

	if err := json.Unmarshal(incData, &data); err != nil {
		panic(err)
	}

	return data
}

// GetAllPlayers returns all registered players
func GetAllPlayers(db *gorm.DB) http.HandlerFunc {

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
func AddPlayer(db *gorm.DB) http.HandlerFunc {

	// SPECIFIC Middleware can go here

	return func(w http.ResponseWriter, r *http.Request) {
		request, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		playerData := unmarshalJSON(request)
		playerData.InsertPlayer(db)
	}
}
