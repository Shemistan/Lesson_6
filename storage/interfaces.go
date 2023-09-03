package storage

import (
	models2 "github.com/Shemistan/Lesson_6/models"
)

type IStorage interface {
	Add(user *models2.User) (int, error)
	Get(userId int) (*models2.User, error)
	GetUsers() ([]*models2.User, error)
	Update(userId int, user *models2.User) error
	Delete(userID int) error
	GetStatistics() *models2.Statistic
}

type IConn interface {
	Open() error
	Close() error
	IsClose() bool
}
