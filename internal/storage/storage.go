package storage

import (
	"errors"
	"github.com/Shemistan/Lesson_6/internal/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
	"log"
	"time"
)

type IStorage interface {
	Auth(req *models.AuthRequest) (int, error)
	UpdateUser(id int, req *models.UserRequest) error
	GetUser(id int) (*models.User, error)
	GetUsers() (*[]models.User, error)
	DeleteUser(id int) error
	GetStatistics() *models.Statistics
}

func New(host string, port, ttl int, conn *Conn) IStorage {
	return &storage{
		db:         make(map[int]*models.User),
		ids:        0,
		Host:       host,
		Port:       port,
		TLL:        ttl,
		conn:       conn,
		statistics: models.Statistics{},
	}
}

type storage struct {
	db         map[int]*models.User
	ids        int
	Host       string
	Port       int
	TLL        int
	conn       *Conn
	statistics models.Statistics
}

func (s *storage) Auth(req *models.AuthRequest) (int, error) {
	if req == nil {
		return 0, errors.New("user is nil")
	}

	err := s.conn.Open()
	if err != nil {
		return 0, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	// проверяем есть ли уже пользователь в бд
	if len(s.db) > 0 {
		for _, u := range s.db {
			if u.Login == req.Login {
				return u.ID, nil
			}
		}
	}

	s.ids++
	user := converters.ApiAuthModelToServiceUserModel(s.ids, *req)
	s.db[s.ids] = user

	s.statistics.GetUserCounts++

	log.Println(s.db)
	log.Printf("user %v is added: %v", s.ids, user)

	return s.ids, nil
}

func (s *storage) UpdateUser(id int, req *models.UserRequest) error {
	if req == nil {
		return errors.New("UserRequest for update is nil")
	}

	err := s.conn.Open()
	if err != nil {
		return err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	if len(s.db) == 0 {
		return errors.New("not users in db")
	}

	_, ok := s.db[id]
	if !ok {
		return errors.New("not exist user in db")
	}

	for _, u := range s.db {
		if u.ID == id {
			u.Name = req.Name
			u.Surname = req.Surname
			u.Status = req.Status
			u.Role = req.Role
			u.UpdateDate = time.Now()

			log.Printf("update user %v", u)
			s.statistics.UpdateCount++
			return nil
		}
	}

	return nil
}

func (s *storage) GetUser(id int) (*models.User, error) {
	err := s.conn.Open()
	if err != nil {
		return &models.User{}, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	user, ok := s.db[id]
	if !ok {
		return &models.User{}, errors.New("user not found")
	}
	log.Printf("get user %v", user)

	s.statistics.GetUserCounts++

	return user, nil
}

func (s *storage) GetUsers() (*[]models.User, error) {
	err := s.conn.Open()
	if err != nil {
		return &[]models.User{}, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	if len(s.db) == 0 {
		return &[]models.User{}, errors.New("not users in db")
	}

	var users []models.User
	for _, user := range s.db {
		users = append(users, *user)
	}

	log.Printf("get users %v", users)
	s.statistics.GetUsersCounts++

	return &users, nil
}

func (s *storage) DeleteUser(id int) error {
	err := s.conn.Open()
	if err != nil {
		return err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	if len(s.db) == 0 {
		return errors.New("not users in db")
	}

	_, ok := s.db[id]
	if !ok {
		return errors.New("user not found")
	} else {
		delete(s.db, id)
		log.Printf("delete user by ID %v", id)
	}

	s.statistics.DeleteUsersCount++

	return nil
}

func (s *storage) GetStatistics() *models.Statistics {
	log.Printf("statistics %v", s.statistics)
	return &s.statistics
}

func NewConnect() *Conn {
	return &Conn{}
}

type Conn struct {
	val bool
}

func (c *Conn) Open() error {
	if c.val {
		return errors.New("failed to open Conn")
	}
	c.val = true
	return nil
}

func (c *Conn) Close() error {
	if !c.val {
		return errors.New("failed to close")
	}
	c.val = false
	return nil
}

func (c *Conn) IsClose() bool {
	if c.val {
		return false
	}

	return true
}
