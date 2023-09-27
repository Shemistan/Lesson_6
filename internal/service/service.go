package service

import (
	"github.com/Shemistan/Lesson_6/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

type IService interface {
	Add(user *models.User) (int64, error)
	UpdateUser(userId int64, user *models.UpdateUserData) error
	GetUserByLogin(login string, password string) (int64, error)
	GetUser(userId int64) (*models.GetUserData, error)
	GetUsers() ([]*models.GetUserData, error)
	DeleteUser(userId int64) error
	GetStatistics() *models.Statistic
}

type service struct {
	stats *models.Statistic
	repo  storage.IStorage
}

func New(repo storage.IStorage) IService {
	return &service{
		stats: &models.Statistic{
			DeletedUsersAccount: 0,
			UpdateCount:         0,
			GetUserCount:        0,
			GetUsersCount:       0,
		},
		repo: repo,
	}
}

func (s *service) Add(user *models.User) (int64, error) {
	res, err := s.repo.Add(user)

	if err != nil {
		return 0, err
	}

	s.stats.UsersCount++

	return res, nil

}

func (s *service) UpdateUser(userId int64, user *models.UpdateUserData) error {
	s.stats.UpdateCount++

	err := s.repo.Update(userId, user)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUserByLogin(login string, password string) (int64, error) {
	id, err := s.repo.GetByLogin(login, password)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) GetUser(userId int64) (*models.GetUserData, error) {
	s.stats.GetUserCount++

	user, err := s.repo.Get(userId)

	if err != nil {
		return nil, err
	}

	getUser := converters.GetServiceToApiModel(user)

	return getUser, nil
}

func (s *service) GetUsers() ([]*models.GetUserData, error) {
	s.stats.GetUsersCount++

	users, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) DeleteUser(userId int64) error {
	s.stats.DeletedUsersAccount++

	err := s.repo.Delete(userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetStatistics() *models.Statistic {
	return s.stats
}
