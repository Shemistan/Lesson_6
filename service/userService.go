package service

import (
	"github.com/Shemistan/Lesson_6/model"
	"github.com/Shemistan/Lesson_6/storage"
	"hash/fnv"
	"strconv"
	"time"
)

type IService interface {
	Add(string, string, string, string) (uint32, error)
	Update(uint32, *model.User) error
	Delete(uint32) error
	Get(uint32) (*model.User, error)
	GetAll() ([]*model.User, error)
	Auth(string, string) (int32, error)
	GetStatistics() []uint32
}

type UserService struct {
	storage    *storage.Storage
	statistics model.Stats
}

func NewUserService(storage *storage.Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (us *UserService) Add(name, surname, login, password string) (uint32, error) {
	us.statistics.AddCounter++

	f := fnv.New32a()
	f.Write([]byte(password))

	return us.storage.Add(&model.User{
		Name:             name,
		Surname:          surname,
		Login:            login,
		HashPassword:     f.Sum32(),
		Status:           "Active",
		Role:             "User",
		RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
		UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
	})
}

func (us *UserService) Update(id uint32, user *model.User) error {
	us.statistics.UpdateCounter++

	f := fnv.New32a()
	f.Write([]byte(strconv.Itoa(int(user.HashPassword))))
	user.HashPassword = f.Sum32()

	user.UpdateDate = time.Now().Format("2006-01-02 15:04:05")

	return us.storage.Update(id, user)
}

func (us *UserService) Delete(id uint32) error {
	us.statistics.DeleteCounter++
	return us.storage.Delete(id)
}

func (us *UserService) Get(id uint32) (*model.User, error) {
	us.statistics.GetUserCounter++
	return us.storage.Get(id)
}

func (us *UserService) GetAll() ([]*model.User, error) {
	us.statistics.GetUsersCounter++
	return us.storage.GetAll()
}

func (us *UserService) GetStatistics() map[string]uint32 {
	outStats := make(map[string]uint32)

	outStats["AddCounter"] = us.statistics.AddCounter
	outStats["UpdateCounter"] = us.statistics.UpdateCounter
	outStats["DeleteCounter"] = us.statistics.DeleteCounter
	outStats["GetUserCounter"] = us.statistics.GetUserCounter
	outStats["GetUsersCounter"] = us.statistics.GetUsersCounter

	return outStats
}

func (us *UserService) Auth(login string, password string) (int32, error) {

	f := fnv.New32a()
	f.Write([]byte(password))

	return us.storage.Auth(login, f.Sum32())
}
