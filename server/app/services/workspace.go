package services

import (
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/repositories"
)

// WorkspaceService todo
type WorkspaceService interface {
	CreateWorkspace(name string) (*orm.Workspace, error)
	UpdateWorkspace(workspace *orm.Workspace, name string) (*orm.Workspace, error)
	GetWorkspaces() ([]*orm.Workspace, error)
	GetWorkspace(id int64) (*orm.Workspace, error)
}

// workspaceService todo
type workspaceService struct {
	WorkspaceService
	workspacesRepository repositories.WorkspacesRepository
}

// NewWorkspaceService todo
func NewWorkspaceService() WorkspaceService {
	return &workspaceService{
		workspacesRepository: repositories.NewWorkspacesRepository(),
	}
}

// CreateWorkspace create workspace
func (s *workspaceService) CreateWorkspace(name string) (*orm.Workspace, error) {
	return s.workspacesRepository.CreateWorkspace(name)
}

// UpdateWorkspace create workspace
func (s *workspaceService) UpdateWorkspace(workspace *orm.Workspace, name string) (*orm.Workspace, error) {
	return s.workspacesRepository.UpdateWorkspace(workspace, name)
}

// GetWorkspaces get workspaces
func (s *workspaceService) GetWorkspaces() ([]*orm.Workspace, error) {
	return s.workspacesRepository.GetWorkspaces()
}

// GetWorkspace get workspace
func (s *workspaceService) GetWorkspace(id int64) (*orm.Workspace, error) {
	return s.workspacesRepository.GetWorkspace(id)
}
