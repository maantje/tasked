package action

import (
	"net/http"
	"os"
	"tasked/authentication"
	"tasked/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(input LoginInput) (*authentication.Token, error) {
	user := &model.User{}

	if !user.Authenticate(input.Email, input.Password) {
		return &authentication.Token{}, echo.NewHTTPError(http.StatusUnauthorized, "Invalid email and/or password")
	}

	exp := time.Now().Add(time.Hour * 1)

	claims := authentication.Claims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    "tasked",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return &authentication.Token{AccessToken: token, ExpiresAt: exp}, err
}
