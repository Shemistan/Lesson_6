package storage

import (
	"errors"
	"fmt"
	"log"

	"github.com/Shemistan/Lesson_6/internal/models"
)

type IStorage interface {
	Auth(user *models.User)(int, error)
	// Get(userId int) (*models.User, error)
	// GetUsers() (map[int]*models.User, error)
	// Update(userId int, user *models.UserDate) error
	// Delete(userID int) error
}

type storage struct {
	ids int
	Host string
	Port int
	TLL int
	conn *Conn
	db map[int]*models.User
}

func New(host string, port, ttl int, conn *Conn)IStorage {
	return &storage{
		db:   make(map[int]*models.User),
		ids:  0,
		Host: host,
		Port: port,
		TLL:  ttl,
		conn: conn,
	}
}

func (s *storage) Auth(user *models.User)(int, error) {
	if user == nil{
		return 0, errors.New("Not founded")
	}

	s.ids ++
	log.Printf("user %v is add: %v", s.ids, user)
	fmt.Println("dsada")
	for k, v := range s.db {
		if v.Login == user.Login {
			if v.HashPassword == user.HashPassword {
				return k, nil
			} else {
				return 0, errors.New("Wrong password")
			}
		} else {
			return k, errors.New("user exists in the database")
		}
	}

	return s.ids, nil
}

func NewConnect() *Conn {
	return &Conn{}
}

type Conn struct {
	val bool
}

func (c *Conn) Close() error {
	if !c.val {
		return errors.New("errors closing")
	}

	return nil
}

func (c *Conn) Open() error {
	if c.val {
		return errors.New("errors opening")
	}

	return nil
}

