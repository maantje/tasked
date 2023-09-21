package controller

import (
	"fmt"
	"net/http"
	"tasked/action"
	"tasked/authentication"
	"tasked/database"
	"tasked/model"

	"github.com/labstack/echo/v4"
)

type TaskController struct{}

func (t *TaskController) Create(c echo.Context) error {
	name := authentication.AuthID(c)
	fmt.Println(name)

	input := &action.CreateTaskInput{
		UserId: authentication.AuthID(c),
	}

	if err := c.Bind(input); err != nil {
		return err
	}

	task, err := action.CreateTask(input)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (t *TaskController) Update(c echo.Context) error {
	input := &action.UpdateTaskInput{
		ID: c.Param("id"),
	}

	if err := c.Bind(input); err != nil {
		return err
	}

	task, err := action.UpdateTask(input)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (t *TaskController) Index(c echo.Context) error {
	var tasks []*model.Task

	if err := database.ORM().Find(&tasks, "user_id = ?", authentication.AuthID(c)).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}

func (t *TaskController) Read(c echo.Context) error {
	var task *model.Task

	err := database.ORM().Where(
		"users_id = ?",
		authentication.AuthID(c),
	).First(&task, c.Param("id")).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (t *TaskController) Delete(c echo.Context) error {
	var task *model.Task

	if err := database.ORM().Where("user_id = ?", authentication.AuthID(c)).First(&task, c.Param("id")).Error; err != nil {
		return err
	}

	if err := database.ORM().Delete(&task).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}
