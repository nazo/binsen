package services

import (
	"testing"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/repositories"
	"github.com/nazo/binsen/server/app/services/auth"
	"golang.org/x/oauth2"
)

type mockUsersRepository struct {
	repositories.UsersRepository
}

// GetUserByEmail get user by email
func (r *mockUsersRepository) GetUserByEmail(email string) (*orm.User, error) {
	return &orm.User{Email: email}, nil
}

// CreateUser create user
func (r *mockUsersRepository) CreateUser(name, email string) (*orm.User, error) {
	return &orm.User{Name: name, Email: email}, nil
}

// UpdateUser update user
func (r *mockUsersRepository) UpdateUser(user *orm.User, name, email string) (*orm.User, error) {
	user.Name = name
	user.Email = email
	return user, nil
}

// DeleteUser update user
func (r *mockUsersRepository) DeleteUser(user *orm.User) error {
	user.Name = "deleted"
	user.Email = ""
	return nil
}

// DeleteUserByEmail delete user by email
func (r *mockUsersRepository) DeleteUserByEmail(email string) error {
	return nil
}

func mockUserService() UserService {
	return &userService{
		usersRepository: &mockUsersRepository{},
	}
}

func TestCreateUser(t *testing.T) {
	name := "test"
	email := "test@example.com"
	user, _ := mockUserService().CreateUser(name, email, nil, nil)
	if user.Email != email {
		t.Fatalf("user.Email expected %s, actual %s", user.Email, email)
	}
}

func TestUpdateUser(t *testing.T) {
	name := "test"
	email := "test@example.com"
	user := &orm.User{
		Name:  "nochange",
		Email: "nochange@example.com",
	}
	_, _ = mockUserService().UpdateUser(user, name, email)
	if user.Email != email {
		t.Fatalf("user.Email expected %s, actual %s", user.Email, email)
	}
}

func TestDeleteUser(t *testing.T) {
	user := &orm.User{
		Name:  "nochange",
		Email: "nochange@example.com",
	}
	_ = mockUserService().DeleteUser(user)
	if user.Email != "" {
		t.Fatalf("user.Email expected %s, actual %s", user.Email, "")
	}
}

func TestFindUserByGoogleUser(t *testing.T) {
	token := &oauth2.Token{}
	googleUser := &auth.GoogleUserinfoResponse{Email: "test@example.com"}
	user, _ := mockUserService().FindUserByGoogleUser(token, googleUser)
	if user.Email != googleUser.Email {
		t.Fatalf("user.Email expected %s, actual %s", user.Email, googleUser.Email)
	}
}
