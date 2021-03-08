package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	model "fbhc.com/api/main/models/player"
	"gorm.io/gorm"
)

func marshalJSONMultiplePlayers(players []model.Player) []byte {
	jsonReady, err := json.Marshal(players)

	if err != nil {
		panic(err)
	}

	return jsonReady
}

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
		var p model.Player

		dbPlayers := p.GetAllPlayers(db)
		allPlayers := marshalJSONMultiplePlayers(dbPlayers)

		w.WriteHeader(http.StatusOK)
		w.Write(allPlayers)
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
