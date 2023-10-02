package service

import (
	"hash/fnv"
	"strconv"
	"time"

	model "github.com/nasrullonurullaev5/Lesson_6/internal/model"
	"github.com/nasrullonurullaev5/Lesson_6/internal/storage"
)

type IUserService interface {
	Register(string, string, string, string) (uint32, error)
	Modify(uint32, *model.User) error
	Delete(uint32) error
	Fetch(uint32) (*model.User, error)
	FetchAll() ([]*model.User, error)
	Authenticate(string, string) (int32, error)
	CollectStatistics() map[string]uint32
}

type UserService struct {
	dataStorage    *storage.DataStorage
	operationStats model.Stats
}

func NewUserService(dataStorage *storage.DataStorage) *UserService {
	return &UserService{
		dataStorage: dataStorage,
	}
}

func (us *UserService) Register(name, surname, login, password string) (uint32, error) {
	us.operationStats.AddCounter++

	passwordHash := fnv.New32a()
	passwordHash.Write([]byte(password))

	return us.dataStorage.CreateUser(&model.User{
		Name:             name,
		Surname:          surname,
		Login:            login,
		HashPassword:     passwordHash.Sum32(),
		Status:           "Active",
		Role:             "User",
		RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
		UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
	})
}

func (us *UserService) Modify(id uint32, user *model.User) error {
	us.operationStats.UpdateCounter++

	passwordHash := fnv.New32a()
	passwordHash.Write([]byte(strconv.Itoa(int(user.HashPassword))))
	user.HashPassword = passwordHash.Sum32()

	user.UpdateDate = time.Now().Format("2006-01-02 15:04:05")

	return us.dataStorage.UpdateUser(id, user)
}

func (us *UserService) Delete(id uint32) error {
	us.operationStats.DeleteCounter++
	return us.dataStorage.DeleteUser(id)
}

func (us *UserService) Fetch(id uint32) (*model.User, error) {
	us.operationStats.GetUserCounter++
	return us.dataStorage.FindUser(id)
}

func (us *UserService) FetchAll() ([]*model.User, error) {
	us.operationStats.GetUsersCounter++
	return us.dataStorage.FindAllUsers()
}

func (us *UserService) CollectStatistics() map[string]uint32 {
	statsMap := make(map[string]uint32)

	statsMap["AddCounter"] = us.operationStats.AddCounter
	statsMap["UpdateCounter"] = us.operationStats.UpdateCounter
	statsMap["DeleteCounter"] = us.operationStats.DeleteCounter
	statsMap["GetUserCounter"] = us.operationStats.GetUserCounter
	statsMap["GetUsersCounter"] = us.operationStats.GetUsersCounter

	return statsMap
}

func (us *UserService) Authenticate(login string, password string) (int32, error) {
	passwordHash := fnv.New32a()
	passwordHash.Write([]byte(password))

	return us.dataStorage.Authenticate(login, passwordHash.Sum32())
}
