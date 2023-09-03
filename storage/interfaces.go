package storage

import (
	"github.com/Shemistan/Lesson_6/storage/models"
)

type IStorage interface {
	Add(user *models.User) (int, error)
	Get(userId int) (*models.User, error)
	Update(userId int, user *models.User) error
	Delete(userID int) error
	GetStatistics() *models.Statistic
}

type IConn interface {
	Open() error
	Close() error
	IsClose() bool
}
