package players

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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

// UnmarshalJSON does Unmarshalling function for json to struct
func UnmarshalJSON(incData []byte) Player {
	var data Player
	if err := json.Unmarshal(incData, &data); err != nil {
		panic(err)
	}

	return data
}

// GetPlayers returns all registered players
func getAllPlayers(w http.ResponseWriter, r *http.Request) {
	jsonMarshalledPlayers, err := json.Marshal(players)
	fmt.Println("Route hit")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMarshalledPlayers)
}

// AddPlayer adds player to roster
func addPlayer(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	playerData := UnmarshalJSON(request)
	players = append(players, playerData)

	fmt.Println("Player successfully added")
	fmt.Printf("%+v\n", playerData)
}

// InitializeEndpoints endpoints for this package
func InitializeEndpoints(r *mux.Router) {
	r.HandleFunc("/players", getAllPlayers).Methods("GET")
	r.HandleFunc("/players", addPlayer).Methods("POST")
}
