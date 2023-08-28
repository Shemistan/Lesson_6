package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

type IService interface {
	Auth(req *models.AuthRequest) (int, error)
	UpdateUser(id int, req *models.UserRequest) error
	GetUser(id int) (*models.User, error)
	GetUsers() (*[]models.User, error)
	DeleteUser(id int) error
	GetStatistics() *models.Statistics
}

func New(repo storage.IStorage) IService {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo storage.IStorage
}

func (s *service) Auth(req *models.AuthRequest) (int, error) {
	//if req.Login == "" {
	//	return 0, errors.New("login is empty")
	//}
	//if req.Password == "" {
	//	return 0, errors.New("password is empty")
	//}

	res, err := s.repo.Auth(req)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows")
	}

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *service) UpdateUser(id int, req *models.UserRequest) error {
	err := s.repo.UpdateUser(id, req)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows")
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUser(id int) (*models.User, error) {
	res, err := s.repo.GetUser(id)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows")
	}

	if err != nil {
		return &models.User{}, err
	}

	return res, nil
}

func (s *service) GetUsers() (*[]models.User, error) {
	res, err := s.repo.GetUsers()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows")
	}

	if err != nil {
		return &[]models.User{}, err
	}

	return res, nil
}

func (s *service) DeleteUser(id int) error {
	err := s.repo.DeleteUser(id)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows")
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetStatistics() *models.Statistics {
	return s.repo.GetStatistics()
}
