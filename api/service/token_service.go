package service

import (
	"context"
	"crypto/rsa"

	"github.com/rGustave07/fbhcapi/api/model"
)

// TokenService used for injecting an implementation of TokenRepository
// for use in service methods along with keys and secrets for signing JWTs
type TokenService struct {
	// TokenRepository model.TokenRepository
	PrivKey       *rsa.PrivateKey
	PubKey        *rsa.PublicKey
	RefreshSecret string
}

// TSConfig will hold repositories that will eventually be injected
// into this service layer
type TSConfig struct {
	// TokenRepository model.TokenRepository
	PrivKey       *rsa.PrivateKey
	PubKey        *rsa.PublicKey
	RefreshSecret string
}

// NewTokenService is a factory function for
// initializing a userService with its repository layer dependency
func NewTokenService(c *TSConfig) model.TokenService {
	return &TokenService{
		PrivKey:       c.PrivKey,
		PubKey:        c.PubKey,
		RefreshSecret: c.RefreshSecret,
	}
}

// NewPairFromUser creates fresh id and refresh tokens for the current user
// If a previous token is included, the previous token is removed from
// the tokens repository
func (s *TokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	panic("not implemented")
}
