package actions

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/services"
	"github.com/nazo/binsen/server/lib/db"
)

type getWorkspacesRequest struct {
}

type getWorkspacesResponse struct {
	Workspaces []*orm.Workspace `json:"workspaces"`
}

// GetWorkspaces get workspaces
func GetWorkspaces(c echo.Context) error {
	db := db.Default(c)
	req := &getWorkspacesRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	workspaceService := services.NewWorkspaceService(db)
	workspaces, err := workspaceService.GetWorkspaces()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid workspaces")
	}
	res := &getWorkspacesResponse{Workspaces: workspaces}
	return c.JSON(http.StatusOK, res)
}

type createWorkspacesRequest struct {
	Name string `query:"name"`
}

type createWorkspacesResponse struct {
	Workspace *orm.Workspace `json:"workspace"`
}

// CreateWorkspace get workspaces
func CreateWorkspace(c echo.Context) error {
	db := db.Default(c)
	req := &createWorkspacesRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	workspaceService := services.NewWorkspaceService(db)
	workspace, err := workspaceService.CreateWorkspace(req.Name)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid workspaces")
	}
	res := &createWorkspacesResponse{Workspace: workspace}
	return c.JSON(http.StatusOK, res)
}
