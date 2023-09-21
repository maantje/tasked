package action

import (
	"net/http"
	"tasked/authentication"
	"tasked/database"
	"tasked/model"

	"github.com/labstack/echo/v4"
)

type UpdateUserInput struct {
	ID       string
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func UpdateUser(input *UpdateUserInput) (*model.User, error) {
	if err := Validate(input); err != nil {
		return &model.User{}, err
	}

	user := &model.User{}

	if err := database.ORM().First(&user, input.ID).Error; err != nil {
		return user, err
	}

	if !user.IsEmailUnique(input.Email) {
		return user, echo.NewHTTPError(http.StatusUnprocessableEntity, "email already in use")
	}

	hash, err := authentication.HashPassword(input.Password)

	if err != nil {
		return user, err
	}

	err = database.ORM().Model(&user).Updates(&model.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: hash,
	}).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
