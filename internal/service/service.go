package service

import (
	"errors"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

var Cache = map[string]int64{
	"Auth":          0,
	"UpdateUser":    0,
	"GetUser":       0,
	"GetUsers":      0,
	"DeleteUser":    0,
	"GetStatistics": 0,
	"DeletedUsers":  0,
}

type IService interface {
	Auth(login, password string) (int64, error)
	UpdateUser(id int64, user *models.User)
	GetUser(id int64) (user *models.User, err error)
	GetUsers() ([]*models.User, error)
	DeleteUser(id int64) (err error)
	GetStatistics() map[string]int64
}

func New(db storage.IStorage) IService {
	return &service{
		repo: db,
	}
}

type service struct {
	repo storage.IStorage
}

func (service *service) Auth(login, password string) (int64, error) {
	Cache["Auth"]++
	id, err := service.repo.Auth(login, password)
	if err != nil {
		return 0, errors.New("error authenticating")
	}
	return id, nil
}

func (service *service) UpdateUser(id int64, user *models.User) {
	Cache["UpdateUser"]++
	service.repo.UpdateUser(id, user.Name, user.Surname, user.Active, user.Role)
}

func (service *service) GetUser(id int64) (user *models.User, err error) {
	Cache["GetUser"]++
	result, err := service.repo.GetUser(id)
	if err != nil {
		return nil, errors.New("wrong id")
	}
	return result, nil
}

func (service *service) GetUsers() ([]*models.User, error) {
	res, err := service.repo.GetUsers()
	Cache["GetUsers"]++
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *service) DeleteUser(id int64) error {
	Cache["DeleteUser"]++
	error := service.repo.DeleteUser(id)
	if error != nil {
		return error
	}
	Cache["DeletedUsers"]++
	return nil
}

func (service *service) GetStatistics() map[string]int64 {
	Cache["GetStatistics"]++
	return Cache
}
