package repositories

import (
	"context"
	"database/sql"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/lib/db"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// WorkspacesRepository todo
type WorkspacesRepository interface {
	GetWorkspaces() ([]*orm.Workspace, error)
	GetWorkspace(id int64) (*orm.Workspace, error)
	CreateWorkspace(name string) (*orm.Workspace, error)
	UpdateWorkspace(workspace *orm.Workspace, name string) (*orm.Workspace, error)
}

// workspacesRepository todo
type workspacesRepository struct {
	WorkspacesRepository
	db *sql.DB
}

// NewWorkspacesRepository todo
func NewWorkspacesRepository() WorkspacesRepository {
	return &workspacesRepository{
		db: db.InjectDB(),
	}
}

// GetWorkspaces get workspaces
func (r *workspacesRepository) GetWorkspaces() ([]*orm.Workspace, error) {
	return orm.Workspaces(qm.OrderBy("id")).All(context.Background(), r.db)
}

// GetWorkspace get workspace
func (r *workspacesRepository) GetWorkspace(id int64) (*orm.Workspace, error) {
	return orm.Workspaces(qm.Where("id = ?", id)).One(context.Background(), r.db)
}

// CreateWorkspace create workspace
func (r *workspacesRepository) CreateWorkspace(name string) (*orm.Workspace, error) {
	workspace := &orm.Workspace{
		Name: name,
	}
	err := workspace.Insert(context.Background(), r.db, boil.Infer())
	return workspace, err
}

// UpdateWorkspace create workspace
func (r *workspacesRepository) UpdateWorkspace(workspace *orm.Workspace, name string) (*orm.Workspace, error) {
	workspace.Name = name
	_, err := workspace.Update(context.Background(), r.db, boil.Infer())
	return workspace, err
}

