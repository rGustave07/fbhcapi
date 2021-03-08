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
