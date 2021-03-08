package model

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// Player is a player
type Player struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Number    int    `json:"number"`
}

// InsertPlayer ...
func (player *Player) InsertPlayer(db *gorm.DB) {
	insertionResult := db.Select("Firstname", "Lastname", "Number", "CreatedAt").Create(&player)

	if insertionResult.Error != nil {
		log.Fatal(insertionResult.Error)
	}

	resultantString := fmt.Sprintf("Rows Affected: %v", insertionResult.RowsAffected)
	fmt.Println(resultantString)
}

// GetAllPlayers ...
func (player *Player) GetAllPlayers(db *gorm.DB) []Player {
	var playerResults []Player

	result := db.Find(&playerResults)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, player := range playerResults {
		fmt.Printf("%+v\n", player)
	}

	return playerResults
}
