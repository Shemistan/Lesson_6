package storage

import (
	"errors"
	"log"
	"time"

	"github.com/Shemistan/Lesson_6/internal/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
)

type IStorage interface {
	Auth(req *models.AuthRequest) (int, error)
	UpdateUser(id int, req *models.UserRequest) error
	GetUser(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	DeleteUser(id int) error
	GetStatistics() *models.Statistics
}

func New(host string, port, ttl int, conn *Conn) IStorage {
	return &storage{
		db:   make(map[int]*models.User),
		ids:  0,
		Host: host,
		Port: port,
		TLL:  ttl,
		conn: conn,
	}
}

type storage struct {
	db   map[int]*models.User
	ids  int
	Host string
	Port int
	TLL  int
	conn *Conn
}

func (s *storage) GetStatistics() *models.Statistics {
	//TODO implement me
	panic("implement me")
}

func (s *storage) Auth(req *models.AuthRequest) (int, error) {
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
	//добавляем юзеру айди
	user := converters.ApiAuthModelToServiceUserModel(*req)
	user.ID = s.ids

	// добавляем в бд нового юзера
	s.db[s.ids] = user

	log.Printf("user %v is added: %v", s.ids, req)

	return s.ids, nil
}

func (s *storage) UpdateUser(id int, req *models.UserRequest) error {
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

	user, ok := s.db[id]
	if !ok {
		return errors.New("failed to updating user: user in not found")
	} else {
		user.Name = req.Name
		user.Surname = req.Surname
		user.Status = req.Status
		user.Role = req.Role
		user.UpdateDate = time.Now()

		log.Printf("update user %v", user)
	}

	return nil
}

func (s *storage) GetUser(id int) (*models.User, error) {
	err := s.conn.Open()
	if err != nil {
		return nil, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	user, ok := s.db[id]
	if !ok {
		return nil, errors.New("failed to get user: user not found")
	}
	log.Printf("get user %v", user)

	return user, nil
}

func (s *storage) GetUsers() ([]*models.User, error) {
	err := s.conn.Open()
	if err != nil {
		return nil, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	if len(s.db) == 0 {
		return nil, nil
	}

	var users []*models.User
	for _, user := range s.db {
		users = append(users, user)
	}

	log.Printf("get users %+v", users)

	return users, nil
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

	_, ok := s.db[id]
	if !ok {
		return errors.New("failed to delete user: user not found")
	} else {
		delete(s.db, id)
		log.Printf("delete user by ID %v", id)
	}

	return nil
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
