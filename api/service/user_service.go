package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rGustave07/fbhcapi/api/model"
)

// UserService acts as a struct for injecting an implementation of
// UserRepository for use in service methods
type UserService struct {
	UserRepository model.UserRepository
}

// USConfig will hold repositories that will eventually be injected
// into this service layer
type USConfig struct {
	UserRepository model.UserRepository
}

// NewUserService is a factor function for
// initializing a UserService with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)

	return u, err
}

func (s *UserService) Signup(ctx context.Context, u *model.User) error {
	panic("Method not implemented")
}
