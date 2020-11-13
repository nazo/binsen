package actions

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/services"
)

type getWorkspacesRequest struct {
}

type getWorkspacesResponse struct {
	Workspaces []*orm.Workspace `json:"workspaces"`
}

// GetWorkspaces get workspaces
func GetWorkspaces(c echo.Context) error {
	req := &getWorkspacesRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	workspaceService := services.NewWorkspaceService()
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
	req := &createWorkspacesRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	workspaceService := services.NewWorkspaceService()
	workspace, err := workspaceService.CreateWorkspace(req.Name)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid workspaces")
	}
	res := &createWorkspacesResponse{Workspace: workspace}
	return c.JSON(http.StatusOK, res)
}

type updateWorkspacesRequest struct {
	ID int64 `query:"id"`
	Name string `query:"name"`
}

type updateWorkspacesResponse struct {
	Workspace *orm.Workspace `json:"workspace"`
}

// UpdateWorkspace get workspaces
func UpdateWorkspace(c echo.Context) error {
	req := &updateWorkspacesRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	workspaceService := services.NewWorkspaceService()
	workspace, err := workspaceService.GetWorkspace(req.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid workspace")
	}
	_, err = workspaceService.UpdateWorkspace(workspace, req.Name)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "workspace update failed")
	}
	res := &updateWorkspacesResponse{Workspace: workspace}
	return c.JSON(http.StatusOK, res)
}
