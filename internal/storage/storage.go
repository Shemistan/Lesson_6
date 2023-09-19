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
	Get(id int64) (*models.User, error)
	GetAll() []*models.User
	Delete(id int64) (int64, error)
	GetStats() map[string]int
}

func New(host string, port int64) IStorage {
	return &storage{
		db:        make(map[int64]*models.User),
		stats:     make(map[string]int),
		host:      host,
		port:      port,
		idCounter: 0,
	}
}

type storage struct {
	db        map[int64]*models.User
	stats     map[string]int
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
		log.Printf("user with %d ID already exists", existingUser.Id)
		return existingUser.Id, nil
	}

	s.idCounter++
	user.Id = s.idCounter
	s.db[s.idCounter] = user
	return user.Id, nil
}

func (s *storage) Update(id int64, updatedUser *models.User) error {

	if id <= 0 {
		return errors.New("id should be greater than 0")
	}

	user, ok := s.db[id]

	if !ok {
		return fmt.Errorf("user with id: %d not found", id)
	}

	if updatedUser.Firstname == "" && updatedUser.Lastname == "" {
		return nil
	}

	if updatedUser.Firstname != "" {
		user.Firstname = updatedUser.Firstname
	}

	if updatedUser.Lastname != "" {
		user.Lastname = updatedUser.Lastname
	}

	user.UpdatedDate = time.Now().String()
	s.stats["update_user"]++
	return nil

}

func (s *storage) Get(id int64) (*models.User, error) {
	user, ok := s.db[id]

	if !ok {
		return &models.User{}, fmt.Errorf("user with id %d doesn't exist", id)
	}
	log.Printf("id: %d, login: %s, firstname: %s, lastname: %s", user.Id, user.Login, user.Firstname, user.Lastname)
	s.stats["get_user"]++
	return user, nil
}

func (s *storage) GetAll() []*models.User {
	var userList []*models.User
	for _, user := range s.db {
		userList = append(userList, user)
	}
	log.Printf("users: %+v", userList)
	s.stats["get_users"]++
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
	s.stats["deleted_users"]++
	return id, nil
}

func (s *storage) GetStats() map[string]int {
	log.Printf("get_user: %d, get_users: %d, update_user: %d, deleted_users: %d", s.stats["get_user"], s.stats["get_users"], s.stats["update_user"], s.stats["deleted_users"])
	return s.stats
}
