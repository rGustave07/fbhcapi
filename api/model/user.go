package model

import (
	"github.com/google/uuid"
)

type User struct {
	UID       uuid.UUID `db:"uid" json:"uid"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	Firstname string    `db:"firstname" json:"firstname"`
	Lastname  string    `db:"lastname" json:"lastname"`
	ImageURL  string    `db:"image_url" json:"imageUrl"`
	Website   string    `db:"website" json:"website"`
}
