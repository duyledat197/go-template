package auth

import (
	"github.com/google/uuid"
	"github.com/stamp-server/models"
	"github.com/stamp-server/utils"
	"github.com/stamp-server/utils/helper"
)

// Service interface
type Service interface {
	Login(email, password string) (models.User, string, error)
	Register(user *models.User, password string) error
}

type service struct {
	userRepo models.UserRepository
}

func (s *service) Login(email, password string) (models.User, string, error) {
	isEmail := helper.IsEmail(email)
	if isEmail != true {
		return models.User{}, "", models.ErrInvalidEmail
	}
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return models.User{}, "", err
	}
	isCorrect := models.IsCorrectPassword(password, user.HashedPassword)
	if isCorrect != true {
		return models.User{}, "", models.ErrWrongEmailOrPassword
	}

	tokenString, err := utils.GenerateToken(user.ID)
	if err != nil {
		return models.User{}, "", err
	}
	return *user, tokenString, nil
}

func (s *service) Register(user *models.User, password string) error {
	isEmail := helper.IsEmail(user.Email)
	if isEmail != true {
		return models.ErrInvalidEmail
	}
	isPassword := helper.IsPassword(password)
	if isPassword != true {
		return models.ErrInvalidPassword
	}
	_, err := s.userRepo.FindByEmail(user.Email)
	if err == nil {
		return models.ErrUserAlreadyExist
	}
	if err != models.ErrUnknowUser {
		return err
	}

	user.HashedPassword, err = models.HashPassword(password)
	uuidNew, err := uuid.NewUUID()
	user.ID = uuid.UUID.String(uuidNew)
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
