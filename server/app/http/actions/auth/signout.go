package auth

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type deleteSignoutResponse struct {
}

// Signout signout user
func Signout(c echo.Context) error {
	userSessionValues, err := session.Get("user", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "userinfo failed")
	}
	userSessionValues.Values["id"] = nil
	userSessionValues.Save(c.Request(), c.Response())
	res := &deleteSignoutResponse{}
	return c.JSON(http.StatusOK, res)
}
