package services

import (
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/repositories"
	"github.com/nazo/binsen/server/app/services/auth"
	"golang.org/x/oauth2"
)

// UserService todo
type UserService interface {
	CreateUser(name, email string, role *orm.Role, group *orm.Group) (*orm.User, error)
	UpdateUser(user *orm.User, name, email string) (*orm.User, error)
	DeleteUser(user *orm.User) error
	FindUserByGoogleUser(token *oauth2.Token, googleUser *auth.GoogleUserinfoResponse) (*orm.User, error)
	GetUser(id int64) (*orm.User, error)
	GetUsers() ([]*orm.User, error)
}

// userService todo
type userService struct {
	UserService
	usersRepository repositories.UsersRepository
}

// NewUserService todo
func NewUserService() UserService {
	return &userService{
		usersRepository: repositories.NewUsersRepository(),
	}
}

// CreateUser create user
func (s *userService) CreateUser(name, email string, role *orm.Role, group *orm.Group) (*orm.User, error) {
	return s.usersRepository.CreateUser(name, email)
}

// UpdateUser update user
func (s *userService) UpdateUser(user *orm.User, name, email string) (*orm.User, error) {
	return s.usersRepository.UpdateUser(user, name, email)
}

// FindUserByGoogleUser find user by google token
func (s *userService) FindUserByGoogleUser(token *oauth2.Token, googleUser *auth.GoogleUserinfoResponse) (*orm.User, error) {
	user, err := s.usersRepository.GetUserByEmail(googleUser.Email)
	if err != nil {
		return nil, err
	}
	// todo update user
	return user, nil
}

// GetUser find user by google token
func (s *userService) GetUser(id int64) (*orm.User, error) {
	user, err := s.usersRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers find user by google token
func (s *userService) GetUsers() ([]*orm.User, error) {
	users, err := s.usersRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) DeleteUser(user *orm.User) error {
	return s.usersRepository.DeleteUser(user)
}
