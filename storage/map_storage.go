package storage

import (
	"errors"
	"fmt"
	"github.com/Shemistan/Lesson_6/models"
	"log"
	"time"
)

func New(host string, port, ttl int, conn IConn) IStorage {
	return &storage{
		Statistic: models.Statistic{
			DeletedUsersCount: 0,
			UpdateCount:       0,
			GetUserCount:      0,
			GetUsersCount:     0,
		},
		db:   make(map[int]*models.User),
		ids:  0,
		Host: host,
		Port: port,
		TLL:  ttl,
		conn: conn,
	}
}

type storage struct {
	Statistic models.Statistic
	db        map[int]*models.User
	ids       int
	Host      string
	Port      int
	TLL       int
	conn      IConn
}

func (s *storage) Add(user *models.User) (int, error) {
	if user == nil {
		return 0, errors.New("user is nil")
	}

	err := s.conn.Open()
	if err != nil {
		return 0, err
	}

	defer s.closeConn()

	if s.isUserExists(user.Login) {
		return 0, errors.New(
			fmt.Sprintf("user with this login already exists %s", user.Login),
		)
	}

	s.ids++
	user.Id = s.ids
	s.db[s.ids] = user

	log.Printf("user %v is added: %v", s.ids, user)

	return s.ids, nil
}

func (s *storage) Get(userId int) (*models.User, error) {
	err := s.conn.Open()
	if err != nil {
		return nil, err
	}

	defer s.closeConn()

	user, err := s.getUserById(userId)
	if err != nil {
		return nil, err
	}

	log.Printf("get user %v", user)

	return user, nil
}

func (s *storage) GetUsers() ([]*models.User, error) {
	err := s.conn.Open()
	if err != nil {
		return nil, err
	}

	defer s.closeConn()

	users := make([]*models.User, 0, len(s.db))

	for _, value := range s.db {
		users = append(users, value)
	}

	return users, nil
}

func (s *storage) Update(userId int, user *models.User) error {
	if user == nil {
		return errors.New("update data is nil")
	}

	err := s.conn.Open()
	if err != nil {
		return err
	}

	defer s.closeConn()

	dbUser, err := s.getUserById(userId)
	if err != nil {
		return err
	}

	dbUser.Name = user.Name
	dbUser.Surname = user.Surname
	dbUser.Status = user.Status
	dbUser.Role = user.Role
	dbUser.UpdateDate = time.Now()

	log.Printf("update user %v", user)

	return nil
}

func (s *storage) Delete(userId int) error {
	err := s.conn.Open()
	if err != nil {
		return err
	}

	defer s.closeConn()

	user, err := s.getUserById(userId)
	if err != nil {
		return err
	}

	delete(s.db, userId)
	log.Printf("delete user %v", user)

	return nil
}

func (s *storage) GetStatistics() *models.Statistic {
	return &s.Statistic
}

func (s *storage) getUserById(id int) (*models.User, error) {
	user, ok := s.db[id]

	if !ok {
		return nil, errors.New(
			fmt.Sprintf("user with this id not exists %d", user.Id),
		)
	}

	return user, nil
}

func (s *storage) isUserExists(login string) bool {
	if len(s.db) > 0 {
		for _, dbUser := range s.db {
			if dbUser.Login == login {
				return true
			}
		}
	}
	return false
}

func (s *storage) closeConn() {
	errClose := s.conn.Close()
	if errClose != nil {
		log.Println(errClose)
	}
}
