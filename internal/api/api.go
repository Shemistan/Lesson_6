package api

import (
	"github.com/Shemistan/Lesson_6/internal/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
)

type IApi interface {
	Auth(login, password string) (int64, error)
	UpdateUser(id int64, apiModel models.Account)
	GetUser(id int64) (user *models.User, err error)
	GetUsers() ([]models.User, error)
	DeleteUser(id int64) (err error)
	GetStatistics() map[string]int64
}

func New(serv service.IService) IApi {
	return &api{
		serv: serv,
	}
}

type api struct {
	serv service.IService
}

func (api *api) Auth(login, password string) (int64, error) {
	res, err := api.serv.Auth(login, password)
	if err != nil {
		return 0, err
	}
	return res, nil
}
func (api *api) UpdateUser(id int64, apiModel models.Account) {
	api.serv.UpdateUser(id, converters.ApiModelToServiceModel(apiModel))
}
func (api *api) GetUser(id int64) (user *models.User, err error) {
	result, err := api.serv.GetUser(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (api *api) GetUsers() ([]models.User, error) {
	res, err := api.serv.GetUsers()
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (api *api) DeleteUser(id int64) (er error) {
	err := api.serv.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
func (api *api) GetStatistics() map[string]int64 {
	res := api.serv.GetStatistics()
	return res
}
