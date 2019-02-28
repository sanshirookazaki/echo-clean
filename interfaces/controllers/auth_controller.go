package controllers

import (
	"html"
	"html/template"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/sanshirookazaki/echo-clean/domain"
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

func Login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/login.html"))
	t.Execute(w, "come on")
}

func (controller *AuthController) Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login", "ログインしてください")
}

func (controller *AuthController) Logout(c echo.Context) error {
	session := session.Default(c)
	session.Clear()
	return c.Render(http.StatusOK, "login", "ログインしてください")
}

func (controller *AuthController) LoginCheck(c echo.Context) error {
	loginform := domain.LoginForm{
		UserName: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	username := html.EscapeString(loginform.UserName)
	password := html.EscapeString(loginform.Password)

	if username == "" && password == "" {
		return c.Render(http.StatusOK, "login", "ユーザー名とパスワードを入力して下さい。")
	} else if username == "" {
		return c.Render(http.StatusOK, "login", "ユーザー名を入力して下さい。")
	} else if password == "" {
		return c.Render(http.StatusOK, "login", "パスワードを入力して下さい。")
	}

	session := session.Default(c)
	hashLoginPassword := controller.Interactor.GetPassword(username)
	err := database.PasswordVerify(hashLoginPassword, password)
	if err != nil {
		userid := controller.Interactor.GetUserID(username, hashLoginPassword)
		session.Set("userid", userid)
		session.Set("password", password)
		session.Save()
		return c.Redirect(http.StatusFound, "/"+username+"/index")
	} else {
		return c.Redirect(http.StatusFound, "login")
	}
}

func (controller *AuthController) LoginNewUser(c echo.Context) error {
	return c.Render(http.StatusFound, "loginnew", "ユーザー名とパスワードを入力して下さい。")
}

func (controller *AuthController) LoginAddUser(c echo.Context) error {
	loginform := domain.LoginForm{
		UserName: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	username := html.EscapeString(loginform.UserName)
	password := html.EscapeString(loginform.Password)
	session := session.Default(c)
	loginname := controller.Interactor.UserUniqueCheck(username)
	if username == loginname {
		return c.Render(http.StatusFound, "loginnew", "このユーザー名は既に使われてます。")
	} else if username == "" || password == "" {
		return c.Render(http.StatusFound, "loginnew", "ユーザー名とパスワードに空白は使えません。")
	} else {
		controller.Interactor.UserAdd(username, password)
		userid := controller.Interactor.GetUserID(username, password)
		session.Set("userid", userid)
		session.Save()
		return c.Redirect(http.StatusFound, "/"+username+"/index")
	}
}
