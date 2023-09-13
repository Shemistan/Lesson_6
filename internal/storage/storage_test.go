package storage

import (
	"github.com/Shemistan/Lesson_5/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(Storage)
	newStorage.db = make(map[int]*models.User)
	userData := &models.User{
		Login:            "login1",
		PasswordHash:     "string",
		Name:             "Name1",
		Surname:          "Surname1",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}

	got, err := newStorage.Add(userData)
	asserts.Nil(err)
	asserts.Equal(1, got)

}

func TestGet(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(Storage)
	newStorage.db = make(map[int]*models.User)
	userData := &models.User{
		Login:            "login2",
		PasswordHash:     "string",
		Name:             "Name2",
		Surname:          "Surname2",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	gotId, _ := newStorage.Add(userData)
	got, err := newStorage.Get(gotId)
	asserts.Nil(err)
	asserts.Equal("Name2", got.Name)
}

func TestGetUsers(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(Storage)
	newStorage.db = make(map[int]*models.User)
	userData := &models.User{
		Login:            "login3",
		PasswordHash:     "string",
		Name:             "Name3",
		Surname:          "Surname3",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	_, err := newStorage.Add(userData)
	if err == nil {
		got, err := newStorage.GetUsers()
		asserts.Nil(err)
		asserts.Equal("Name3", got[1].Name)
	}
}

func TestUpdate(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(Storage)
	newStorage.db = make(map[int]*models.User)
	userData := &models.User{
		Login:            "login4",
		PasswordHash:     "string",
		Name:             "Name4",
		Surname:          "Surname4",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	updated := &models.UserDate{
		Name:       "Name5",
		Surname:    "Surname5",
		Status:     "Active",
		Role:       "User",
		UpdateDate: "string",
	}
	gotId, _ := newStorage.Add(userData)
	err := newStorage.Update(gotId, updated)
	asserts.Nil(err)
}

func TestDelete(t *testing.T) {
	asserts := assert.New(t)
	var newStorage = new(Storage)
	newStorage.db = make(map[int]*models.User)
	userData := &models.User{
		Login:            "login2",
		PasswordHash:     "string",
		Name:             "Name2",
		Surname:          "Surname2",
		Status:           "Active",
		Role:             "User",
		RegistrationDate: "string",
		UpdateDate:       "string",
	}
	gotId, _ := newStorage.Add(userData)
	err := newStorage.Delete(gotId)
	asserts.Nil(err)
}
