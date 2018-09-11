package repositories

import (
	"context"
	"database/sql"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/lib/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GroupsRepository todo
type GroupsRepository interface {
	GetGroups() ([]*orm.Group, error)
	GetGroup(id int) (*orm.Group, error)
	CreateGroup(name string) (*orm.Group, error)
	UpdateGroup(group *orm.Group, name string) (*orm.Group, error)
	DeleteGroup(group *orm.Group) error
}

// groupsRepository todo
type groupsRepository struct {
	GroupsRepository
	db *sql.DB
}

// NewGroupsRepository todo
func NewGroupsRepository() GroupsRepository {
	return &groupsRepository{
		db: db.InjectDB(),
	}
}

// GetGroups get groups
func (r *groupsRepository) GetGroups() ([]*orm.Group, error) {
	return orm.Groups(qm.OrderBy("id")).All(context.Background(), r.db)
}

// GetGroup get group
func (r *groupsRepository) GetGroup(id int) (*orm.Group, error) {
	return orm.Groups(qm.Where("id = ?", id)).One(context.Background(), r.db)
}

// CreateGroup create group
func (r *groupsRepository) CreateGroup(name string) (*orm.Group, error) {
	group := &orm.Group{
		Name: name,
	}
	err := group.Insert(context.Background(), r.db, boil.Infer())
	return group, err
}

// UpdateGroup create group
func (r *groupsRepository) UpdateGroup(group *orm.Group, name string) (*orm.Group, error) {
	group.Name = name
	_, err := group.Update(context.Background(), r.db, boil.Infer())
	return group, err
}

// DeleteGroup delete group
func (r *groupsRepository) DeleteGroup(group *orm.Group) error {
	_, err := group.Delete(context.Background(), r.db)
	return err
}
