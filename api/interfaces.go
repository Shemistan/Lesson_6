package api

import "github.com/Shemistan/Lesson_6/api/dtos"

type IApi interface {
	Auth(req dtos.AuthRequest) (dtos.AuthResponse, error)
	UpdateUser(req dtos.UpdateUserRequest) (dtos.UpdateUserResponse, error)
	GetUser(req dtos.GetUserRequest) (dtos.UserDto, error)
	GetUsers() ([]dtos.UserDto, error)
	DeleteUser(req dtos.DeleteUserRequest) error
	GetStatistics() dtos.Statistics
}
