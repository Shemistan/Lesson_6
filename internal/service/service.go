package service

import (
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

type IService interface {
	Auth(user *models.User)(int, error)
	// Get(userId int) (*models.User, error)
	// GetUsers() (map[int]*models.User, error)
	// Update(userId int, user *models.UserDate) error
	// Delete(userID int) error
}

type service struct {
	CounterAuth int
	repo        storage.IStorage
}

func New(repo storage.IStorage)IService{
	return &service{
		CounterAuth: 0,
		repo: repo,
	}
}

func (s *service)Auth(user *models.User)(int, error){
	res, err := s.repo.Auth(user)
	if err != nil {
		return 0, err
	}
	s.CounterAuth++
	return res, nil
}