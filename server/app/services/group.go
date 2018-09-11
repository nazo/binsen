package services

import (
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/repositories"
)

// GroupService todo
type GroupService interface {
	CreateGroup(name string) (*orm.Group, error)
	UpdateGroup(group *orm.Group, name string) (*orm.Group, error)
	DeleteGroup(group *orm.Group) error
	GetDefaultGroup() (*orm.Group, error)
	GetGroups() ([]*orm.Group, error)
	GetGroup(id int) (*orm.Group, error)
}

// groupService todo
type groupService struct {
	GroupService
	groupsRepository repositories.GroupsRepository
}

// NewGroupService todo
func NewGroupService() GroupService {
	return &groupService{
		groupsRepository: repositories.NewGroupsRepository(),
	}
}

// CreateGroup create group
func (s *groupService) CreateGroup(name string) (*orm.Group, error) {
	return s.groupsRepository.CreateGroup(name)
}

// GetGroups get groups
func (s *groupService) GetGroups() ([]*orm.Group, error) {
	return s.groupsRepository.GetGroups()
}

// GetGroup get group
func (s *groupService) GetGroup(id int) (*orm.Group, error) {
	return s.groupsRepository.GetGroup(id)
}

// GetDefaultGroup get group
func (s *groupService) GetDefaultGroup() (*orm.Group, error) {
	defaultGroupID := 1
	return s.groupsRepository.GetGroup(defaultGroupID)
}

// UpdateGroup update group
func (s *groupService) UpdateGroup(group *orm.Group, name string) (*orm.Group, error) {
	return s.groupsRepository.UpdateGroup(group, name)
}

func (s *groupService) DeleteGroup(group *orm.Group) error {
	return s.groupsRepository.DeleteGroup(group)
}
