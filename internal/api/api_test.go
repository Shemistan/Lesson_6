package api

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/Lesson_6/internal/models"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
)

func TestApi_Auth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для интерфейса IStorage
	storage := mock_storage.NewMockIStorage(ctrl)

	// Создаем сервис, передавая мок IStorage
	serv := New(storage)

	api := New(serv)

	tests := []struct {
		input *models.AuthRequest
		resp  int
		error error
	}{
		{input: &models.AuthRequest{
			Login:    "Test",
			Password: "Test",
		},
			resp:  1,
			error: nil,
		},
		{input: &models.AuthRequest{
			Login:    "Test",
			Password: "Test",
		},
			resp:  0,
			error: errors.New("Error"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("tc.input.Login", func(t *testing.T) {
			storage.EXPECT().Auth(tc.input).Return(tc.resp, tc.error)
			res, err := api.Auth(tc.input)
			assert.Equal(t, tc.resp, res)
			assert.Equal(t, tc.error, err)
		})
	}

	t.Run("request is nil", func(t *testing.T) {
		_, err := api.Auth(nil)
		if err == nil {
			assert.Error(t, err, "request is nil")
		}
	})

}

func TestApi_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для интерфейса IStorage
	storage := mock_storage.NewMockIStorage(ctrl)

	// Создаем сервис, передавая мок IStorage
	serv := New(storage)

	api := New(serv)

	tests := []struct {
		input *models.UserRequest
		id    int
		error error
	}{
		{input: &models.UserRequest{
			Name:             "Test",
			Surname:          "Test",
			Status:           "Active",
			Role:             "user",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
			id:    1,
			error: nil,
		},
		{input: &models.UserRequest{
			Name:             "Test",
			Surname:          "Test",
			Status:           "Active",
			Role:             "user",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
			id:    1,
			error: errors.New("bd error"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("tc.input.Login", func(t *testing.T) {
			storage.EXPECT().UpdateUser(tc.id, tc.input).Return(tc.error)
			err := api.UpdateUser(tc.id, tc.input)
			assert.Equal(t, tc.error, err)
		})
	}

	t.Run("request is nil", func(t *testing.T) {
		err := api.UpdateUser(0, nil)
		if err == nil {
			assert.Error(t, err, "request is nil")
		}
	})
}

func TestApi_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для интерфейса IStorage
	storage := mock_storage.NewMockIStorage(ctrl)

	// Создаем сервис, передавая мок IStorage
	serv := New(storage)

	api := New(serv)

	tests := []struct {
		input  int
		output *models.User
		error  error
	}{
		{input: 1,
			output: &models.User{},
			error:  nil,
		},
		{input: 1,
			output: &models.User{},
			error:  errors.New("get user error"),
		},
		{input: 1,
			output: &models.User{
				ID:               1,
				Login:            "Test",
				Name:             "Test",
				Surname:          "Test",
				Status:           "active",
				Role:             "user",
				RegistrationDate: time.Time{},
				UpdateDate:       time.Time{},
			},
			error: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("GET USER", func(t *testing.T) {
			storage.EXPECT().GetUser(tc.input).Return(tc.output, tc.error)
			res, err := api.GetUser(tc.input)
			assert.Equal(t, tc.error, err)
			assert.Equal(t, tc.output, res)
		})
	}
}

func TestApi_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для интерфейса IStorage
	storage := mock_storage.NewMockIStorage(ctrl)

	// Создаем сервис, передавая мок IStorage
	serv := New(storage)

	api := New(serv)

	tests := []struct {
		output []*models.User
		error  error
	}{
		{
			output: []*models.User{},
			error:  errors.New("get users error"),
		},
		{
			output: []*models.User{},
			error:  nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("GET USERS", func(t *testing.T) {
			storage.EXPECT().GetUsers().Return(tc.output, tc.error)
			res, err := api.GetUsers()
			assert.Equal(t, tc.error, err)
			assert.Equal(t, tc.output, res)
		})
	}
}

func TestApi_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для интерфейса IStorage
	storage := mock_storage.NewMockIStorage(ctrl)

	// Создаем сервис, передавая мок IStorage
	serv := New(storage)

	api := New(serv)

	tests := []struct {
		input int
		error error
	}{
		{input: 1,
			error: nil,
		},
		{input: 1,
			error: errors.New("get user error"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("DELETE USER", func(t *testing.T) {
			storage.EXPECT().DeleteUser(tc.input).Return(tc.error)
			err := api.DeleteUser(tc.input)
			assert.Equal(t, tc.error, err)
		})
	}
}
