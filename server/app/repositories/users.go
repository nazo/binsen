package repositories

import (
	"context"
	"database/sql"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/lib/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// UsersRepository todo
type UsersRepository interface {
	GetUsers() ([]*orm.User, error)
	GetUserByID(id int64) (*orm.User, error)
	GetUserByEmail(email string) (*orm.User, error)
	CreateUser(name, email string) (*orm.User, error)
	UpdateUser(user *orm.User, name, email string) (*orm.User, error)
	DeleteUserByEmail(email string) error
	DeleteUserByID(id int64) error
	DeleteUser(user *orm.User) error
}

// usersRepository todo
type usersRepository struct {
	UsersRepository
	db *sql.DB
}

// NewUsersRepository todo
func NewUsersRepository() UsersRepository {
	return &usersRepository{
		db: db.InjectDB(),
	}
}

// GetUserByEmail get user by email
func (r *usersRepository) GetUserByEmail(email string) (*orm.User, error) {
	return orm.Users(qm.Where("email = ?", email)).One(context.Background(), r.db)
}

// CreateUser create user
func (r *usersRepository) CreateUser(name, email string) (*orm.User, error) {
	user := &orm.User{
		Name:  name,
		Email: email,
	}
	err := user.Insert(context.Background(), r.db, boil.Infer())
	return user, err
}

// UpdateUser create user
func (r *usersRepository) UpdateUser(user *orm.User, name, email string) (*orm.User, error) {
	user.Name = name
	user.Email = email
	_, err := user.Update(context.Background(), r.db, boil.Infer())
	return user, err
}

// DeleteUser delete user
func (r *usersRepository) DeleteUser(user *orm.User) error {
	_, err := user.Delete(context.Background(), r.db)
	return err
}

// DeleteUser delete user by id
func (r *usersRepository) DeleteUserByID(id int64) error {
	user, err := r.GetUserByID(id)
	if err != nil {
		return err
	}
	return r.DeleteUser(user)
}

// DeleteUserByEmail delete user by email
func (r *usersRepository) DeleteUserByEmail(email string) error {
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return err
	}
	return r.DeleteUser(user)
}

// GetUserByID get user by id
func (r *usersRepository) GetUserByID(id int64) (*orm.User, error) {
	return orm.Users(qm.Where("id = ?", id)).One(context.Background(), r.db)
}

// GetUsers get users
func (r *usersRepository) GetUsers() ([]*orm.User, error) {
	return orm.Users(qm.OrderBy("id ASC")).All(context.Background(), r.db)
}
