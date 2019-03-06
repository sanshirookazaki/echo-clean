package infrastructure

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sanshirookazaki/echo-clean/interfaces/controllers"
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

func Init() {
	t := NewTemplate("views/*.html")

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

	userController := controllers.NewUserController(NewSQLHandler())
	authController := controllers.NewAuthController(NewSQLHandler())
	//store := session.NewCookieStore([]byte("secret-key"))
	//store.MaxAge(86400)
	//e.Use(session.Sessions("ESESSION", store))
	e.Static("/static", "static")
	//e.GET("/", authController.Login)
	//e.GET("/login", authController.Login)
	e.POST("/login", authController.LoginCheck)
	e.POST("/logout", authController.Logout)
	e.GET("/login/new", authController.LoginNewUser)
	e.POST("/login/new", authController.LoginAddUser)
	e.GET("/:username/index", userController.UserIndex)
	e.GET("/:username/task/:id", userController.UserDetailTask)
	e.GET("/:username/task/add", userController.UserAddTask)
	e.POST("/:username/task/add", userController.UserAddTaskPost)
	e.POST("/:username/task/delete", userController.UserDeleteTask)
	e.GET("/:username/task/:id/edit", userController.UserEditTask)
	e.POST("/:username/task/:id/edit", userController.UserEditTaskPost)
	e.POST("/:username/task/finish", userController.UserFinishTask)
	e.GET("/:username/task/history", userController.UserTaskHistory)
	//e.Logger.Fatal(e.Start(":1323"))

	n := negroni.Classic()
	//store := cookiestore.New([]byte("secret123"))
	store := sessions.NewCookieStore([]byte("secret"))
	session := sessions.NewSession(store, "secret")
	fmt.Println(session)
	r := mux.NewRouter()
	r.HandleFunc("/login", authController.Login).Methods("GET")

	//srv := &http.Server{
	//	Handler: r,
	//	Addr:    "127.0.0.1:8000",
	//	// Good practice: enforce timeouts for servers you create!
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}

	//log.Fatal(srv.ListenAndServe())

	n.UseHandler(r)
	n.Run(":3000")
}
