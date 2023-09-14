package service

import (
	"log"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

type IService interface {
	Auth(req *models.AuthRequest) (int, error)
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

func (s *service) Auth(req *models.AuthRequest) (int, error) {
	s.statistics.GetUserCounts++

	res, err := s.repo.Auth(req)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *service) UpdateUser(id int, req *models.UserRequest) error {
	s.statistics.UpdateCount++

	err := s.repo.UpdateUser(id, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUser(id int) (*models.User, error) {
	s.statistics.GetUserCounts++

	res, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) GetUsers() ([]*models.User, error) {
	s.statistics.GetUsersCounts++

	res, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) DeleteUser(id int) error {
	s.statistics.DeleteUsersCount++

	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetStatistics() *models.Statistics {
	log.Printf("statistics %v", s.statistics)
	return &s.statistics
}
