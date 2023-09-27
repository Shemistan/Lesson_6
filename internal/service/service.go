package service

import (
	"fmt"
	"golang-api/internal/models"
	"golang-api/internal/storage"
	"time"
)

type IService interface {
	AddUser(user models.User) (int, error)
	GetUser(userId int) (models.User, error)
	GetUsers() map[int64]*models.User
	UpdateUser(userId int, user models.User) error
	DeleteUser(userId int) error
	GetStatictics() models.Statistics
}

func New(repo storage.IStorage) IService {
	return &service{
		statistics: models.Statistics{},
		repo:       storage.New(),
	}
}

type service struct {
	statistics models.Statistics
	repo       storage.IStorage
}

func (s *service) AddUser(user models.User) (int, error) {
	user.Status = models.Active
	user.Role = models.Manager
	user.RegistrationDate = time.Now()
	user.UpdateDate = time.Now()
	userId, err := s.repo.AddUser(&user)
	if err != nil {
		fmt.Print("failed to add a user")
		return 0, err
	}

	s.statistics.AddUserCount++
	return userId, nil

}

func (s *service) GetUser(userId int) (models.User, error) {
	user, err := s.repo.GetUser(userId)
	if err != nil {
		fmt.Println("cannot get a user")
		return models.User{}, err
	}

	s.statistics.GetUserCount++
	return *user, nil
}

func (s *service) GetUsers() map[int64]*models.User {
	s.statistics.GetUsersCount++
	return s.repo.GetUsers()
}

func (s *service) UpdateUser(userId int, user models.User) error {
	s.statistics.UpdatedUsersCount++
	user.UpdateDate = time.Now()
	err := s.repo.UpdateUser(userId, &user)
	if err != nil {
		fmt.Println("failed to update a user")
		return err
	}
	return nil
}

func (s *service) DeleteUser(userId int) error {
	s.statistics.DeletedUsersCount++
	err := s.repo.DeleteUser(userId)
	if err != nil {
		fmt.Println("failed to update a user")
		return err
	}
	return nil
}

func (s *service) GetStatictics() models.Statistics {
	return s.statistics
}
