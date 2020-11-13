package middlewares

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Security user auth middleware
func Security() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionValues, err := session.Get("binsen-session", c)
			if err != nil {
				c.Logger().Error(err)
				return echo.ErrUnauthorized
			}
			user := sessionValues.Values["id"]
			if user == nil {
				c.Logger().Error(err)
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
