package storage

import (
	"errors"
	"time"

	"github.com/Shemistan/Lesson_6/internal/models"
)

type IStorage interface {
	Auth(a, b string) (int64, error)
	UpdateUser(id int64, name, surname string, active bool, role string)
	GetUser(int64) (user *models.User, err error)
	GetUsers() ([]models.User, error)
	DeleteUser(int64) (err error)
}

var CurrentTime = time.Now().Format("02-01-2006 15:04")

func New() IStorage {
	return &Storage{
		dataBase: make(map[int64]*models.User),
		ID:       0,
	}
}

type Storage struct {
	dataBase map[int64]*models.User
	ID       int64
}

func (s *Storage) addUser(login, password string) int64 {
	s.ID++
	s.dataBase[s.ID] = &models.User{
		Login:            login,
		Password:         password,
		Active:           true,
		Role:             "User",
		RegistrationDate: CurrentTime,
	}
	return s.ID
}

func (s *Storage) Auth(log, psw string) (int64, error) {
	for i := range s.dataBase {
		if s.dataBase[i].Login == log {
			if s.dataBase[i].Password == psw {
				return i, nil
			} else {
				return 0, errors.New("wrong password")
			}

		}
	}
	return s.addUser(log, psw), nil
}

func (s *Storage) UpdateUser(id int64, name, surname string, active bool, role string) {
	currTime := time.Now().Format("02-01-2006 15:04")
	s.dataBase[id].Name = name
	s.dataBase[id].Surname = surname
	s.dataBase[id].Active = active
	s.dataBase[id].Role = role
	s.dataBase[id].UpdateDate = currTime
}

func (s *Storage) GetUser(id int64) (*models.User, error) {
	for key := range s.dataBase {
		if key == id {
			return s.dataBase[key], nil
		}
	}
	return nil, errors.New("there is no such user")
}

func (s *Storage) GetUsers() ([]models.User, error) {
	list := make([]models.User, len(s.dataBase))
	if len(s.dataBase) != 0 {
		for i, val := range s.dataBase {
			list[i-1] = *val
		}
		return list, nil
	} else {
		return nil, errors.New("database is empty")
	}
}

func (s *Storage) DeleteUser(id int64) (err error) {
	if _, ok := s.dataBase[id]; ok {
		delete(s.dataBase, id)
		return nil
	} else {
		return errors.New("no such id")
	}
}
