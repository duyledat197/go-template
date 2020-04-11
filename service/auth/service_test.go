package auth

import (
	"testing"

	"github.com/duyledat197/go-template/models/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userLoginRequest struct {
	Email    string
	Password string
}

func TestRegister_Service_InvalidEmail(t *testing.T) {
	var userRegister domain.User = domain.User{
		Email: "test002",
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)
	err := service.Register(&userRegister, "password")
	assert.Equal(t, domain.ErrInvalidEmail, err)
}

func TestRegister_Service_InvalidPassword(t *testing.T) {
	var userRegister domain.User = domain.User{
		Email: "test002@yopmail.com",
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)
	err := service.Register(&userRegister, "")
	assert.Equal(t, domain.ErrInvalidPassword, err)
}

func TestRegister_Service_ExistEmail(t *testing.T) {
	var userRegister domain.User = domain.User{
		Email: "test001@yopmail.com",
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)
	userRepo.On("Create", &userRegister).Return(domain.ErrUserAlreadyExist, nil)
	userRepo.On("FindByEmail", userRegister.Email).Return(&userRegister, nil, nil)
	err := service.Register(&userRegister, "password")
	assert.Equal(t, domain.ErrUserAlreadyExist, err)
}

func TestRegister_Service_Success(t *testing.T) {
	var userRegister domain.User = domain.User{
		Email: "test002@yopmail.com",
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)
	userRepo.On("Create", &userRegister).Return(nil, nil)
	userRepo.On("FindByEmail", userRegister.Email).Return(nil, domain.ErrUnknowUser, nil)
	err := service.Register(&userRegister, "password")
	assert.Equal(t, nil, err)
}

func TestLogin_Service_InvalidEmail(t *testing.T) {
	var userLogin userLoginRequest = userLoginRequest{
		Email:    "test001yopmail.com",
		Password: "Test@123",
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)
	user, token, err := service.Login(userLogin.Email, userLogin.Password)
	assert.Equal(t, domain.ErrInvalidEmail, err)
	assert.Equal(t, domain.User{}, user)
	assert.Equal(t, "", token)
}

func TestLogin_Service_EmptyEmail(t *testing.T) {
	var userLogin userLoginRequest = userLoginRequest{
		Email:    "",
		Password: "Test@123",
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)
	user, token, err := service.Login(userLogin.Email, userLogin.Password)
	assert.Equal(t, domain.ErrInvalidEmail, err)
	assert.Equal(t, domain.User{}, user)
	assert.Equal(t, "", token)
}

func TestLogin_Service_NotExistEmail(t *testing.T) {
	var userLogin userLoginRequest = userLoginRequest{
		Email:    "test001@yopmail.com",
		Password: "Test@123",
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)

	userRepo.On("FindByEmail", userLogin.Email).Return(nil, domain.ErrUnknowUser, nil)
	user, token, err := service.Login(userLogin.Email, userLogin.Password)
	assert.Equal(t, domain.ErrUnknowUser, err)
	assert.Equal(t, "", token)
	assert.Equal(t, domain.User{}, user)
}

func TestLogin_Service_WrongPassword(t *testing.T) {
	var userLogin userLoginRequest = userLoginRequest{
		Email:    "test001@yopmail.com",
		Password: "wrongpassword",
	}
	rightPassword := "Test@123"
	hashPassword, err := domain.HashPassword(rightPassword)
	expectUser := domain.User{
		Email:          userLogin.Email,
		HashedPassword: hashPassword,
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)
	userRepo.On("FindByEmail", userLogin.Email).Return(&expectUser, nil, nil)
	user, token, err := service.Login(userLogin.Email, userLogin.Password)
	assert.Equal(t, domain.ErrWrongEmailOrPassword, err)
	assert.Equal(t, "", token)
	assert.Equal(t, domain.User{}, user)
}

func TestLogin_Service_Success(t *testing.T) {
	var userLogin userLoginRequest = userLoginRequest{
		Email:    "test001@yopmail.com",
		Password: "Test@123",
	}
	hashPassword, err := domain.HashPassword(userLogin.Password)
	expectUser := domain.User{
		Email:          userLogin.Email,
		HashedPassword: hashPassword,
	}
	userRepo := &mockUserRepository{}
	service := NewService(userRepo)

	userRepo.On("FindByEmail", userLogin.Email).Return(&expectUser, nil, nil)
	user, token, err := service.Login(userLogin.Email, userLogin.Password)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", token)
	assert.NotEqual(t, domain.User{}, user)
}

// mock
type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(1)
}

func (m *mockUserRepository) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	var r0 *domain.User
	if rf, ok := args.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*domain.User)
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
func (m *mockUserRepository) FindByUserID(userID string) (*domain.User, error) {
	return nil, nil
}
func (m *mockUserRepository) FindAll() ([]*domain.User, error) {
	return []*domain.User{}, nil
}
func (m *mockUserRepository) Update(userID string, user *domain.User) error {
	return nil
}
func (m *mockUserRepository) Delete(userID string) error {
	return nil
}
