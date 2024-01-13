package main

import (
	"errors"
	"net/http"
	"tasked/controller"
	"tasked/database"
	"tasked/database/migration"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	database.Init()
	migration.Migrate()

	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "not found",
			})

			return
		}

		e.DefaultHTTPErrorHandler(err, c)
	}

	// e.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	// 	Skipper: func(c echo.Context) bool {
	// 		return c.Request().URL.String() == "/login" || (c.Request().URL.String() == "/users" && c.Request().Method == "POST")
	// 	},
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(authentication.Claims)
	// 	},
	// }))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tc := &controller.TaskController{}

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})

	e.GET("/tasks", tc.Index)
	e.GET("/tasks/:id", tc.Read)
	e.PATCH("/tasks/:id", tc.Update)
	e.DELETE("/tasks/:id", tc.Delete)
	e.POST("/tasks", tc.Create)

	uc := &controller.UserController{}

	e.GET("/users", uc.Index)
	e.GET("/users/:id", uc.Read)
	e.PATCH("/users/:id", uc.Update)
	e.DELETE("/users/:id", uc.Delete)
	e.POST("/users", uc.Create)

	ac := &controller.AuthController{}

	e.POST("/login", ac.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
