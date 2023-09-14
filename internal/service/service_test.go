package service

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/Lesson_6/internal/models"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок для интерфейса IStorage
	storage := mock_storage.NewMockIStorage(ctrl)

	// Создаем сервис, передавая мок IStorage
	serv := New(storage)

	statistic := &models.Statistics{}

	t.Run("request is nil", func(t *testing.T) {
		storage.EXPECT().Auth(nil).Return(0, errors.New("some error"))
		storage.EXPECT().UpdateUser(0, nil).Return(errors.New("some error"))
		storage.EXPECT().GetUser(0).Return(nil, errors.New("some error"))
		storage.EXPECT().GetUsers().Return(nil, errors.New("some error"))

		_, err := serv.Auth(nil)
		statistic.GetUserCounts++
		if err == nil {
			assert.Error(t, err, "ошибка при выполнении аутентификации")
		}
		err = serv.UpdateUser(0, nil)
		statistic.UpdateCount++
		if err == nil {
			assert.Error(t, err, "ошибка при выполнении UpdateUser")
		}

		_, err = serv.GetUser(0)
		statistic.GetUserCounts++
		if err == nil {
			assert.Error(t, err, "ошибка при выполнении GetUser")
		}
		_, err = serv.GetUsers()
		statistic.GetUsersCounts++
		if err == nil {
			assert.Error(t, err, "ошибка при выполнении GetUsers")
		}
	})

	t.Run("Auth", func(t *testing.T) {
		statistic.GetUserCounts++

		auth := models.AuthRequest{
			Login:    "Test",
			Password: "123123",
		}
		storage.EXPECT().Auth(gomock.Any()).Return(1, nil)

		res, err := serv.Auth(&auth)

		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 1, res)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		statistic.UpdateCount++

		storage.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)

		err := serv.UpdateUser(1, nil)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, nil, err)
	})

	t.Run("GetUser", func(t *testing.T) {
		statistic.GetUserCounts++

		user := models.User{
			ID:               0,
			Login:            "",
			Name:             "",
			Surname:          "",
			Status:           "",
			Role:             "",
			RegistrationDate: time.Time{},
			UpdateDate:       time.Time{},
		}

		storage.EXPECT().GetUser(gomock.Any()).Return(&user, nil)

		res, err := serv.GetUser(1)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, user, *res)
	})

	t.Run("GetUsers", func(t *testing.T) {
		statistic.GetUsersCounts++

		var users []*models.User

		storage.EXPECT().GetUsers().Return(users, nil)

		res, err := serv.GetUsers()

		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, users, res)
	})

	t.Run("DeleteUser return error", func(t *testing.T) {
		statistic.DeleteUsersCount++

		storage.EXPECT().DeleteUser(gomock.Any()).Return(errors.New("DeleteUser error"))

		err := serv.DeleteUser(1)

		assert.Equal(t, errors.New("DeleteUser error"), err)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		statistic.DeleteUsersCount++

		storage.EXPECT().DeleteUser(gomock.Any()).Return(nil)

		err := serv.DeleteUser(1)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, nil, err)
	})

	t.Run("Get statistic", func(t *testing.T) {
		result := serv.GetStatistics()

		assert.Equal(t, *statistic, *result)
	})

}
