package storage

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Shemistan/Lesson_6/internal/models"
)

type IStorage interface {
	Add(user *models.User) (int64, error)
	Update(id int64, user *models.User) error
	Get(id int64) (models.User, error)
	GetAll() []*models.User
	Delete(id int64) (int64, error)
}

func New(host string, port int64) IStorage {
	return &storage{
		db:        make(map[int64]*models.User),
		host:      host,
		port:      port,
		idCounter: 0,
	}
}

type storage struct {
	db        map[int64]*models.User
	host      string
	port      int64
	idCounter int64
}

func (s *storage) Add(user *models.User) (int64, error) {
	if user == nil {
		return 0, errors.New("user is null")
	}

	existingUser, ok := s.db[user.Id]

	if ok {
		log.Printf("user with id %d already exists", user.Id)
		return existingUser.Id, nil
	}

	

	s.idCounter++
	user.Id = s.idCounter
	s.db[s.idCounter] = user
	log.Printf("user with login: %s is added - Created at: %s\n", user.Login, user.RegistrationDate)
	return user.Id, nil
}

func (s *storage) Update(id int64, updatedUser *models.User) error {

	if id < 0 {
		return errors.New("id can't be less than 0")
	}

	user, ok := s.db[id]

	if !ok {
		return fmt.Errorf("user with id: %d not found", id)
	}

	if updatedUser.Firstname != "" {
		user.Firstname = updatedUser.Firstname
	}

	if updatedUser.Lastname != "" {
		user.Lastname = updatedUser.Lastname
	}

	user.UpdatedDate = time.Now().String()

	return nil

}

func (s *storage) Get(id int64) (models.User, error) {
	user, ok := s.db[id]

	if !ok {
		return models.User{}, fmt.Errorf("user with id %d doesn't exist", id)
	}
	log.Printf("id: %d, login: %s, firstname: %s, lastname: %s", user.Id, user.Login, user.Firstname, user.Lastname)
	return *user, nil
}

func (s *storage) GetAll() []*models.User {
	var userList []*models.User
	for _, user := range s.db {
		userList = append(userList, user)
	}
	log.Printf("users: %+v", userList)
	return userList
}

func (s *storage) Delete(id int64) (int64, error) {
	if len(s.db) == 0 {
		return 0, errors.New("there are no users to delete")
	}

	if id <= 0 {
		return 0, errors.New("invalid id")
	}

	_, ok := s.db[id]

	if !ok {
		return 0, fmt.Errorf("user with id %d doesn't exist", id)
	}

	delete(s.db, id)
	log.Printf("user with id %d was successfully deleted", id)

	return id, nil
}
