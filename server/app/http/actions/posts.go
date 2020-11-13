package actions

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nazo/binsen/server/app/http/context"
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/services"
	"github.com/pkg/errors"
)

type getPostsRequest struct {
	WorkspaceID int64 `query:"workspace_id" validate:"required"`
	Page        int   `query:"page" validate:"required"`
}

type getPostsResponse struct {
	Posts []*orm.Post `json:"posts"`
}

// GetPosts get posts
func GetPosts(c echo.Context) error {
	req := &getPostsRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	workspaceService := services.NewWorkspaceService()
	workspace, err := workspaceService.GetWorkspace(req.WorkspaceID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid workspace")
	}
	postService := services.NewPostService()
	posts, err := postService.GetPosts(workspace, req.Page)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid posts")
	}
	res := &getPostsResponse{Posts: posts}
	return c.JSON(http.StatusOK, res)
}

type getPostRequest struct {
	ID int64 `query:"id" validate:"required"`
}

type getPostResponse struct {
	Post *orm.Post `json:"post"`
}

// GetPost get post
func GetPost(c echo.Context) error {
	req := &getPostRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	postService := services.NewPostService()
	post, err := postService.GetPost(req.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid posts")
	}
	res := &getPostResponse{Post: post}
	return c.JSON(http.StatusOK, res)
}

type createPostsRequest struct {
	WorkspaceID int64  `json:"workspace_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Body        string `json:"body" validate:"required"`
}

type createPostsResponse struct {
	Post *orm.Post `json:"post"`
}

// CreatePost get posts
func CreatePost(c echo.Context) error {
	req := &createPostsRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	user, err := context.GetLoginUser(c)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid user")
	}
	workspaceService := services.NewWorkspaceService()
	workspace, err := workspaceService.GetWorkspace(req.WorkspaceID)
	if err != nil {
		c.Logger().Error(errors.Wrap(err, fmt.Sprintf("workspaceService.GetWorkspace ID:[%d]", req.WorkspaceID)))
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid workspace")
	}
	postService := services.NewPostService()
	post, err := postService.CreatePost(workspace, req.Title, req.Body, user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid posts")
	}
	res := &createPostsResponse{Post: post}
	return c.JSON(http.StatusOK, res)
}

type updatePostsRequest struct {
	ID    int64  `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

type updatePostsResponse struct {
	Post *orm.Post `json:"post"`
}

// UpdatePost get posts
func UpdatePost(c echo.Context) error {
	req := &updatePostsRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	if err := c.Validate(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	user, err := context.GetLoginUser(c)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid user")
	}
	postService := services.NewPostService()
	post, err := postService.GetPost(req.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid post ID")
	}
	post, err = postService.UpdatePost(post, req.Title, req.Body, user)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid posts")
	}
	res := &updatePostsResponse{Post: post}
	return c.JSON(http.StatusOK, res)
}
