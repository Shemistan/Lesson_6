package service

import (
	"log"

	"golang-api/internal/models"
	"golang-api/internal/storage"
)

type IService interface {
	Register(user *models.User) (int64, error)
	UpdateUser(id int64, updatedUser *models.User) error
	GetUser(id int64) (*models.User, error)
	GetAllUsers() []*models.User
	DeleteUser(id int64) (int64, error)
	GetStats() map[string]int
}

func New(db storage.IStorage) IService {
	return &service{
		repo: db,
	}
}

type service struct {
	repo storage.IStorage
}

func (s *service) Register(request *models.User) (int64, error) {
	id, err := s.repo.Add(request)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) UpdateUser(id int64, updatedUser *models.User) error {
	err := s.repo.Update(id, updatedUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUser(id int64) (*models.User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		log.Println(err)
		return &models.User{}, err
	}

	return user, nil
}

func (s *service) GetAllUsers() []*models.User {
	usersList := s.repo.GetAll()

	return usersList
}

func (s *service) DeleteUser(id int64) (int64, error) {
	id, err := s.repo.Delete(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) GetStats() map[string]int {
	stats := s.repo.GetStats()
	return stats
}