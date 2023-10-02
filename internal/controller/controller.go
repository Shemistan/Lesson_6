package controller

import (
	"errors"

	"github.com/nasrullonurullaev5/Lesson_6/internal/model"
	"github.com/nasrullonurullaev5/Lesson_6/internal/service"
)

type IUserController interface {
	Authenticate(string, string) (int32, error)
	AddUser(string, string, string, string) (uint32, error)
	GetUser(uint32) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	UpdateUser(uint32, *model.User) error
	DeleteUser(uint32) error
	CollectStatistics() map[string]uint32
}

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Authenticate(login string, password string) (int32, error) {
	if login == "" || password == "" {
		return 0, errors.New("Empty fields not accepted")
	}

	result, err := uc.userService.Authenticate(login, password)

	return result, err
}

func (uc *UserController) AddUser(name, surname, login, password string) (uint32, error) {
	return uc.userService.Register(name, surname, login, password)
}

func (uc *UserController) GetUser(id uint32) (*model.User, error) {
	return uc.userService.Fetch(id)
}

func (uc *UserController) GetAllUsers() ([]*model.User, error) {
	return uc.userService.FetchAll()
}

func (uc *UserController) UpdateUser(id uint32, user *model.User) error {
	return uc.userService.Modify(id, user)
}

func (uc *UserController) DeleteUser(id uint32) error {
	return uc.userService.Delete(id)
}

func (uc *UserController) CollectStatistics() map[string]uint32 {
	return uc.userService.CollectStatistics()
}
