package auth

import (
	"context"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
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
	session := session.Default(c)
	session.Set("google-oauth2-state", state)
	session.Save()
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
	session := session.Default(c)
	if session.Get("google-oauth2-state") != req.State {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid state")
	}
	session.Set("google-oauth2-state", nil)
	token, err := googleService.Exchange(req.Code)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "token validate failed")
	}
	userResponse, err := googleService.GetUser(token)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "userinfo failed")
	}
	user, err := userService.FindUserByGoogleUser(token, userResponse)
	session.Set("user_id", user.ID)
	session.Save()
	res := &getGoogleCallbackResponse{User: user}
	return c.JSON(http.StatusOK, res)
}
