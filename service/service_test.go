package service

import (
	"errors"
	"github.com/Shemistan/Lesson_6/model"
	"github.com/Shemistan/Lesson_6/storage"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	st := storage.NewStorage()
	service := NewUserService(st)

	t.Run("Add", func(t *testing.T) {
		result, _ := service.Add("test", "test", "test", "123")
		assert.Equal(t, uint32(0), result)
	})

	//service.Add("test1", "test1", "test23", "123")

	t.Run("Get", func(t *testing.T) {
		_, err := service.Get(1)
		assert.Equal(t, errors.New("User not found"), err)
	})

	t.Run("Get", func(t *testing.T) {
		result, _ := service.Get(0)
		assert.Equal(t, "test", result.Name)
	})

	t.Run("GetAll", func(t *testing.T) {
		arr, _ := service.GetAll()
		assert.Equal(t, 1, len(arr))
	})

	t.Run("Update", func(t *testing.T) {
		err := service.Update(0, &model.User{
			Name:             "test1",
			Surname:          "test2",
			Login:            "test",
			HashPassword:     123,
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
			UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
		})
		assert.Equal(t, nil, err)
	})

	t.Run("Update", func(t *testing.T) {
		err := service.Update(2, &model.User{
			Name:             "test1",
			Surname:          "test2",
			Login:            "test",
			HashPassword:     123,
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
			UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
		})
		assert.Equal(t, errors.New("User not found"), err)
	})

	t.Run("Auth", func(t *testing.T) {
		result, _ := service.Auth("test", "123")
		assert.Equal(t, int32(0), result)
	})

	t.Run("Delete", func(t *testing.T) {
		err := service.Delete(1)
		assert.Equal(t, errors.New("User not found"), err)
	})

	t.Run("Delete", func(t *testing.T) {
		err := service.Delete(0)
		assert.Equal(t, nil, err)
	})
}
