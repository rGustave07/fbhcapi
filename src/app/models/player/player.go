package model

// Player is a player
type Player struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Number    int    `json:"number"`
}
