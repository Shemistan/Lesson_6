package service

import (
	"testing"
	"time"

	"golang-api/internal/models"
	"golang-api/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {

	mockStorage := storage.New()

	userService := New(mockStorage)

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

	userId, err := userService.AddUser(newUser)
	assert.NoError(t, err, "AddUser should not return an error in this case")
	assert.NotEqual(t, 0, userId, "UserID should not be zero")

	retrievedUser, err := userService.GetUser(userId)
	assert.NoError(t, err, "GetUser should not return an error in this case")

	newUser.Id = retrievedUser.Id
	newUser.UpdateDate = retrievedUser.UpdateDate
	assert.Equal(t, newUser, retrievedUser, "Retrieved user should match the added user")

	newUser.Name = "UpdatedName"
	err = userService.UpdateUser(userId, newUser)
	assert.NoError(t, err, "UpdateUser should not return an error in this case")

	updatedUser, err := userService.GetUser(userId)
	assert.NoError(t, err, "GetUser should not return an error in this case")

	newUser.UpdateDate = updatedUser.UpdateDate
	assert.Equal(t, newUser, updatedUser, "Updated user data should match")

	err = userService.DeleteUser(userId)
	assert.NoError(t, err, "DeleteUser should not return an error in this case")

	_, err = userService.GetUser(userId)
	assert.Error(t, err, "GetUser should return an error after deletion")

	users := userService.GetUsers()
	assert.NotNil(t, users, "GetUsers should not return nil")

	statistics := userService.GetStatictics()
	assert.Equal(t, 1, statistics.AddUserCount, "AddUserCount should be 1")
	assert.Equal(t, 2, statistics.GetUserCount, "GetUserCount should be 2")
	assert.Equal(t, 1, statistics.UpdatedUsersCount, "UpdatedUsersCount should be 1")
	assert.Equal(t, 1, statistics.DeletedUsersCount, "DeletedUsersCount should be 1")
	assert.Equal(t, 1, statistics.GetUsersCount, "GetUsersCount should be 1")
}
