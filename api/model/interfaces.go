package model

import (
	"context"

	"github.com/google/uuid"
)

// UserService defines methods the handler layer expects
// any user service it interacts with to implement
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
}

type TokenService interface {
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
}

// UserRepository defines methods the service layer expects
// any user repository it interacts with to implement
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}
