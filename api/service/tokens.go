package service

import (
	"crypto/rsa"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rGustave07/fbhcapi/api/model"
)

// IDTokenCustomClaims holds structure of JWT Claims of idToken
type IDTokenCustomClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

// RefreshTokenCustomClaims holds the payload of a refresh token
// This can be used to extract user id for subsequent
// application operations (IE, fetch user in Redis)
type RefreshTokenCustomClaims struct {
	UID uuid.UUID `json:"uid"`
	jwt.StandardClaims
}

// RefreshToken holds the actual signed jwt string along with the ID
// We return the id so it can be used without re-parsing the JWT from signed string
type RefreshToken struct {
	SS        string
	ID        string
	ExpiresIn time.Duration
}

// generateIDToken generates an IDToken which is a jwt with myCustomClaims
// Could call this GenerateIDTokenString, but the signature makes this fairly clear
func generateIDToken(u *model.User, key *rsa.PrivateKey) (string, error) {
	unixTime := time.Now().Unix()
	tokenExp := unixTime + 60*15 // 15 Mins from current time

	claims := IDTokenCustomClaims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  unixTime,
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)

	if err != nil {
		log.Println("Failed to sign id token string")
		return "", err
	}

	return ss, nil
}

// generateRefreshToken creates a refresh token
// The refresh token stores only the user's ID, a string
func generateRefreshToken(uid uuid.UUID, key string) (*RefreshToken, error) {
	currentTime := time.Now()

	tokenExp := currentTime.AddDate(0, 0, 3) // 3 days
	tokenID, err := uuid.NewRandom()         // v4 uuid in google uuid lib

	if err != nil {
		log.Println("Failed to generate refresh token ID")
		return nil, err
	}

	claims := RefreshTokenCustomClaims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        tokenID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		log.Println("Failed to sign refresh token string")
		return nil, err
	}

	return &RefreshToken{
		SS:        ss,
		ID:        tokenID.String(),
		ExpiresIn: tokenExp.Sub(currentTime),
	}, nil
}
