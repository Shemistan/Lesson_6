package storage

import (
	"errors"

	"github.com/Shemistan/Lesson_6/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
)

type IStorage interface {
	Add(user *models.User) (int64, error)
	Update(userId int64, user *models.UpdateUserData) error
	Get(userId int64) (*models.User, error)
	GetByLogin(login string, password string) (int64, error)
	GetAll() ([]*models.GetUserData, error)
	Delete(userId int64) error
}

type storage struct {
	db  map[int64]*models.User
	ids int64
}

func New() IStorage {
	return &storage{
		db:  make(map[int64]*models.User),
		ids: 0,
	}
}

func (s *storage) Add(user *models.User) (int64, error) {
	if user == nil {
		return 0, errors.New("user's data is incorrect")
	}
	s.ids++
	s.db[s.ids] = user

	return s.ids, nil
}

func (s *storage) Update(userId int64, user *models.UpdateUserData) error {
	_, ok := s.db[userId]
	if !ok {
		return errors.New("Such user doesn't exist")
	}
	s.db[userId].Name = user.Name
	s.db[userId].Surname = user.Surname
	s.db[userId].Status = user.Status
	s.db[userId].Role = user.Role

	return nil
}

func (s *storage) Get(userId int64) (*models.User, error) {
	_, ok := s.db[userId]
	if !ok {
		return nil, errors.New("Such user doesn't exist")
	}

	return s.db[userId], nil
}

func (s *storage) GetByLogin(login string, password string) (int64, error) {

	for id, v := range s.db {
		if v.Login == login && v.Password == password {
			return id, nil
		}
	}

	return 0, errors.New("Such user doesn't exist")
}

func (s *storage) GetAll() ([]*models.GetUserData, error) {
	if len(s.db) == 0 {
		return nil, errors.New("Db is empty")
	}

	users := make([]*models.GetUserData, len(s.db))
	i := 0

	for _, v := range s.db {
		users[i] = converters.GetServiceToApiModel(v)
		i++
	}

	return users, nil
}

func (s *storage) Delete(userId int64) error {
	_, ok := s.db[userId]
	if !ok {
		return errors.New("Such user doesn't exist")
	}

	delete(s.db, userId)

	return nil
}
