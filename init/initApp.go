package init

import (
	"auth-service/infrastructure"
	"auth-service/registry"

	"github.com/labstack/echo/v4"
)

func StartApp() {
	conn := Conn()
	c := registry.Registry(conn)
	e := echo.New()
	infrastructure.NewRoute(e, c)
	e.Logger.Fatal(e.Start(":8081"))
}
