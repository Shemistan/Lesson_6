package services

import (
	models2 "github.com/Shemistan/Lesson_6/models"
	"github.com/Shemistan/Lesson_6/storage"
)

type service struct {
	repo storage.IStorage
}

func New(repo storage.IStorage) IService {
	return &service{
		repo,
	}
}

func (s *service) Auth(user *models2.User) (int, error) {
	id, err := s.repo.Add(user)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) UpdateUser(id int, user *models2.User) error {
	s.repo.GetStatistics().UpdateCount++

	return s.repo.Update(id, user)
}

func (s *service) GetUser(id int) (*models2.User, error) {
	s.repo.GetStatistics().GetUserCount++

	user, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUsers() ([]*models2.User, error) {
	s.repo.GetStatistics().GetUsersCount++

	users, err := s.repo.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) DeleteUser(id int) error {
	s.repo.GetStatistics().DeletedUsersCount++

	return s.repo.Delete(id)
}

func (s *service) GetStatistics() *models2.Statistic {
	return s.repo.GetStatistics()
}
