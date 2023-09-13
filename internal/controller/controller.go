package controller

import (
	"errors"
	"github.com/Shemistan/Lesson_6/internal/model"
	"github.com/Shemistan/Lesson_6/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Auth(login string, password string) (int32, error) {

	if login == "" || password == "" {
		return 0, errors.New("Empty fields not excepted")
	}

	result, err := uc.userService.Auth(login, password)

	return result, err
}

func (uc *UserController) Add(name, surname, login, password string) (uint32, error) {
	return uc.userService.Add(name, surname, login, password)
}

func (uc *UserController) Get(id uint32) (*model.User, error) {
	return uc.userService.Get(id)
}

func (uc *UserController) GetAll() ([]*model.User, error) {
	return uc.userService.GetAll()
}

func (uc *UserController) Update(id uint32, user *model.User) error {
	return uc.userService.Update(id, user)
}

func (uc *UserController) Delete(id uint32) error {
	return uc.userService.Delete(id)
}

func (uc *UserController) GetStatistics() map[string]uint32 {
	return uc.userService.GetStatistics()
}
