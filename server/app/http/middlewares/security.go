package middlewares

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Security user auth middleware
func Security() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := session.Get("user_id", c)
			if err != nil {
				return echo.ErrUnauthorized
			}
			if user == nil {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
