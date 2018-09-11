package services

import (
	"testing"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/repositories"
)

type mockGroupsRepository struct {
	repositories.GroupsRepository
}

// CreateGroup create group
func (r *mockGroupsRepository) CreateGroup(name string) (*orm.Group, error) {
	return &orm.Group{Name: name}, nil
}

// UpdateGroup update group
func (r *mockGroupsRepository) UpdateGroup(group *orm.Group, name string) (*orm.Group, error) {
	group.Name = name
	return group, nil
}

// DeleteGroup update group
func (r *mockGroupsRepository) DeleteGroup(group *orm.Group) error {
	group.Name = "deleted"
	return nil
}

func mockGroupService() GroupService {
	return &groupService{
		groupsRepository: &mockGroupsRepository{},
	}
}

func TestCreateGroup(t *testing.T) {
	name := "test"
	group, _ := mockGroupService().CreateGroup(name)
	if group.Name != name {
		t.Fatalf("group.Name expected %s, actual %s", group.Name, name)
	}
}

func TestUpdateGroup(t *testing.T) {
	name := "test"
	group := &orm.Group{
		Name: "nochange",
	}
	_, _ = mockGroupService().UpdateGroup(group, name)
	if group.Name != name {
		t.Fatalf("group.Name expected %s, actual %s", group.Name, name)
	}
}

func TestDeleteGroup(t *testing.T) {
	group := &orm.Group{
		Name: "nochange",
	}
	_ = mockGroupService().DeleteGroup(group)
	if group.Name != "deleted" {
		t.Fatalf("group.Name expected %s, actual %s", group.Name, "deleted")
	}
}
