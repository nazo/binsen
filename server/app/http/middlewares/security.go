package middlewares

import (
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// Security user auth middleware
func Security() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session := session.Default(c)
			user := session.Get("user_id")
			if user == nil {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
