package services

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/storage/models"
)

type IService interface {
	Auth(req dtos.AuthRequest) (int, error)
	UpdateUser(req dtos.UpdateUserRequest) error
	GetUser(req dtos.GetUserRequest) (models.User, error)
	GetUsers(req dtos.GetUsersRequest) ([]models.User, error)
	DeleteUser(req dtos.DeleteUserRequest) error
	GetStatistics(req dtos.GetStatisticsRequest) models.Statistic
}
