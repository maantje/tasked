package controller

import (
	"net/http"
	"tasked/action"
	"tasked/database"
	"tasked/model"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func (t *UserController) Create(c echo.Context) error {
	input := &action.CreateUserInput{}

	if err := c.Bind(input); err != nil {
		return err
	}

	user, err := action.CreateUser(input)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (t *UserController) Update(c echo.Context) error {
	input := &action.UpdateUserInput{
		ID: c.Param("id"),
	}

	if err := c.Bind(input); err != nil {
		return err
	}

	user, err := action.UpdateUser(input)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (t *UserController) Index(c echo.Context) error {
	var users []*model.User

	if err := database.ORM().Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (t *UserController) Read(c echo.Context) error {
	var user *model.User

	if err := database.ORM().First(&user, c.Param("id")).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (t *UserController) Delete(c echo.Context) error {
	var user *model.User

	if err := database.ORM().First(&user, c.Param("id")).Error; err != nil {
		return err
	}

	if err := database.ORM().Delete(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
