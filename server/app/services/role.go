package services

import (
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/repositories"
)

// RoleService todo
type RoleService interface {
	CreateRole(name string) (*orm.Role, error)
	GetRoles() ([]*orm.Role, error)
	GetRole(id int) (*orm.Role, error)
	GetDefaultRole() (*orm.Role, error)
}

// roleService todo
type roleService struct {
	RoleService
	rolesRepository repositories.RolesRepository
}

// NewRoleService todo
func NewRoleService() RoleService {
	return &roleService{
		rolesRepository: repositories.NewRolesRepository(),
	}
}

// CreateRole create role
func (s *roleService) CreateRole(name string) (*orm.Role, error) {
	return s.rolesRepository.CreateRole(name)
}

// GetRoles get roles
func (s *roleService) GetRoles() ([]*orm.Role, error) {
	return s.rolesRepository.GetRoles()
}

// GetRole get role
func (s *roleService) GetRole(id int) (*orm.Role, error) {
	return s.rolesRepository.GetRole(id)
}

// GetDefaultRole get role
func (s *roleService) GetDefaultRole() (*orm.Role, error) {
	defaultRoleID := 1
	return s.rolesRepository.GetRole(defaultRoleID)
}
