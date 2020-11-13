package http

import (
	"os"

	"github.com/facebookgo/grace/gracehttp"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nazo/binsen/server/app/http/routes"
	"github.com/nazo/binsen/server/lib/db"
	"gopkg.in/boj/redistore.v1"
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

	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer store.Close()
	store.Options.HttpOnly = true
	// store.Options.Secure = true;
	// store.Options.SameSite = http.SameSiteNoneMode;
	e.Use(session.Middleware(store))
	conn, err := db.NewDB(&db.NewDBConfig{
		Host:     "postgres",
		Port:     "5432",
		User:     "binsen",
		Password: "binsen",
		Database: "binsen_dev",
		SSLMode:  "disable",
	})
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
