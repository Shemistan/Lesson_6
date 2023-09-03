package storage

import (
	"github.com/Shemistan/Lesson_6/storage/models"
)

func New(host string, port, ttl int, conn *IConn) IStorage {
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
	conn *IConn
}
