package main

import (
	"html/template"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sanshirookazaki/echo-clean/infrastructure"
	"github.com/sanshirookazaki/echo-clean/interfaces/controllers"
	"github.com/sanshirookazaki/echo-mvc/util"
)

type Template struct {
	templates *template.Template
}

func NewTemplate(path string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	t := util.NewTemplate("views/*.html")

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Renderer = t

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.Render(http.StatusNotFound, "404", nil)
			} else {
				c.Render(http.StatusInternalServerError, "500", nil)
			}
		}
	}
	userController := controllers.NewUserController(infrastructure.NewSQLHandler())
	authController := controllers.NewAuthController(infrastructure.NewSQLHandler())
	e.Static("/static", "static")
	e.GET("/", userController.Index)
	e.GET("/login", authController.Login)
	e.Logger.Fatal(e.Start(":1323"))
}
