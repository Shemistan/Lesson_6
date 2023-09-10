package storage

import (
	"errors"
	"github.com/Shemistan/Lesson_6/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConn(t *testing.T) {
	t.Run("NewConn should return new conn", func(t *testing.T) {
		conn := NewConn()

		if conn == nil {
			t.Error("conn is nil")
		}

		err := conn.Open()

		if err != nil {
			t.Error(err)
		}

		if conn.IsClose() {
			t.Error("conn is close after open")
		}

		err = conn.Close()

		if err != nil {
			t.Error(err)
		}

		if !conn.IsClose() {
			t.Error("conn is open after close")
		}
	})
}

func TestStorage(t *testing.T) {
	host := "localhost"
	port := 8080
	ttl := 10

	t.Run("NewStorage should return new storage", func(t *testing.T) {

		storage := New(host, port, ttl, NewConn())

		if storage == nil {
			t.Error("storage is nil")
		}
	})

	t.Run("Add should return error if user is nil", func(t *testing.T) {
		storage := New("localhost", port, 10, NewConn())

		_, err := storage.Add(nil)

		if err == nil {
			t.Error("err is nil")
		}

		assert.Equal(t, err, errors.New("user is nil"))
	})

	t.Run("Add should return error if user with this login already exists", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())

		user := &models.User{
			Login: "test",
		}

		_, err := storage.Add(user)
		_, err = storage.Add(user)

		if err == nil {
			t.Error("err is nil")
		}

		assert.Equal(t, err, errors.New("user with this login already exists test"))
	})

	t.Run("Add should return id if storage successful added user", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())

		user := &models.User{
			Login: "Denis",
		}

		id, err := storage.Add(user)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, id, 1)
	})

	t.Run("Get should return error if user with provide id not exists", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())

		_, err := storage.Get(1)

		if err == nil {
			t.Error("err is nil")
		}

		assert.Equal(t, err, errors.New("user with this id not exists 1"))
	})

	t.Run("Get should return user if storage successful get user", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())

		user := &models.User{
			Login: "Denis",
		}

		id, err := storage.Add(user)

		if err != nil {
			t.Error(err)
		}

		user, err = storage.Get(id)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Id, id)
		assert.Equal(t, "Denis", user.Login)
	})

	t.Run("Update should return error if user with provide id not exists", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())

		err := storage.Update(1, &models.User{Login: "Daniel"})

		if err == nil {
			t.Error("err is nil")
		}

		assert.Equal(t, err, errors.New("user with this id not exists 1"))
	})

	t.Run("Update should return error if update data is nil", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())

		err := storage.Update(1, nil)

		if err == nil {
			t.Error("err is nil")
		}

		assert.Equal(t, err, errors.New("update data is nil"))
	})

	t.Run("Update should successfully update user", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())

		user := &models.User{
			Name:    "Denis",
			Surname: "Denisov",
			Role:    models.RoleModerator,
			Status:  models.StatusActive,
		}

		id, err := storage.Add(user)

		if err != nil {
			t.Error(err)
		}

		err = storage.Update(id, &models.User{
			Name:    "Daniel",
			Surname: "Lee",
			Role:    models.RoleUser,
			Status:  models.StatusBlocked,
		})

		if err != nil {
			t.Error(err)
		}

		newUser, err := storage.Get(id)

		assert.Equal(t, err, nil)
		assert.Equal(t, "Daniel", newUser.Name)
		assert.Equal(t, "Lee", newUser.Surname)
		assert.Equal(t, models.RoleUser, newUser.Role)
		assert.Equal(t, models.StatusBlocked, newUser.Status)
	})

	t.Run("Delete should return error if user with provide id not exists", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())

		err := storage.Delete(1)

		if err == nil {
			t.Error("err is nil")
		}

		assert.Equal(t, errors.New("user with this id not exists 1"), err)
	})

	t.Run("Delete should successfully delete user", func(t *testing.T) {
		storage := New(host, port, ttl, NewConn())
		storage.Add(&models.User{})

		err := storage.Delete(1)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, err, nil)
	})
}
