package infrastructure

import (
	"auth-service/controller"

	"github.com/labstack/echo/v4"
)

func NewRoute(c *echo.Echo, u controller.AuthController) {
	c.GET("get-token", func(c echo.Context) error {
		return nil
	})

	c.POST("/user", func(c echo.Context) error {
		return u.CreateAccount(c)
	})

	c.POST("/user/login", func(c echo.Context) error {
		return u.Login(c)
	})
}
