package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (controller *UserController) UserIndex(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	session.Save(r, w)
	userid := session.Values["userid"]
	uid := userid.(int)
	username, password := controller.UserInteractor.GetUserNamePassword(uid)
	tasks := controller.TaskInteractor.GetTaskAll(uid)
	u := domain.User{
		Userid:   uid,
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "index", u)
	json.NewEncoder(w).Encode(u)
}

func (controller *UserController) UserDetailTask(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	userid := session.Values["userid"]
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTask(id)
	u := domain.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "task", u)
	json.NewEncoder(w).Encode(u)
}

func (controller *UserController) UserAddTask(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	userid := session.Values["userid"]
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "add", u)
	json.NewEncoder(w).Encode(u)
}

func (controller *UserController) UserAddTaskPost(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	userid := session.Values["userid"]
	task := r.FormValue("task")
	controller.TaskInteractor.AddTask(userid.(int), task)
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "index", u)
	json.NewEncoder(w).Encode(u)
}

func (controller *UserController) UserDeleteTask(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	userid := session.Values["userid"]
	id, _ := strconv.Atoi(r.FormValue("id"))
	controller.TaskInteractor.DeleteTask(id)
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "index", u)
	json.NewEncoder(w).Encode(u)
}

func (controller *UserController) UserFinishTask(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	userid := session.Values["userid"]
	id, _ := strconv.Atoi(r.FormValue("id"))
	controller.TaskInteractor.FinishTask(id)
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "index", u)
	json.NewEncoder(w).Encode(u)
}

func (controller *UserController) UserTaskHistory(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	userid := session.Values["userid"]
	tasks := controller.TaskInteractor.TaskHistory(userid.(int))
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	u := domain.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "history", u)
	json.NewEncoder(w).Encode(u)
}

func (controller *UserController) UserEditTask(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	userid := session.Values["userid"]
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "edit", u)
	json.NewEncoder(w).Encode(u)
}

func (controller *UserController) UserEditTaskPost(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	userid := session.Values["userid"]
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	task := r.FormValue("task")
	controller.TaskInteractor.EditTask(task, id)
	username, password := controller.UserInteractor.GetUserNamePassword(userid.(int))
	tasks := controller.TaskInteractor.GetTaskAll(userid.(int))
	u := domain.User{
		Userid:   userid.(int),
		Username: username,
		Password: password,
		Tasks:    tasks,
	}
	//T.Render(w, "index", u)
	json.NewEncoder(w).Encode(u)
}
