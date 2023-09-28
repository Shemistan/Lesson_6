package api

import (
	"golang-api/internal/converters"
	"golang-api/internal/models"
	"golang-api/internal/service"
)

type IApi interface{
	Auth(req *models.AddRequest)(int, error)
	GetUser(userId int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(userId int, user *models.UserDate) error
	DeleteUser(userId int) error
	GetStatistics() map[string]int
}

func New(serv service.IService)IApi {
	return &api{serv: serv}
}

type api struct {
	serv service.IService
}

func (a *api) Auth(req *models.AddRequest)(int, error) {
	result, err := a.serv.Auth(converters.ApiAuthModelToServiceUserModel(*req))
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (a *api) GetUser(userId int)(*models.User, error) {
	result, err := a.serv.GetUser(userId)
	if err != nil {
		return &models.User{}, err
	}

	return result, nil
}

func (a *api) GetUsers()([]*models.User, error) {
	result, err := a.serv.GetUsers()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *api)UpdateUser(userId int, user *models.UserDate) error {
	err := a.serv.UpdateUser(userId, user)
	if err != nil {
		return  err
	}

	return nil
}

func (a *api)DeleteUser(userId int) error {
	err := a.serv.DeleteUser(userId)
	if err != nil {
		return  err
	}

	return nil
}

func (a *api)GetStatistics() map[string]int {
	
	return a.serv.GetStatistics()
}