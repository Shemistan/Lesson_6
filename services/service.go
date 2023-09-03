package services

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/storage"
	"github.com/Shemistan/Lesson_6/storage/models"
)

type service struct {
	repo storage.IStorage
}

func New(repo storage.IStorage) IService {
	return &service{
		repo,
	}
}

func (s service) Auth(req dtos.AuthRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) UpdateUser(req dtos.UpdateUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s service) GetUser(req dtos.GetUserRequest) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetUsers(req dtos.GetUsersRequest) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteUser(req dtos.DeleteUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s service) GetStatistics(req dtos.GetStatisticsRequest) models.Statistic {
	return s.repo.GetStatistics()
}
