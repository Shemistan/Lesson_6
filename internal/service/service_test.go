package service

import (
	"errors"
	"testing"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	storage := storage.New()

	service := New(storage)

	t.Run("Add user", func(t *testing.T) {
		user := &models.User{
			Login:            "log",
			Password:         "pass",
			Name:             "test1",
			Surname:          "he",
			Status:           "Active",
			Role:             "User",
			RegistrationDate: "18.09.2023",
			UpdateDate:       "18.09.2023",
		}

		res, _ := service.Add(user)

		assert.Equal(t, int64(1), res)
	})

	t.Run("Update", func(t *testing.T) {
		user := &models.UpdateUserData{
			Name:    "test2",
			Surname: "he2",
			Status:  "Active",
			Role:    "User",
		}

		err := service.UpdateUser(1, user)

		assert.Equal(t, nil, err)
	})

	t.Run("Get by login", func(t *testing.T) {

		id, _ := service.GetUserByLogin("log", "pass")

		assert.Equal(t, int64(1), id)

	})

	t.Run("Get", func(t *testing.T) {

		_, err := service.GetUser(2)

		assert.Equal(t, errors.New("Such user doesn't exist"), err)

	})

	t.Run("Get all", func(t *testing.T) {

		users, err := service.GetUsers()

		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(users))

	})

	t.Run("Delete", func(t *testing.T) {

		err := service.DeleteUser(1)

		assert.Equal(t, nil, err)

	})

	t.Run("Update", func(t *testing.T) {
		user := &models.UpdateUserData{
			Name:    "test2",
			Surname: "he2",
			Status:  "Active",
			Role:    "User",
		}

		err := service.UpdateUser(1, user)

		assert.Equal(t, errors.New("Such user doesn't exist"), err)
	})

}
