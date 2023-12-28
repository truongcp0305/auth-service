package registry

import (
	"auth-service/controller"
	"auth-service/repository"
	"auth-service/service"

	"github.com/go-pg/pg/v10"
)

func Registry(pg *pg.DB) controller.AuthController {
	db := repository.NewDatabase(pg)
	return *controller.NewAuthController(
		*service.NewAccountService(db),
	)
}
