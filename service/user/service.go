package user

import (
	"github.com/duyledat197/go-template/models/domain"
)

// Service interface
type Service interface {
	CreateUser(user domain.User) error
	GetUserByEmail(email string) (domain.User, error)
	GetUserByID(userID string) (domain.User, error)
	GetListUser() ([]domain.User, error)
	UpdateUser(userID string, user domain.User) error
	DeleteUser(userID string) error
}

type service struct {
	userRepo domain.UserRepository
}

func (s *service) CreateUser(user domain.User) error {
	err := s.userRepo.Create(&user)
	return err
}

func (s *service) GetUserByEmail(email string) (domain.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return *user, err
}

func (s *service) GetUserByID(userID string) (domain.User, error) {
	user, err := s.userRepo.FindByUserID(userID)
	if err != nil {
		return domain.User{}, err
	}
	return *user, err
}

func (s *service) GetListUser() ([]domain.User, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	var result []domain.User
	for _, user := range users {
		result = append(result, *user)
	}
	return result, nil
}

func (s *service) UpdateUser(userID string, user domain.User) error {
	err := s.userRepo.Update(userID, &user)
	return err
}

func (s *service) DeleteUser(userID string) error {
	err := s.userRepo.Delete(userID)
	return err
}

// NewService ...
func NewService(userRepo domain.UserRepository) Service {
	return &service{
		userRepo: userRepo,
	}
}
