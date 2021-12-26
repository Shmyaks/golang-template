package internal

import (
	_ "golang-template-module/docs"
	"golang-template-module/internal/delivery/handlers"
	httpv1 "golang-template-module/internal/delivery/http/v1"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunListening() {
	e := echo.New()
	e.Validator = &handlers.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := e.Group("/api/v1")
	{
		v1.POST("/auth", httpv1.RouterAuth)
		v1.POST("/create", httpv1.RouterRegistration)
		user := v1.Group("/user")
		{
			user.GET("/:id", httpv1.RouterUserGetId)
			user.GET("/list", httpv1.RouterUserGetList)
		}
	}
	e.Logger.Fatal(e.Start(":1323"))
}
