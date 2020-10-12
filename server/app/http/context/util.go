package context

import (
	"errors"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/services"
)

// GetLoginUser get login user
func GetLoginUser(c echo.Context) (*orm.User, error) {
	sessionValues, err := session.Get("binsen-session", c)
	if err != nil {
		return nil, err
	}
	userID := sessionValues.Values["id"]
	if userID == nil {
		return nil, errors.New("not authorized")
	}
	userService := services.NewUserService()
	user, err := userService.GetUser(userID.(int64))
	if err != nil {
		return nil, err
	}
	return user, nil
}
