package controllers

import (
	"go-clean-architecture/domain"
	"go-clean-architecture/interfaces/database"
	"go-clean-architecture/usecase"
	"strconv"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) error {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return err
	}
	c.JSON(201, user)
	return nil
}

func (controller *UserController) Index(c Context) error {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return err
	}
	c.JSON(200, users)
	return nil
}

func (controller *UserController) Show(c Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return err
	}
	c.JSON(200, user)
	return nil
}
