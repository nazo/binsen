package services

import (
	"database/sql"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/repositories"
)

// WorkspaceService todo
type WorkspaceService interface {
	CreateWorkspace(name string) (*orm.Workspace, error)
	GetWorkspaces() ([]*orm.Workspace, error)
	GetWorkspace(id int64) (*orm.Workspace, error)
}

// workspaceService todo
type workspaceService struct {
	WorkspaceService
	workspacesRepository repositories.WorkspacesRepository
}

// NewWorkspaceService todo
func NewWorkspaceService(db *sql.DB) WorkspaceService {
	return &workspaceService{
		workspacesRepository: repositories.NewWorkspacesRepository(db),
	}
}

// CreateWorkspace create workspace
func (s *workspaceService) CreateWorkspace(name string) (*orm.Workspace, error) {
	return s.workspacesRepository.CreateWorkspace(name)
}

// GetWorkspaces get workspaces
func (s *workspaceService) GetWorkspaces() ([]*orm.Workspace, error) {
	return s.workspacesRepository.GetWorkspaces()
}

// GetWorkspace get workspace
func (s *workspaceService) GetWorkspace(id int64) (*orm.Workspace, error) {
	return s.workspacesRepository.GetWorkspace(id)
}
