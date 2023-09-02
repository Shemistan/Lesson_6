package service

import (
	"log"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

type IService interface {
	Auth(req *models.User) (int, error)
	UpdateUser(id int, req *models.UserRequest) error
	GetUser(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	DeleteUser(id int) error
	GetStatistics() *models.Statistics
}

func New(repo storage.IStorage) IService {
	return &service{
		repo:       repo,
		statistics: models.Statistics{},
	}
}

type service struct {
	repo       storage.IStorage
	statistics models.Statistics
}

func (s *service) Auth(req *models.User) (int, error) {
	res, err := s.repo.Auth(req)
	if err != nil {
		return 0, err
	}
	s.statistics.GetUserCounts++
	return res, nil
}

func (s *service) UpdateUser(id int, req *models.UserRequest) error {
	err := s.repo.UpdateUser(id, req)
	if err != nil {
		return err
	}

	s.statistics.UpdateCount++

	return nil
}

func (s *service) GetUser(id int) (*models.User, error) {
	res, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	s.statistics.GetUserCounts++

	return res, nil
}

func (s *service) GetUsers() ([]*models.User, error) {
	res, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	s.statistics.GetUsersCounts++

	return res, nil
}

func (s *service) DeleteUser(id int) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	s.statistics.DeleteUsersCount++

	return nil
}

func (s *service) GetStatistics() *models.Statistics {
	log.Printf("statistics %v", s.statistics)
	return &s.statistics
}
