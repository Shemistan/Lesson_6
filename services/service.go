package services

import (
	"github.com/Shemistan/Lesson_6/storage"
	"github.com/Shemistan/Lesson_6/storage/models"
)

type service struct {
	repo storage.IStorage
}

func New(repo storage.IStorage) IService {
	return &service{
		repo,
	}
}

func (s service) Auth(user *models.User) (int, error) {
	id, err := s.repo.Add(user)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s service) UpdateUser(id int, user *models.User) error {

	panic("implement me")
}

func (s service) GetUser(id int) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetUsers() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func (s service) GetStatistics() models.Statistic {
	return s.repo.GetStatistics()
}
