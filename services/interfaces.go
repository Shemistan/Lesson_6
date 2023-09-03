package services

import (
	models2 "github.com/Shemistan/Lesson_6/models"
)

type IService interface {
	Auth(user *models2.User) (int, error)
	UpdateUser(id int, user *models2.User) error
	GetUser(id int) (*models2.User, error)
	GetUsers() ([]*models2.User, error)
	DeleteUser(id int) error
	GetStatistics() *models2.Statistic
}
