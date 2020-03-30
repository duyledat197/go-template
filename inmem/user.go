package inmem

import (
	"github.com/stamp-server/models"
)

type userRepository struct {
	users map[string]*models.User
}

func (r *userRepository) Create(user *models.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *userRepository) FindByUserName(userName string) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) FindByUserID(userID string) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) FindAll() ([]*models.User, error) {
	return nil, nil
}

func (r *userRepository) Update(userID string, user *models.User) error {
	return nil
}

func (r *userRepository) Delete(userID string) error {
	return nil
}

// NewUserRepository ...
func NewUserRepository() models.UserRepository {
	return &userRepository{
		users: make(map[string]*models.User),
	}
}
