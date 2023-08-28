package api

import (
	"errors"
	"github.com/Shemistan/Lesson_6/internal/models"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
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
		input    models.AuthRequest
		expected error
	}{
		{input: models.AuthRequest{
			Login:    "Test",
			Password: "Test",
		}, expected: nil},
		{input: models.AuthRequest{
			Login:    "",
			Password: "Test",
		}, expected: errors.New("login is empty")},
		{input: models.AuthRequest{
			Login:    "Test",
			Password: "",
		}, expected: errors.New("password is empty")},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("tc.input.Login", func(t *testing.T) {
			storage.EXPECT().Auth(&tc.input)
			_, err := api.Auth(&tc.input)
			assert.NoError(t, err)
		})
	}

}
