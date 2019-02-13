package controllers

import (
	"net/http"
	"strconv"

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
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", u)
}

func (controller *UserController) UserDetailTask(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	id, _ := strconv.Atoi(c.Param("id"))
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTask(id)
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "task", u)
}

func (controller *UserController) UserAddTask(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "add", u)
}

func (controller *UserController) UserAddTaskPost(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	task := c.FormValue("task")
	controller.TaskInteractor.AddTask(userid.(int), task)
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", u)
}

func (controller *UserController) UserDeleteTask(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	id, _ := strconv.Atoi(c.FormValue("id"))
	controller.TaskInteractor.DeleteTask(id)
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", u)
}

func (controller *UserController) UserFinishTask(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	id, _ := strconv.Atoi(c.FormValue("id"))
	controller.TaskInteractor.FinishTask(id)
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", u)
}

func (controller *UserController) UserTaskHistory(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	tasks := controller.TaskInteractor.TaskHistory(userid.(int))
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "history", u)
}

func (controller *UserController) UserEditTask(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "edit", u)
}

func (controller *UserController) UserEditTaskPost(c echo.Context) error {
	session := session.Default(c)
	userid := session.Get("userid")
	id, _ := strconv.Atoi(c.FormValue("id"))
	task := c.FormValue("task")
	controller.TaskInteractor.EditTask(task, id)
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		UserID:   userid.(int),
		UserName: username,
		Password: password,
		Tasks:    tasks,
	}
	return c.Render(http.StatusOK, "index", u)
}
