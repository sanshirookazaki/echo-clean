package controllers

import (
	"html"
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
	loginpass := controller.Interactor.GetPassword(username)
	if password == loginpass {
		userid := controller.Interactor.GetUserID(username, password)
		//tasks := models.GetTaskAll(userid)
		session.Set("userid", userid)
		session.Save()
		return c.Redirect(http.StatusFound, "/"+username+"/index")
	} else {
		return c.Redirect(http.StatusFound, "login")
	}
}
