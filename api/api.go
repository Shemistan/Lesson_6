package api

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/services"
)

type api struct {
	service services.IService
}

func New(service services.IService) IApi {
	return &api{
		service: service,
	}
}
func (a api) Auth(req dtos.AuthRequest) (dtos.AuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a api) UpdateUser(req dtos.UpdateUserRequest) (dtos.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a api) GetUser(req dtos.GetUserRequest) (dtos.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a api) GetUsers(req dtos.GetUsersRequest) (dtos.GetUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a api) DeleteUser(req dtos.DeleteUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (a api) GetStatistics(req dtos.GetStatisticsRequest) (dtos.GetStatisticsResponse, error) {
	//TODO implement me
	panic("implement me")
}
