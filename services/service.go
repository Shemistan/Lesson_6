package services

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/storage"
)

type service struct {
	repo storage.IStorage
}

func New(repo storage.IStorage) IService {
	return &service{
		repo,
	}
}

func (s service) Auth(req dtos.AuthRequest) (dtos.AuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) UpdateUser(req dtos.UpdateUserRequest) (dtos.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetUser(req dtos.GetUserRequest) (dtos.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetUsers(req dtos.GetUsersRequest) (dtos.GetUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteUser(req dtos.DeleteUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s service) GetStatistics(req dtos.GetStatisticsRequest) (dtos.GetStatisticsResponse, error) {
	//TODO implement me
	panic("implement me")
}
