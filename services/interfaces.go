package services

import (
	"github.com/Shemistan/Lesson_6/models"
)

type IService interface {
	Auth(user *models.User) (int, error)
	UpdateUser(id int, user *models.User) error
	GetUser(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	DeleteUser(id int) error
	GetStatistics() *models.Statistic
}
