package controllers

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sanshirookazaki/echo-clean/domain"
	"github.com/sanshirookazaki/echo-clean/interfaces/database"
	"github.com/sanshirookazaki/echo-clean/template"
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

var (
	T     = template.NewTemplate("views/*.html")
	Store = sessions.NewCookieStore([]byte("SESSION_KEY"))
)

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	session.Values["foo"] = "bar"
	fmt.Println(session.Values["foo"])
	session.Save(r, w)
	T.Render(w, "login", "commeon")
}

func (controller *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	fmt.Println(session.Values["foo"])
	T.Render(w, "login", "commeon")
}

func (controller *AuthController) LoginCheck(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	loginform := domain.LoginForm{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	username := html.EscapeString(loginform.UserName)
	password := html.EscapeString(loginform.Password)

	if username == "" && password == "" {
		T.Render(w, "login", "ユーザー名とパスワードを入力して下さい。")
	} else if username == "" {
		T.Render(w, "login", "ユーザー名を入力して下さい。")
	} else if password == "" {
		T.Render(w, "login", "パスワードを入力して下さい。")
	}

	hashLoginPassword := controller.Interactor.GetPassword(username)
	err := database.PasswordVerify(hashLoginPassword, password)
	if err != nil {
		userid := controller.Interactor.GetUserID(username, hashLoginPassword)
		session.Values["userid"] = userid
		session.Values["password"] = password
		session.Save(r, w)
		http.Redirect(w, r, "/"+username+"/index", 301)
	} else {
		http.Redirect(w, r, "/login", 301)
	}
}

func (controller *AuthController) LoginNewUser(w http.ResponseWriter, r *http.Request) {
	T.Render(w, "loginnew", "ユーザー名とパスワードを入力して下さい。")
}

func (controller *AuthController) LoginAddUser(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "SESSION_KEY")
	loginform := domain.LoginForm{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	username := html.EscapeString(loginform.UserName)
	password := html.EscapeString(loginform.Password)
	loginname := controller.Interactor.UserUniqueCheck(username)
	if username == loginname {
		T.Render(w, "loginnew", "このユーザー名は既に使われてます。")
	} else if username == "" || password == "" {
		T.Render(w, "loginnew", "ユーザー名とパスワードに空白は使えません。")
	} else {
		controller.Interactor.UserAdd(username, password)
		userid := controller.Interactor.GetUserID(username, password)
		session.Values["userid"] = userid
		session.Save(r, w)
		http.Redirect(w, r, "/"+username+"/index", 301)
	}
}
