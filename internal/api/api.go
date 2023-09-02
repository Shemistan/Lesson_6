package api

import (
	"github.com/Shemistan/Lesson_6/internal/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
)

type IApi interface {
	Auth(req *models.AuthRequest) (int, error)
	UpdateUser(id int, req *models.UserRequest) error
	GetUser(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	DeleteUser(id int) error
	GetStatistics() *models.Statistics
}

func New(serv service.IService) IApi {
	return &api{serv: serv}
}

type api struct {
	serv service.IService
}

func (a *api) Auth(req *models.AuthRequest) (int, error) {
	err := ValidateAuthRequest(req)
	if err != nil {
		return 0, err
	}

	user := converters.ApiAuthModelToServiceUserModel(*req)

	res, err := a.serv.Auth(user)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (a *api) UpdateUser(id int, req *models.UserRequest) error {
	err := ValidateUpdateUser(id, req)
	if err != nil {
		return err
	}

	err = a.serv.UpdateUser(id, req)
	if err != nil {
		return err
	}

	return nil
}

func (a *api) GetUser(id int) (*models.User, error) {
	res, err := a.serv.GetUser(id)
	if err != nil {
		return &models.User{}, err
	}

	return res, nil
}

func (a *api) GetUsers() ([]*models.User, error) {
	res, err := a.serv.GetUsers()
	if err != nil {
		return []*models.User{}, err
	}

	return res, nil
}

func (a *api) DeleteUser(id int) error {
	err := a.serv.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *api) GetStatistics() *models.Statistics {
	res := a.serv.GetStatistics()

	return res
}
