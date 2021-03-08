package referees

// Perhaps make a referee controller that contains a route struct

// Referee is not a player and has no number
type Referee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
