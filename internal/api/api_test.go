package api

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockIStorage(ctrl)
	serv := service.New(storage)
	api := New(serv)

	t.Run("Get all users", func(t *testing.T) {
		storage.EXPECT().GetUsers().Return([]*models.User{}, nil)

		userList, err := api.GetUsers()

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, []*models.User{}, userList)

	})

	userLogin := &models.AddRequest {
		AuthParams: models.AuthData{
			Login: "nods",
			Password: "dasmkdmak12231",
		},
	}

	t.Run("Auth return error ", func(t *testing.T) {
		storage.EXPECT().Auth(gomock.Any()).Return(0, errors.New("some errors"))

		_, err := api.Auth(userLogin)

		assert.Equal(t, errors.New("some errors"), err)

	})

	t.Run("Auth return success ", func(t *testing.T) {
		storage.EXPECT().Auth(gomock.Any()).Return(1, nil)


		res, _ := api.Auth(userLogin)

		assert.Equal(t, 1, res)

	})

	t.Run("GetUser return error ", func(t *testing.T) {
		storage.EXPECT().GetUser(gomock.Any()).Return(&models.User{}, errors.New("some errors"))

		_, err := api.GetUser(0)

		assert.Equal(t, errors.New("some errors"), err)

	})

	t.Run("GetUser return success ", func(t *testing.T) {

		testUser := models.User{
			Id:               1,
			Login:            "nods",
			HashPassword:     "dmsafmsa2f3",
			Name:             "Nodir",
			Surname:          "Sulaymonov",
			Status:           "active",
			Role:             "user",
			RegistrationDate: time.Time{},
			UpdateDate:       time.Time{},
		}

		storage.EXPECT().GetUser(gomock.Any()).Return(&testUser, nil)

		res, _ := api.GetUser(1)

		assert.Equal(t, &testUser, res)

	})

	t.Run("UpdateUser return error ", func(t *testing.T) {
		storage.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New("some errors"))

		res := api.UpdateUser(0, nil)

		assert.Equal(t, errors.New("some errors"), res)

	})

	t.Run("UpdateUser return success ", func(t *testing.T) {
		storage.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)

		res := api.UpdateUser(1, nil)

		assert.Equal(t, nil, res)

	})

	t.Run("DeleteUser return error", func(t *testing.T) {
		storage.EXPECT().DeleteUser(gomock.Any()).Return(errors.New("some errors"))

		err := api.DeleteUser(0)

		assert.Equal(t, errors.New("some errors"), err)
	})

	t.Run("DeleteUser return success", func(t *testing.T) {
		storage.EXPECT().DeleteUser(gomock.Any()).Return(nil)

		err := api.DeleteUser(1)

		assert.Equal(t, nil, err)
	})


}
	