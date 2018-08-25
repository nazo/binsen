package auth

import (
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

type deleteSignoutResponse struct {
}

// Signout signout user
func Signout(c echo.Context) error {
	session := session.Default(c)
	session.Delete("user_id")
	session.Save()
	res := &deleteSignoutResponse{}
	return c.JSON(http.StatusOK, res)
}
