package controller

import (
	"auth-service/auth"
	"auth-service/service"
	"auth-service/view/incoming"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	userService service.AccountService
}

func NewAuthController(s service.AccountService) *AuthController {
	return &AuthController{
		userService: s,
	}
}

func (a *AuthController) GetToken(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (a *AuthController) Login(c echo.Context) error {
	var params incoming.LoginParam
	c.Bind(&params)
	if params.UserName == "" || params.Password == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := params.GetModel()
	result, err := a.userService.Login(user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}
	result.Token, err = auth.CreateJwt(*user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (a *AuthController) CreateAccount(c echo.Context) error {
	var params incoming.CreateUserParam
	c.Bind(&params)
	if params.UserName == "" || params.Password == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := params.GetModel()
	result, err := a.userService.CreateAccount(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result.Token, err = auth.CreateJwt(*user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
