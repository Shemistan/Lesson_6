package api

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/helpers/converters"
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
	user := converters.AuthToUserModel(req)

	id, err := a.service.Auth(&user)

	if err != nil {
		return dtos.AuthResponse{}, err
	}

	return dtos.AuthResponse{
		Id: id,
	}, nil
}

func (a api) UpdateUser(req dtos.UpdateUserRequest) (dtos.UpdateUserResponse, error) {
	user := converters.UserDtoToModel(req.User)

	err := a.service.UpdateUser(req.Id, &user)

	if err != nil {
		return dtos.UpdateUserResponse{
			Success: false,
		}, err
	}

	return dtos.UpdateUserResponse{
		Success: true,
	}, nil
}

func (a api) GetUser(req dtos.GetUserRequest) (dtos.UserDto, error) {
	user, err := a.service.GetUser(req.Id)

	if err != nil {
		return dtos.UserDto{}, err
	}

	return converters.UserModelToDto(*user), nil
}

func (a api) GetUsers() ([]dtos.UserDto, error) {
	users, err := a.service.GetUsers()

	if err != nil {
		return []dtos.UserDto{}, err
	}

	usersDtos := make([]dtos.UserDto, len(users))

	for _, user := range users {
		usersDtos = append(usersDtos, converters.UserModelToDto(*user))
	}

	return usersDtos, nil
}

func (a api) DeleteUser(req dtos.DeleteUserRequest) error {
	return a.service.DeleteUser(req.Id)
}

func (a api) GetStatistics() dtos.Statistics {
	return converters.StatisticModelToDto(*a.service.GetStatistics())
}
