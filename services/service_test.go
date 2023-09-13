package services

import (
	"errors"
	"github.com/Shemistan/Lesson_6/models"
	mock_storage "github.com/Shemistan/Lesson_6/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)

	storage := mock_storage.NewMockIStorage(ctrl)

	serv := New(storage)

	t.Run("Auth should return error if storage return error", func(t *testing.T) {
		storage.EXPECT().Add(gomock.Any()).Return(1, errors.New("some error"))

		_, err := serv.Auth(nil)

		assert.Equal(t, errors.New("some error"), err)
	})

	t.Run("Auth should return id if storage successful auth user", func(t *testing.T) {
		mockId := 1

		storage.EXPECT().Add(gomock.Any()).Return(mockId, nil)

		id, err := serv.Auth(nil)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, mockId, id)
	})

	t.Run("UpdateUser should return error if storage return error", func(t *testing.T) {
		statistic := models.Statistic{}

		storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("some error"))
		storage.EXPECT().GetStatistics().Return(&statistic)

		err := serv.UpdateUser(1, nil)

		assert.Equal(t, errors.New("some error"), err)
	})

	t.Run("UpdateUser should return nil if storage successful update user", func(t *testing.T) {
		statistic := models.Statistic{}

		storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		storage.EXPECT().GetStatistics().Return(&statistic)

		err := serv.UpdateUser(1, nil)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, nil, err)
	})

	//Write test for all methods
	t.Run("GetUser should return error if storage return error", func(t *testing.T) {
		statistic := models.Statistic{}

		storage.EXPECT().Get(gomock.Any()).Return(nil, errors.New("some error"))
		storage.EXPECT().GetStatistics().Return(&statistic)

		_, err := serv.GetUser(1)

		assert.Equal(t, errors.New("some error"), err)
	})

	t.Run("GetUser should return user if storage successful get user", func(t *testing.T) {
		statistic := models.Statistic{}

		mockUser := models.User{
			Id:               0,
			Login:            "",
			Password:         "",
			Name:             "",
			Surname:          "",
			Status:           "",
			Role:             "",
			RegistrationDate: time.Time{},
			UpdateDate:       time.Time{},
		}

		storage.EXPECT().Get(gomock.Any()).Return(&mockUser, nil)
		storage.EXPECT().GetStatistics().Return(&statistic)

		user, err := serv.GetUser(1)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, mockUser, *user)
	})

	t.Run("GetUsers should return error if storage return error", func(t *testing.T) {
		statistic := models.Statistic{}

		var users []*models.User

		storage.EXPECT().GetUsers().Return(users, errors.New("some error"))
		storage.EXPECT().GetStatistics().Return(&statistic)

		_, err := serv.GetUsers()

		assert.Equal(t, errors.New("some error"), err)
	})

	t.Run("GetUsers should return users if storage successful get users", func(t *testing.T) {
		statistic := models.Statistic{}

		var mockUsers []*models.User

		storage.EXPECT().GetUsers().Return(mockUsers, nil)
		storage.EXPECT().GetStatistics().Return(&statistic)

		users, err := serv.GetUsers()

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, mockUsers, users)
	})

	t.Run("DeleteUser should return error if storage return error", func(t *testing.T) {
		statistic := models.Statistic{}

		storage.EXPECT().Delete(gomock.Any()).Return(errors.New("some error"))
		storage.EXPECT().GetStatistics().Return(&statistic)

		err := serv.DeleteUser(1)

		assert.Equal(t, errors.New("some error"), err)
	})

	t.Run("DeleteUser should return nil if storage successful delete user", func(t *testing.T) {
		statistic := models.Statistic{}
		storage.EXPECT().GetStatistics().Return(&statistic)

		storage.EXPECT().Delete(gomock.Any()).Return(nil)

		err := serv.DeleteUser(1)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, nil, err)
	})

	t.Run("Get statistic should return statistic from storage", func(t *testing.T) {
		statistic := models.Statistic{}
		storage.EXPECT().GetStatistics().Return(&statistic)

		result := serv.GetStatistics()

		assert.Equal(t, statistic, *result)
	})
}
