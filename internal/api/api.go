package api

import (
	"golang-api/internal/converters"
	"golang-api/internal/models"
	"golang-api/internal/service"
)

type IApi interface {
	Register(req *models.Request) (int64, error)
	Update(Id int64, req *models.UserUpdateRequest) error
	Get(Id int64) (*models.User, error)
	GetAllUsers() []*models.User
	DeleteUser(Id int64) (int64, error)
	GetStats() map[string]int
}

func New(serv service.IService) IApi {
	return &api{
		serv: serv,
	}
}

type api struct {
	serv service.IService
}

func (a *api) Register(req *models.Request) (int64, error) {
	id, err := a.serv.Register(converters.ApiModelToServiceModel(*req))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *api) Update(id int64, req *models.UserUpdateRequest) error {
	err := a.serv.UpdateUser(id, converters.UserUpdate(*req))
	if err != nil {
		return err
	}

	return nil
}

func (a *api) Get(id int64) (*models.User, error) {
	user, err := a.serv.GetUser(id)
	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (a *api) GetAllUsers() []*models.User {
	usersList := a.serv.GetAllUsers()

	return usersList
}

func (a *api) DeleteUser(id int64) (int64, error) {
	id, err := a.serv.DeleteUser(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *api) GetStats() map[string]int {
	stats := a.serv.GetStats()
	return stats
}
