package services

import (
	"errors"
	"loyalty-program-api/models"
	"loyalty-program-api/repositories"
)

type UserService interface {
	RegisterUser(name, email string) (*models.User, error)
	GetPoints(userID uint) (float64, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(name, email string) (*models.User, error) {
	user := &models.User{Name: name, Email: email}
	err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetPoints(userID uint) (float64, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return 0, errors.New("user not found")
	}
	return user.Points, nil
}
