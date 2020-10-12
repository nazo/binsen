package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/services"
	authServices "github.com/nazo/binsen/server/app/services/auth"
)

type getGoogleAuthResponse struct {
	RedirectURI string `json:"redirectUri"`
}

// GetGoogleAuth begin google auth flow
func GetGoogleAuth(c echo.Context) error {
	googleService := authServices.NewGoogleService(context.Background())
	url, state := googleService.GetDefaultAuthCodeURL()
	sessionValues, err := session.Get("binsen-session", c)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid state 1")
	}
	sessionValues.Values["state"] = state
	err = sessionValues.Save(c.Request(), c.Response())
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid state 2")
	}
	res := &getGoogleAuthResponse{
		RedirectURI: url,
	}
	return c.JSON(http.StatusOK, res)
}

type getGoogleCallbackRequest struct {
	Code  string `query:"code"`
	State string `query:"state"`
}

type getGoogleCallbackResponse struct {
	User *orm.User `json:"user"`
}

// GetGoogleCallback receive google auth callback action
func GetGoogleCallback(c echo.Context) error {
	googleService := authServices.NewGoogleService(context.Background())
	userService := services.NewUserService()
	req := &getGoogleCallbackRequest{}
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "missing required parameters")
	}
	sessionValues, err := session.Get("binsen-session", c)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid state 1")
	}
	if req.State != sessionValues.Values["state"] {
		c.Logger().Error("invalid state")
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid state 2")
	}
	sessionValues.Values["state"] = nil
	token, err := googleService.Exchange(req.Code)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "token validate failed")
	}
	userResponse, err := googleService.GetUser(token)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "userinfo failed 1")
	}
	user, err := userService.FindUserByGoogleUser(token, userResponse)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "userinfo failed 2")
	}
	sessionValues.Values["id"] = user.ID
	err = sessionValues.Save(c.Request(), c.Response())
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "userinfo failed 4")
	}
	res := &getGoogleCallbackResponse{User: user}
	return c.JSON(http.StatusOK, res)
}
