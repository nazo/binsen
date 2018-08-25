package db

import (
	"database/sql"

	"github.com/labstack/echo"
)

const (
	// DefaultKey todo
	DefaultKey = "db_connection_key"
)

// Default todo
func Default(ctx echo.Context) *sql.DB {
	db := ctx.Get(DefaultKey)
	if db == nil {
		return nil
	}
	return ctx.Get(DefaultKey).(*sql.DB)
}

// DatabaseMiddleware todo
func DatabaseMiddleware(db *sql.DB) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set(DefaultKey, db)
			return h(ctx)
		}
	}
}
