package service

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"golang-api/internal/models"
	mock_storage "golang-api/internal/storage/mocks"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockIStorage(ctrl)
	serv := New(storage)

	t.Run("Successful Auth ", func(t *testing.T) {
		storage.EXPECT().Auth(gomock.Any()).Return(1, nil)

		userId, err := serv.Auth(nil)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 1, userId)
	})

	t.Run("Error  Auth", func(t *testing.T) {
		storage.EXPECT().Auth(gomock.Any()).Return(1, errors.New("user not add"))

		_, err := serv.Auth(nil)
		if err == nil {
			t.Error(err)
		}
	})

	t.Run("Successful GetUser", func(t *testing.T) {
		testUser := models.User{
			Id:               0,
			Login:            "user_1",
			HashPassword:     "jwtsha1sja256",
			Name:             "Faxa",
			Surname:          "Faxa",
			Status:           "active",
			Role:             "user",
			RegistrationDate: time.Time{},
			UpdateDate:       time.Time{},
		}

		storage.EXPECT().GetUser(gomock.Any()).Return(&testUser, nil)

		user, err := serv.GetUser(1)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, &testUser, user)
	})

	t.Run("Error GetUser", func(t *testing.T) {
		storage.EXPECT().GetUser(gomock.Any()).Return(&models.User{}, errors.New("some errors"))

		_, err := serv.GetUser(1)

		assert.Equal(t, errors.New("some errors"), err)
	})

	t.Run("Successful GetUsers", func(t *testing.T) {
		var testUserList []*models.User
		storage.EXPECT().GetUsers().Return(testUserList, nil)

		userList, err := serv.GetUsers()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, testUserList, userList)
	})

	t.Run("Error GetUsers", func(t *testing.T) {
		var testUserList []*models.User
		storage.EXPECT().GetUsers().Return(testUserList, errors.New("some errors"))

		_, err := serv.GetUsers()

		assert.Equal(t, errors.New("some errors"), err)
	})

	t.Run("Successful UpdateUser", func(t *testing.T) {
		userTest := &models.UserDate{
			Name:             "Faxa",
			Surname:          "Faxa",
			Status:           "online",
			Role:             "user",
			RegistrationDate: time.Time{},
			UpdateDate:       time.Time{},
		}

		storage.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)

		err := serv.UpdateUser(1, userTest)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, nil, err)
	})

	t.Run("Error UpdateUser", func(t *testing.T) {
		storage.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New("some errors"))

		err := serv.UpdateUser(1, nil)

		assert.Equal(t, "some errors", err.Error())
	})

	t.Run("Successful DeleteUser", func(t *testing.T) {
		storage.EXPECT().DeleteUser(gomock.Any()).Return(nil)

		err := serv.DeleteUser(1)

		assert.Equal(t, nil, err)
	})

	t.Run("Error DeleteUser", func(t *testing.T) {
		storage.EXPECT().DeleteUser(gomock.Any()).Return(errors.New("some errors"))

		err := serv.DeleteUser(1)

		assert.Equal(t, "some errors", err.Error())
	})

}
