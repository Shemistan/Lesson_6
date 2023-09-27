package storage

import (
	"errors"
	"golang-api/internal/models"
)

type IStorage interface {
	AddUser(user *models.User) (int, error)
	GetUser(userId int) (*models.User, error)
	GetUsers() map[int64]*models.User
	UpdateUser(userId int, user *models.User) error
	DeleteUser(userId int) error
}

type storage struct {
	inc int64
	db  map[int64]*models.User
}

func New() IStorage {
	return &storage{
		inc: 1,
		db:  make(map[int64]*models.User),
	}
}

func (s *storage) AddUser(user *models.User) (int, error) {
	userId := s.inc
	s.inc++
	user.Id = userId
	s.db[userId] = user
	return int(userId), nil
}

func (s *storage) GetUser(userId int) (*models.User, error) {
	val, ok := s.db[int64(userId)]
	if !ok {
		return nil, errors.New("user not found")
	}

	return val, nil
}

func (s *storage) GetUsers() map[int64]*models.User {
	return s.db
}

func (s *storage) UpdateUser(userId int, user *models.User) error {
	_, ok := s.db[int64(userId)]
	if !ok {
		return errors.New("user not found")
	}

	s.db[int64(userId)] = user
	return nil
}

func (s *storage) DeleteUser(userId int) error {
	_, ok := s.db[int64(userId)]
	if !ok {
		return errors.New("user not found")
	}

	delete(s.db, int64(userId))
	return nil
}
