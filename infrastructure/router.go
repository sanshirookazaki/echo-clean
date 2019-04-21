package infrastructure

import (
	"net/http"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/sanshirookazaki/echo-clean/interfaces/controllers"
)

func Init() {
	userController := controllers.NewUserController(NewSQLHandler())
	authController := controllers.NewAuthController(NewSQLHandler())
	n := negroni.Classic()
	r := mux.NewRouter()
	r.HandleFunc("/login", authController.Login).Methods("GET")
	r.HandleFunc("/login", authController.LoginCheck).Methods("POST")
	r.HandleFunc("/logout", authController.Logout).Methods("POST")
	r.HandleFunc("/{username}/index", userController.UserIndex).Methods("GET")
	r.HandleFunc("/login/new", authController.LoginNewUser).Methods("GET")
	r.HandleFunc("/login/new", authController.LoginAddUser).Methods("POST")
	r.HandleFunc("/{username}/task/detail/{id}", userController.UserDetailTask).Methods("GET")
	r.HandleFunc("/{username}/task/add", userController.UserAddTask).Methods("GET")
	r.HandleFunc("/{username}/task/add", userController.UserAddTaskPost).Methods("POST")
	r.HandleFunc("/{username}/task/delete", userController.UserDeleteTask).Methods("POST")
	r.HandleFunc("/{username}/task/{id}/edit", userController.UserEditTask).Methods("GET")
	r.HandleFunc("/{username}/task/{id}/edit", userController.UserEditTaskPost).Methods("POST")
	r.HandleFunc("/{username}/task/finish", userController.UserFinishTask).Methods("POST")
	r.HandleFunc("/{username}/task/history", userController.UserTaskHistory).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./vue"))))
	n.UseHandler(r)
	n.Run(":3000")
}
