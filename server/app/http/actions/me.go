package actions

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nazo/binsen/server/app/http/context"
	"github.com/nazo/binsen/server/app/orm"
)

type getMeRequest struct {
}

type getMeResponse struct {
	User *orm.User `json:"user"`
}

// GetMe get me
func GetMe(c echo.Context) error {
	user, err := context.GetLoginUser(c)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid user")
	}
	res := &getMeResponse{User: user}
	return c.JSON(http.StatusOK, res)
}
