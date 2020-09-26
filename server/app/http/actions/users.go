package actions

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/services"
)

type getUsersRequest struct {
}

type getUsersResponse struct {
	Users []*orm.User `json:"users"`
}

// GetUsers get users
func GetUsers(c echo.Context) error {
	req := &getUsersRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	userService := services.NewUserService()
	users, err := userService.GetUsers()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid users")
	}
	res := &getUsersResponse{Users: users}
	return c.JSON(http.StatusOK, res)
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type createUserResponse struct {
	User *orm.User `json:"user"`
}

// CreateUser create user
func CreateUser(c echo.Context) error {
	req := &createUserRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	group, err := services.NewGroupService().GetDefaultGroup()
	if err != nil {
		panic(err)
	}
	role, err := services.NewRoleService().GetDefaultRole()
	if err != nil {
		panic(err)
	}
	userService := services.NewUserService()
	user, err := userService.CreateUser(req.Name, req.Email, role, group)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "user create failed")
	}
	res := &createUserResponse{User: user}
	return c.JSON(http.StatusOK, res)
}

type updateUserRequest struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type updateUserResponse struct {
	User *orm.User `json:"user"`
}

// UpdateUser update user
func UpdateUser(c echo.Context) error {
	req := &updateUserRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	userService := services.NewUserService()
	user, err := userService.GetUser(req.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid user")
	}
	_, err = userService.UpdateUser(user, req.Name, req.Email)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "user update failed")
	}
	res := &updateUserResponse{User: user}
	return c.JSON(http.StatusOK, res)
}

type deleteUserRequest struct {
	ID int64 `query:"id"`
}

type deleteUserResponse struct {
}

// DeleteUser delete user
func DeleteUser(c echo.Context) error {
	req := &deleteUserRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	userService := services.NewUserService()
	user, err := userService.GetUser(req.ID)
	c.Logger().Error(req.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid user")
	}
	err = userService.DeleteUser(user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "user delete failed")
	}
	res := &deleteUserResponse{}
	return c.JSON(http.StatusOK, res)
}
