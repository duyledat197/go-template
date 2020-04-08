package user

import (
	"github.com/stamp-server/models"
)

// Service interface
type Service interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(userID string) (models.User, error)
	GetListUser() ([]models.User, error)
	UpdateUser(userID string, user models.User) error
	DeleteUser(userID string) error
}

type service struct {
	userRepo models.UserRepository
}

func (s *service) CreateUser(user models.User) error {
	err := s.userRepo.Create(&user)
	return err
}

func (s *service) GetUserByEmail(email string) (models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	return *user, err
}

func (s *service) GetUserByID(userID string) (models.User, error) {
	user, err := s.userRepo.FindByUserID(userID)
	if err != nil {
		return models.User{}, err
	}
	return *user, err
}

func (s *service) GetListUser() ([]models.User, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var result []models.User
	for _, user := range users {
		result = append(result, *user)
	}
	return result, nil
}

func (s *service) UpdateUser(userID string, user models.User) error {
	err := s.userRepo.Update(userID, &user)
	return err
}

func (s *service) DeleteUser(userID string) error {
	err := s.userRepo.Delete(userID)
	return err
}

// NewService ...
func NewService(userRepo models.UserRepository) Service {
	return &service{
		userRepo: userRepo,
	}
}
