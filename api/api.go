package api

import "github.com/Shemistan/Lesson_6/api/dtos"

type api struct {
}

func New() IApi {
	return &api{}
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
