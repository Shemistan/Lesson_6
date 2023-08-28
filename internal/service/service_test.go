package service

import (
	"errors"
	"github.com/Shemistan/Lesson_6/internal/models"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для интерфейса IStorage
	storage := mock_storage.NewMockIStorage(ctrl)

	// Создаем сервис, передавая мок IStorage
	serv := New(storage)

	t.Run("request is nil", func(t *testing.T) {
		storage.EXPECT().Auth(nil).Return(0, errors.New("some error"))
		storage.EXPECT().UpdateUser(0, nil).Return(errors.New("some error"))
		storage.EXPECT().GetUser(0).Return(&models.User{}, errors.New("some error"))
		storage.EXPECT().GetUsers().Return(&[]models.User{}, errors.New("some error"))

		_, err := serv.Auth(nil)
		if err == nil {
			assert.Error(t, err, "ошибка при выполнении аутентификации")
		}
		err = serv.UpdateUser(0, nil)
		if err == nil {
			assert.Error(t, err, "ошибка при выполнении UpdateUser")
		}
		_, err = serv.GetUser(0)
		if err == nil {
			assert.Error(t, err, "ошибка при выполнении GetUser")
		}
		_, err = serv.GetUsers()
		if err == nil {
			assert.Error(t, err, "ошибка при выполнении GetUsers")
		}
	})

}
