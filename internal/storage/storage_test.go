package storage

import (
	"errors"
	"testing"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	storage := New()

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

		res, _ := storage.Add(user)

		assert.Equal(t, int64(1), res)
	})

	t.Run("Add user", func(t *testing.T) {

		_, err := storage.Add(nil)

		assert.Equal(t, errors.New("user's data is incorrect"), err)
	})

	t.Run("Add user", func(t *testing.T) {
		user := &models.User{
			Login:            "login",
			Password:         "password",
			Name:             "another",
			Surname:          "he12",
			Status:           "Active",
			Role:             "User",
			RegistrationDate: "18.09.2023",
			UpdateDate:       "18.09.2023",
		}

		_, err := storage.Add(user)

		assert.Equal(t, nil, err)
	})

	t.Run("Update", func(t *testing.T) {
		user := &models.UpdateUserData{
			Name:    "test2",
			Surname: "he2",
			Status:  "Active",
			Role:    "User",
		}

		err := storage.Update(1, user)

		assert.Equal(t, nil, err)
	})

	t.Run("Get by login", func(t *testing.T) {

		id, _ := storage.GetByLogin("login", "password")

		assert.Equal(t, int64(2), id)

	})

	t.Run("Get", func(t *testing.T) {

		_, err := storage.Get(2)

		assert.Equal(t, nil, err)
	})

	t.Run("Get all", func(t *testing.T) {

		users, err := storage.GetAll()

		assert.Equal(t, nil, err)
		assert.Equal(t, 2, len(users))

	})

	t.Run("Delete", func(t *testing.T) {

		err := storage.Delete(1)

		assert.Equal(t, nil, err)

	})

	t.Run("Update", func(t *testing.T) {
		user := &models.UpdateUserData{
			Name:    "test2",
			Surname: "he2",
			Status:  "Active",
			Role:    "User",
		}

		err := storage.Update(2, user)

		assert.Equal(t, nil, err)
	})

}
