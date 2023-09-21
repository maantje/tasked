package controller

import (
	"net/http"
	"tasked/action"

	"github.com/labstack/echo/v4"
)

type AuthController struct{}

func (a *AuthController) Login(c echo.Context) error {
	input := &action.LoginInput{}

	if err := c.Bind(input); err != nil {
		return err
	}

	token, err := action.Login(*input)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}
