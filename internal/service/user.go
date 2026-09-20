package service

import (
	"github.com/aruncs31s/go_gin_api_starterkit/internal/application/dto"
	"github.com/aruncs31s/go_gin_api_starterkit/internal/domain/models"
	"github.com/aruncs31s/go_gin_api_starterkit/internal/repository"
)

type UserService interface {
	// CreateUser creates a new user.
	CreateUser(name string) (*dto.UserResponse, error)

	// GetUserByID retrieves a user by their ID.
	GetUserByID(id int) (*dto.UserResponse, error)

	// UpdateUser updates an existing user's information.
	UpdateUser(id int, name string) (*dto.UserResponse, error)

	// DeleteUser deletes a user by their ID.
	DeleteUser(id int) error

	// GetAllUsers retrieves all users.
	GetAllUsers() ([]*dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(name string) (*dto.UserResponse, error) {
	user := &models.User{Name: name}
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	return &dto.UserResponse{ID: user.ID, Name: user.Name}, nil
}

func (s *userService) GetUserByID(id int) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{ID: user.ID, Name: user.Name}, nil
}

func (s *userService) UpdateUser(id int, name string) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	if err := s.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}
	return &dto.UserResponse{ID: user.ID, Name: user.Name}, nil
}

func (s *userService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}

func (s *userService) GetAllUsers() ([]*dto.UserResponse, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userResponses := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = &dto.UserResponse{ID: user.ID, Name: user.Name}
	}
	return userResponses, nil
}
