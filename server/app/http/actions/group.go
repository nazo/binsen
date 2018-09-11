package actions

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/services"
)

type getGroupsRequest struct {
}

type getGroupsResponse struct {
	Groups []*orm.Group `json:"groups"`
}

// GetGroups get groups
func GetGroups(c echo.Context) error {
	req := &getGroupsRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	groupService := services.NewGroupService()
	groups, err := groupService.GetGroups()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid groups")
	}
	res := &getGroupsResponse{Groups: groups}
	return c.JSON(http.StatusOK, res)
}

type createGroupRequest struct {
	Name string `json:"name"`
}

type createGroupResponse struct {
	Group *orm.Group `json:"group"`
}

// CreateGroup create group
func CreateGroup(c echo.Context) error {
	req := &createGroupRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	groupService := services.NewGroupService()
	group, err := groupService.CreateGroup(req.Name)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "group create failed")
	}
	res := &createGroupResponse{Group: group}
	return c.JSON(http.StatusOK, res)
}

type updateGroupRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type updateGroupResponse struct {
	Group *orm.Group `json:"group"`
}

// UpdateGroup update group
func UpdateGroup(c echo.Context) error {
	req := &updateGroupRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	groupService := services.NewGroupService()
	group, err := groupService.GetGroup(req.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid group")
	}
	_, err = groupService.UpdateGroup(group, req.Name)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "group update failed")
	}
	res := &updateGroupResponse{Group: group}
	return c.JSON(http.StatusOK, res)
}

type deleteGroupRequest struct {
	ID int `query:"id"`
}

type deleteGroupResponse struct {
}

// DeleteGroup delete group
func DeleteGroup(c echo.Context) error {
	req := &deleteGroupRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	groupService := services.NewGroupService()
	group, err := groupService.GetGroup(req.ID)
	c.Logger().Error(req.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid group")
	}
	err = groupService.DeleteGroup(group)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "group delete failed")
	}
	res := &deleteGroupResponse{}
	return c.JSON(http.StatusOK, res)
}
