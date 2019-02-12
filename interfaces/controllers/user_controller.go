package controllers

import (
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/sanshirookazaki/echo-clean/domain"
	"github.com/sanshirookazaki/echo-clean/interfaces/database"
	"github.com/sanshirookazaki/echo-clean/usecase"
)

type UserController struct {
	UserInteractor usecase.UserInteractor
	TaskInteractor usecase.TaskInteractor
}

func NewUserController(SQLHandler database.SQLHandler) *UserController {
	return &UserController{
		UserInteractor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: SQLHandler,
			},
		},
		TaskInteractor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}
func (controller *UserController) UserIndex(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", u)
}
