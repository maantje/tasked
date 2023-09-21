package action

import (
	"net/http"
	"tasked/authentication"
	"tasked/database"
	"tasked/model"

	"github.com/labstack/echo/v4"
)

type CreateUserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func CreateUser(input *CreateUserInput) (*model.User, error) {
	if err := Validate(input); err != nil {
		return &model.User{}, err
	}

	user := &model.User{}

	if !user.IsEmailUnique(input.Email) {
		return user, echo.NewHTTPError(http.StatusUnprocessableEntity, "email already in use")
	}

	hash, err := authentication.HashPassword(input.Password)

	if err != nil {
		return &model.User{}, err
	}

	user = &model.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: hash,
	}

	if err := database.ORM().Create(&user).Error; err != nil {
		return &model.User{}, err
	}

	return user, nil
}
