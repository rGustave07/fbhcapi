package model

import (
	"fmt"
	"time"

	"fbhc.com/api/main/utility"
	"gorm.io/gorm"
)

//User Defines the fields for a user
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

// HashOwnPassword ...
func (u *User) HashOwnPassword() (string, error) {
	hashedPassword, err := utility.HashString(u.Password)

	u.Password = hashedPassword

	return hashedPassword, err
}

// CreateUserTable ...
func CreateUserTable(db *gorm.DB) {
	tableExists := db.Migrator().HasTable(&User{})

	// Pretty self explanatory code - Ritter Gustave
	if tableExists {
		return
	}

	fmt.Println("User table not found, creating table")
	err := db.Migrator().CreateTable(&User{})

	if err != nil {
		panic(err.Error())
	}

	return
}
