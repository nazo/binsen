package http

import (
	"os"

	"github.com/facebookgo/grace/gracehttp"

	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nazo/binsen/server/app/http/routes"
	"github.com/nazo/binsen/server/lib/db"
	"gopkg.in/go-playground/validator.v9"
)

// Server run server
func Server() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("API_CLIENT_ORIGIN")},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	store, err := session.NewRedisStore(32, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	e.Use(session.Sessions("binsen_session", store))
	conn, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	e.Use(db.DatabaseMiddleware(conn))
	e.Validator = &CustomValidator{validator: validator.New()}

	// Routes
	routes.Main(e)

	e.Server.Addr = ":5000"

	// Start server
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
