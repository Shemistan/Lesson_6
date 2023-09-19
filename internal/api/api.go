package api

import (
	"golang-api/internal/converters"
	"golang-api/internal/models"
	"golang-api/internal/service"
)

type IApi interface {
	AddUser(req *models.AddUserRequest) (int, error)
	GetUser(req *models.GetUserRequest) (models.User, error)
	GetUsers(req *models.GetUsersRequest) []models.User
	UpdateUser(req *models.UpdateUserRequest) error
	DeleteUser(req *models.DeleteUserRequest) error
	GetStatustics() models.Statistics
}

func New(serv service.IService) IApi {
	return &api{
		serv: serv,
	}
}

type api struct {
	serv service.IService
}

func (a *api) AddUser(req *models.AddUserRequest) (int, error) {

	res, err := a.serv.AddUser(converters.ApiUserModelToServiceModel(*req))
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (a *api) GetUser(req *models.GetUserRequest) (models.User, error) {
	return a.serv.GetUser(req.UserId)
}

func (a api) GetUsers(req *models.GetUsersRequest) []models.User {
	var users []models.User
	for _, user := range a.serv.GetUsers() {
		users = append(users, *user)
	}
	return users
}

func (a *api) UpdateUser(req *models.UpdateUserRequest) error {
	return a.serv.UpdateUser(req.Id, converters.ApiUpdateUserModelToServiceModel(*req))
}

func (a *api) DeleteUser(req *models.DeleteUserRequest) error {
	return a.serv.DeleteUser(req.Id)
}

func (a api) GetStatustics() models.Statistics {
	return a.serv.GetStatictics()
}
