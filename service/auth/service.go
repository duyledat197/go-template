package auth

import (
	"github.com/stamp-server/models"
)

// Service interface
type Service interface {
	Login(userName string, password string) (models.User, string, error)
	Register(user *models.User, password string) error
}

type service struct {
	userRepo models.UserRepository
}

func (s *service) Login(userName string, password string) (models.User, string, error) {
	user, err := s.userRepo.FindByUserName(userName)
	if err != nil {
		return models.User{}, "", err
	}
	match := models.IsCorrectPassword(password, user.HashedPassword)
	if match != true {
		return models.User{}, "", models.ErrWrongUserNameOrPassword
	}
	return *user, "", nil
}

func (s *service) Register(user *models.User, password string) error {
	_, err := s.userRepo.FindByUserName(user.UserName)
	if err == nil {
		return models.ErrUserAlreadyExist
	}
	if err != models.ErrUnknowUser {
		return err
	}
	user.HashedPassword, err = models.HashPassword(password)
	if err != nil {
		return err
	}
	err = s.userRepo.Create(user)
	return err
}

// NewService ...
func NewService(userRepo models.UserRepository) Service {
	return &service{
		userRepo: userRepo,
	}
}
