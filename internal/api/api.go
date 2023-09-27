package api

import (
	"log"

	"github.com/Shemistan/Lesson_6/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
)

type IApi interface {
	Auth(req *models.AddRequest) (int64, error)
	UpdateUser(userId int64, user *models.UpdateUserData) error
	GetUser(userId int64) (*models.GetUserData, error)
	GetUsers() ([]*models.GetUserData, error)
	DeleteUser(userId int64) error
	GetStatistics() *models.Statistic
}

func New(serv service.IService) IApi {
	return &api{
		serv: serv,
	}
}

type api struct {
	serv service.IService
}

func (a *api) Auth(req *models.AddRequest) (int64, error) {
	id, err := a.serv.GetUserByLogin(req.AuthParams.Login, req.AuthParams.Password)

	if err != nil {
		res, err := a.serv.Add(converters.ApiToServiceModel(req))

		if err != nil {
			return 0, err
		}
		log.Printf("user added %v", res)
		return res, nil
	}

	log.Printf("User is exist %d", id)
	return id, nil
}

func (a *api) UpdateUser(userId int64, user *models.UpdateUserData) error {
	err := a.serv.UpdateUser(userId, user)

	if err != nil {
		return err
	}

	return nil
}

func (a *api) GetUser(userId int64) (*models.GetUserData, error) {
	user, err := a.serv.GetUser(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *api) GetUsers() ([]*models.GetUserData, error) {
	users, err := a.serv.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a *api) DeleteUser(userId int64) error {
	err := a.serv.DeleteUser(userId)

	if err != nil {
		return err
	}

	return nil
}

func (a *api) GetStatistics() *models.Statistic {
	return a.serv.GetStatistics()
}
