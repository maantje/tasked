package action

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func Validate(i interface{}) error {
	if err := validator.New().Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return nil
}
