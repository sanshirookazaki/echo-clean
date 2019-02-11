package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sanshirookazaki/echo-clean/interfaces/database"
	"github.com/sanshirookazaki/echo-clean/usecase"
)

type AuthController struct {
	Interactor usecase.AuthInteractor
}

func NewAuthController(SQLHandler database.SQLHandler) *AuthController {
	return &AuthController{
		Interactor: usecase.AuthInteractor{
			AuthRepository: &database.AuthRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

func (controller *AuthController) Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login", "ログインしてください")
}
