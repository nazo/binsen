package context

import (
	"errors"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/services"
)

// GetLoginUser get login user
func GetLoginUser(c echo.Context) (*orm.User, error) {
	session := session.Default(c)
	userID := session.Get("user_id")
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
