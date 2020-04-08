package wallet

import (
	"testing"

	"github.com/stamp-server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWallet_Service_UserIDNotExist(t *testing.T) {
	walletRepo := &mockWalletRepository{}
	userRepo := &mockUserRepository{}
	service := NewService(walletRepo, userRepo)
	userID := "userID"
	userRepo.On("FindByUserID", userID).Return(nil, models.ErrUnknowUser, nil)
	address, publicKey, err := service.Create(userID)
	assert.Equal(t, models.ErrUnknowUser, err)
	assert.Equal(t, "", address)
	assert.Equal(t, "", publicKey)
}

// mock

type mockWalletRepository struct {
	mock.Mock
}

func (m *mockWalletRepository) Create(wallet *models.Wallet) error {
	args := m.Called(wallet)
	return args.Error(1)
}

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(1)
}

func (m *mockUserRepository) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	var r0 *models.User
	if rf, ok := args.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
func (m *mockUserRepository) FindByUserID(userID string) (*models.User, error) {
	args := m.Called(userID)
	var r0 *models.User
	if rf, ok := args.Get(0).(func(string) *models.User); ok {
		r0 = rf(userID)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
func (m *mockUserRepository) FindAll() ([]*models.User, error) {
	return []*models.User{}, nil
}
func (m *mockUserRepository) Update(userID string, user *models.User) error {
	return nil
}
func (m *mockUserRepository) Delete(userID string) error {
	return nil
}
