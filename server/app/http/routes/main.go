package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/nazo/binsen/server/app/http/actions"
	"github.com/nazo/binsen/server/app/http/actions/auth"
	"github.com/nazo/binsen/server/app/http/middlewares"
)

// Main router
func Main(e *echo.Echo) {
	routeAuth := e.Group("/api/auth/v1")
	routeAuth.GET("/google/auth", auth.GetGoogleAuth)
	routeAuth.GET("/google/callback", auth.GetGoogleCallback)
	routeAuth.DELETE("/signout", auth.Signout)

	routeWorkspace := e.Group("/api/workspace/v1")
	routeWorkspace.Use(middlewares.Security())
	routeWorkspace.GET("/list", actions.GetWorkspaces)
	routeWorkspace.PUT("/create", actions.CreateWorkspace)

	routePost := e.Group("/api/post/v1")
	routePost.Use(middlewares.Security())
	routePost.GET("/list", actions.GetPosts)
	routePost.PUT("/create", actions.CreatePost)
	routePost.PATCH("/update", actions.UpdatePost)
	routePost.GET("/get", actions.GetPost)

	routeMe := e.Group("/api/me/v1")
	routeMe.Use(middlewares.Security())
	routeMe.GET("/get", actions.GetMe)

	routeUser := e.Group("/api/user/v1")
	routeUser.Use(middlewares.Security())
	routeUser.GET("/list", actions.GetUsers)
	routeUser.PUT("/create", actions.CreateUser)
	routeUser.PATCH("/update", actions.UpdateUser)
	routeUser.DELETE("/delete", actions.DeleteUser)

	routeGroup := e.Group("/api/group/v1")
	routeGroup.Use(middlewares.Security())
	routeGroup.GET("/list", actions.GetGroups)
	routeGroup.PUT("/create", actions.CreateGroup)
	routeGroup.PATCH("/update", actions.UpdateGroup)
	routeGroup.DELETE("/delete", actions.DeleteGroup)
}
