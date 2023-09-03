package services

import "github.com/Shemistan/Lesson_6/api/dtos"

type IService interface {
	Auth(req dtos.AuthRequest) (int, error)
	UpdateUser(req dtos.UpdateUserRequest) error
	GetUser(req dtos.GetUserRequest) (dtos.GetUserResponse, error)
	GetUsers(req dtos.GetUsersRequest) (dtos.GetUsersResponse, error)
	DeleteUser(req dtos.DeleteUserRequest) error
	GetStatistics(req dtos.GetStatisticsRequest) dtos.GetStatisticsResponse
}
