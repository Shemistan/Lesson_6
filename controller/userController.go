package controller

import (
	"errors"
	"fmt"
	"github.com/Shemistan/Lesson_6/model"
	"github.com/Shemistan/Lesson_6/service"
)

func cin[T any](msg string, value T) {
	fmt.Print(msg)
	fmt.Scan(&value)
}

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Auth(login string, password string) (uint32, error) {

	if login == "" || password == "" {
		return 0, errors.New("Empty fields not excepted")
	}

	result, _ := uc.userService.Auth(login, password)

	if result == -1 {
		var name, surname string
		cin("Введите имя: ", name)
		cin("Введите фамилию: ", surname)

		uc.userService.Add(name, surname, login, password)
	}

	return uint32(result), nil
}

func (uc *UserController) Get(id uint32) (model.User, error) {
	return uc.userService.Get(id)
}

func (uc *UserController) GetAll() ([]model.User, error) {
	return uc.userService.GetAll()
}

func (uc *UserController) Update(id uint32, user model.User) error {
	return uc.userService.Update(id, user)
}

func (uc *UserController) Delete(id uint32) error {
	return uc.userService.Delete(id)
}
