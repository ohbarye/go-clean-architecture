package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-clean-architecture/interfaces/controllers"
)

func Run()  {
	e := echo.New()

	userController := controllers.NewUserController(NewSqlHandler())

	e.Use(middleware.Logger())

	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })

	e.Logger.Fatal(e.Start(":1313"))
}
