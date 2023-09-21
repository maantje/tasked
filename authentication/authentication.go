package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthID(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims)
	return claims.UserId
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
