package utility

import "golang.org/x/crypto/bcrypt"

// HashString HashString hashes the given string.
func HashString(str string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashedBytes), err
}
