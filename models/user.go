package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                 string `bson:"_id"`
	Email              string `bson:"email"`
	HashedPassword     string `bson:"hashedPassword"`
	StampCreatedAmount int    `bson:"stampCreatedAmount"`
}

// HashPassword : hash password using crypto
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// IsCorrectPassword : check password with passwordhash
func IsCorrectPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// UserRepository ...
type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByUserID(userID string) (*User, error)
	FindAll() ([]*User, error)
	Update(userID string, user *User) error
	Delete(userID string) error
}

// ErrUnknowUser ...
var ErrUnknowUser = errors.New("unknown user")

// ErrUserAlreadyExist ...
var ErrUserAlreadyExist = errors.New("user already exist")

// ErrWrongEmailOrPassword ...
var ErrWrongEmailOrPassword = errors.New("wrong email or password")

// ErrInvalidEmail ...
var ErrInvalidEmail = errors.New("invalid email")

// ErrInvalidPassword ...
var ErrInvalidPassword = errors.New("password must in 8 - 32 characters")
