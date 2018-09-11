package repositories

import (
	"context"
	"database/sql"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/lib/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// RolesRepository todo
type RolesRepository interface {
	GetRoles() ([]*orm.Role, error)
	GetRole(id int) (*orm.Role, error)
	CreateRole(name string) (*orm.Role, error)
}

// rolesRepository todo
type rolesRepository struct {
	RolesRepository
	db *sql.DB
}

// NewRolesRepository todo
func NewRolesRepository() RolesRepository {
	return &rolesRepository{
		db: db.InjectDB(),
	}
}

// GetRoles get roles
func (r *rolesRepository) GetRoles() ([]*orm.Role, error) {
	return orm.Roles(qm.OrderBy("id")).All(context.Background(), r.db)
}

// GetRole get role
func (r *rolesRepository) GetRole(id int) (*orm.Role, error) {
	return orm.Roles(qm.Where("id = ?", id)).One(context.Background(), r.db)
}

// CreateRole create role
func (r *rolesRepository) CreateRole(name string) (*orm.Role, error) {
	role := &orm.Role{
		Name: name,
	}
	err := role.Insert(context.Background(), r.db, boil.Infer())
	return role, err
}
