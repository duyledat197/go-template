package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                 string `bson:"_id"`
	UserName           string `bson:"userName"`
	HashedPassword     string `bson:"hashedPassword"`
	Name               string `bson:"name"`
	Address            string `bson:"address"`
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
	FindByUserName(userName string) (*User, error)
	FindByUserID(userID string) (*User, error)
	FindAll() ([]*User, error)
	Update(userID string, user *User) error
	Delete(userID string) error
}

// ErrUnknowUser ...
var ErrUnknowUser = errors.New("unknown user")

// ErrUserAlreadyExist ...
var ErrUserAlreadyExist = errors.New("user already exist")

// ErrWrongUserNameOrPassword ...
var ErrWrongUserNameOrPassword = errors.New("wrong user name or password")
