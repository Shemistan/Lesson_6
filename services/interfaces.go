package services

import (
	"github.com/Shemistan/Lesson_6/storage/models"
)

type IService interface {
	Auth(req models.User) (int, error)
	UpdateUser(id int, user models.User) error
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	DeleteUser(id int) error
	GetStatistics() models.Statistic
}
