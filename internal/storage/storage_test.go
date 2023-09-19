package storage

import (
	"github.com/stretchr/testify/assert"
	"golang-api/internal/models"
	"testing"
	"time"
)

func TestGetUsers(t *testing.T) {
	storage := New()
	users := storage.GetUsers()
	assert.Equal(t, users, make(map[int64]*models.User), "must be equal")
}

func TestAddUser(t *testing.T) {
	newUser := models.User{
		Login:            "manager",
		PasswordHash:     "password_hash",
		Name:             "Test_User",
		Surname:          "Surname",
		Status:           models.Active,
		Role:             models.Manager,
		RegistrationDate: time.Now(),
		UpdateDate:       time.Now(),
	}
	storage := New()
	userId, err := storage.AddUser(&newUser)
	assert.NoError(t, err, "storage should not return an error in this case")
	assert.Equal(t, 1, userId)

	user, err := storage.GetUser(userId)
	assert.NoError(t, err, "storage should not return an error in this case")
	assert.Equal(t, &newUser, user)
}

func TestGetUser(t *testing.T) {
	newUser := models.User{
		Login:            "manager",
		PasswordHash:     "password_hash",
		Name:             "Test_User",
		Surname:          "Surname",
		Status:           models.Active,
		Role:             models.Manager,
		RegistrationDate: time.Now(),
		UpdateDate:       time.Now(),
	}
	storage := New()
	userId, err := storage.AddUser(&newUser)
	assert.NoError(t, err, "storage should not return an error in this case")
	assert.Equal(t, 1, userId)

	newUser2 := models.User{
		Login:            "manager",
		PasswordHash:     "password_hash",
		Name:             "Test_User",
		Surname:          "Surname",
		Status:           models.Active,
		Role:             models.Manager,
		RegistrationDate: time.Now(),
		UpdateDate:       time.Now(),
	}

	userId2, err2 := storage.AddUser(&newUser2)
	assert.NoError(t, err2, "storage should not return an error in this case")
	assert.Equal(t, 2, userId2)

	user, err := storage.GetUser(userId)
	assert.NoError(t, err, "storage should not return an error in this case")
	assert.Equal(t, &newUser, user)
}

func TestUpdateUser(t *testing.T) {
	newUser := models.User{
		Login:            "manager",
		PasswordHash:     "password_hash",
		Name:             "Test_User",
		Surname:          "Surname",
		Status:           models.Active,
		Role:             models.Manager,
		RegistrationDate: time.Now(),
		UpdateDate:       time.Now(),
	}
	storage := New()
	userId, err := storage.AddUser(&newUser)
	assert.NoError(t, err, "storage should not return an error in this case")
	assert.Equal(t, 1, userId)

	newUser.Login = "UpdatedLogin"

	err = storage.UpdateUser(1, &newUser)
	assert.NoError(t, err, "storage should not return an error in this case")

	user, err := storage.GetUser(userId)
	assert.NoError(t, err, "storage should not return an error in this case")
	assert.Equal(t, &newUser, user)
	assert.Equal(t, user.Login, "UpdatedLogin")
}

func TestDeleteUser(t *testing.T) {
	newUser := models.User{
		Login:            "manager",
		PasswordHash:     "password_hash",
		Name:             "Test_User",
		Surname:          "Surname",
		Status:           models.Active,
		Role:             models.Manager,
		RegistrationDate: time.Now(),
		UpdateDate:       time.Now(),
	}
	storage := New()
	userId, err := storage.AddUser(&newUser)
	assert.NoError(t, err, "storage should not return an error in this case")
	assert.Equal(t, 1, userId)

	err = storage.DeleteUser(userId)
	assert.NoError(t, err, "storage should not return an error in this case")

	_, err = storage.GetUser(userId)
	assert.Error(t, err, "must me an error 'not found user")
}
