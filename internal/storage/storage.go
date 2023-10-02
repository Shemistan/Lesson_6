package storage

import (
	"errors"

	"github.com/nasrullonurullaev5/Lesson_6/internal/model"
)

type IDataStorage interface {
	CreateUser(*model.User) (uint32, error)
	FindUser(uint32) (*model.User, error)
	FindAllUsers() ([]*model.User, error)
	DeleteUser(uint32) error
	UpdateUser(uint32, *model.User) error
	Authenticate(string, uint32) (int32, error)
}

type DataStorage struct {
	users map[uint32]*model.User
}

func NewDataStorage() *DataStorage {
	return &DataStorage{
		users: make(map[uint32]*model.User),
	}
}

func (dataStorage *DataStorage) CreateUser(user *model.User) (uint32, error) {
	dataStorage.users[uint32(len(dataStorage.users))] = user
	return uint32(len(dataStorage.users)) - 1, nil
}

func (dataStorage *DataStorage) FindUser(id uint32) (*model.User, error) {
	if len(dataStorage.users) > 0 && id > uint32(len(dataStorage.users))-1 {
		return nil, errors.New("user not found")
	}
	return dataStorage.users[id], nil
}

func (dataStorage *DataStorage) FindAllUsers() ([]*model.User, error) {
	var users []*model.User
	for _, value := range dataStorage.users {
		users = append(users, value)
	}
	return users, nil
}

func (dataStorage *DataStorage) DeleteUser(id uint32) error {
	if id > uint32(len(dataStorage.users))-1 {
		return errors.New("user not found")
	}
	delete(dataStorage.users, id)
	return nil
}

func (dataStorage *DataStorage) UpdateUser(id uint32, user *model.User) error {
	if int(id) > len(dataStorage.users)-1 {
		return errors.New("user not found")
	}
	user.RegistrationDate = dataStorage.users[id].RegistrationDate
	dataStorage.users[id] = user
	return nil
}

func (dataStorage *DataStorage) Authenticate(login string, hashPassword uint32) (int32, error) {
	for id, user := range dataStorage.users {
		if user.Login == login && user.HashPassword == hashPassword {
			return int32(id), nil
		}
	}
	return -1, errors.New("user not found")
}
