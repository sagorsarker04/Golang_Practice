package service

import (
	"fmt"

	"test/models"
)

// UserService interface
type UserService interface {
	GetUser(id int) (*model.User, error)
}

// userService struct
type userService struct{}

// Constructor
func NewUserService() UserService {
	return &userService{}
}

// Implement GetUser method
func (s *userService) GetUser(id int) (*model.User, error) {
	if id == 1 {
		return &model.User{ID: 1, Name: "Tanvir", Email: "tanvir@example.com"}, nil
	}
	return nil, fmt.Errorf("user with id %d not found", id)
}
