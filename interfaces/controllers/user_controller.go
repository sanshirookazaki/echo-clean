package controllers

import (
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

