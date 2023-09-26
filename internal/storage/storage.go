package storage

import (
	"errors"
	"fmt"
	"time"

	"github.com/Shemistan/Lesson_6/internal/models"
)

type IStorage interface {
	Add(user *models.SUser) (int32,error)
	Get(userId int32) (*models.SUser, error)
	GetAll()
	Update(userId int32, user *models.SUser) error
	Delete(userId int32) error
	GetMap() (map1 map[int32]*models.SUser)
}

type SStorage struct{
	idg int32
	db map[int32]*models.SUser
}

func NewIStorage() IStorage{
	return &SStorage{idg: 0,db: make(map[int32]*models.SUser)}
}

func (s *SStorage) GetMap() (map[int32]*models.SUser){
	return s.db
}

func (s *SStorage) GetAll() {
	for key, value := range s.db{
		fmt.Println("Key: ", key, "Value: ", value)
	} 
}

func (s *SStorage) Add(user *models.SUser)(int32,error){
	if user == nil{
		return 0, errors.New("Error User not found")
	}

	if s.idg == 0 {
		s.db[0]=user
		t := time.Now()
		user.RegistrationDate = t.Format("2006-01-02 15:04:05")
		s.idg++
		return 0, nil
	}
		
	s.db[s.idg]=user
	t := time.Now()
	user.RegistrationDate = t.Format("2006-01-02 15:04:05")
	s.idg++

	return s.idg, nil
	
}

func (s *SStorage) Get(userId int32) (*models.SUser, error){
	if value, ok := s.db[userId]; ok {
		return value, nil
	}

	return nil, errors.New("Error cant get user")
}

func (s *SStorage) Update(userId int32, user *models.SUser) error{
	if user == nil {
		return errors.New("Error User not found")
	}
	if _, ok := s.db[userId]; !ok {
		return errors.New("Error USerID not found")
	}

	for key := range s.db {
		if key == userId {
			s.db[userId] = user
			t := time.Now()
			user.UpdateDate = t.Format("2006-01-02 15:04:05")
			return nil
		}
	}

	return errors.New("Error Cant Update User")
}

func (s *SStorage) Delete(userId int32) error{
	for key := range s.db {
		if key == userId {
			delete(s.db, userId)
			return nil
		}
	}
	
	return errors.New("Error Cant Delete User")
}




