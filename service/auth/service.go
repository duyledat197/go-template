package auth

import (
	"github.com/duyledat197/go-template/models/domain"
	"github.com/duyledat197/go-template/utils"
	"github.com/duyledat197/go-template/utils/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service interface
type Service interface {
	Login(email, password string) (domain.User, string, error)
	Register(user *domain.User, password string) error
}

type service struct {
	userRepo domain.UserRepository
}

func (s *service) Login(email, password string) (domain.User, string, error) {
	isEmail := helper.IsEmail(email)
	if isEmail != true {
		return domain.User{}, "", domain.ErrInvalidEmail
	}
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return domain.User{}, "", err
	}
	isCorrect := domain.IsCorrectPassword(password, user.HashedPassword)
	if isCorrect != true {
		return domain.User{}, "", domain.ErrWrongEmailOrPassword
	}

	tokenString, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		return domain.User{}, "", err
	}
	return *user, tokenString, nil
}

func (s *service) Register(user *domain.User, password string) error {
	isEmail := helper.IsEmail(user.Email)
	if isEmail != true {
		return domain.ErrInvalidEmail
	}
	isPassword := helper.IsPassword(password)
	if isPassword != true {
		return domain.ErrInvalidPassword
	}
	_, err := s.userRepo.FindByEmail(user.Email)
	if err == nil {
		return domain.ErrUserAlreadyExist
	}
	if err != domain.ErrUnknowUser {
		return err
	}

	user.HashedPassword, err = domain.HashPassword(password)
	user.ID = primitive.NewObjectID()
	if err != nil {
		return err
	}
	err = s.userRepo.Create(user)
	return err
}

// NewService ...
func NewService(userRepo domain.UserRepository) Service {
	return &service{
		userRepo: userRepo,
	}
}
