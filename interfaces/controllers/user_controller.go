package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sanshirookazaki/echo-clean/domain"
	"github.com/sanshirookazaki/echo-clean/interfaces/database"
	"github.com/sanshirookazaki/echo-clean/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(SQLHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

func (controller *UserController) Index(c echo.Context) error {
	//u := domain.User{}
	//c.Bind(&u)
	tasks := controller.Interactor.GetTaskAll(1)
	u := domain.User{
		UserID:   1,
		UserName: "test",
		Password: "test",
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", u)
}
