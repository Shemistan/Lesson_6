package storage

import (
	"errors"
	"github.com/Shemistan/Lesson_6/model"
)

type IStorage interface {
	Add(*model.User) (uint32, error)
	Get(uint32) (*model.User, error)
	GetAll() ([]*model.User, error)
	Delete(uint32) error
	Update(*model.User) error
	Auth(string, string) (uint32, error)
}

type Storage struct {
	db map[uint32]*model.User
}

func NewStorage() *Storage {
	return &Storage{
		db: make(map[uint32]*model.User),
	}
}

func (storage *Storage) Add(user *model.User) (uint32, error) {
	storage.db[(uint32)(len(storage.db))] = user
	return (uint32)(len(storage.db)) - 1, nil
}

func (storage *Storage) Get(id uint32) (*model.User, error) {
	if id > (uint32)(len(storage.db))-1 {
		return nil, errors.New("User not found")
	}
	return storage.db[id], nil
}

func (storage *Storage) GetAll() ([]*model.User, error) {
	var user []*model.User
	for _, value := range storage.db {
		user = append(user, value)
	}
	return user, nil
}

func (storage *Storage) Delete(id uint32) error {
	if id > (uint32)(len(storage.db))-1 {
		return errors.New("User not found")
	}
	delete(storage.db, id)
	return nil
}

func (storage *Storage) Update(id uint32, user *model.User) error {
	if int(id) > len(storage.db)-1 {
		return errors.New("User not found")
	}
	user.RegistrationDate = storage.db[id].RegistrationDate
	storage.db[id] = user
	return nil
}

func (storage *Storage) Auth(login string, hashPassword uint32) (int32, error) {
	for id, user := range storage.db {
		if user.Login == login && user.HashPassword == hashPassword {
			return int32(id), nil
		}
	}
	return -1, errors.New("User not found")
}
