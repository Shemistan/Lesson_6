package api

import (
	"errors"
	"testing"

	"github.com/Shemistan/Lesson_6/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
	mock_service "github.com/Shemistan/Lesson_6/internal/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMockApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serv := mock_service.NewMockIService(ctrl)
	api := New(serv)

	t.Run("Auth", func(t *testing.T) {
		req1 := &models.AddRequest{
			AuthParams: models.UserAuth{
				Login:    "asd",
				Password: "asd1",
			},
			Data: models.UserData{
				Name:    "user1",
				Surname: "asdasd",
			},
		}

		req2 := &models.AddRequest{
			AuthParams: models.UserAuth{
				Login:    "login",
				Password: "pass",
			},
			Data: models.UserData{
				Name:    "user2",
				Surname: "userov",
			},
		}

		upd := &models.UpdateUserData{
			Name: "newName",
		}
		serv.EXPECT().GetUserByLogin("asd", "asd1").Return(int64(0), errors.New("Such user doesn't exist"))
		serv.EXPECT().Add(converters.ApiToServiceModel(req1)).Return(int64(1), nil)
		serv.EXPECT().GetUserByLogin("asd", "asd1").Return(int64(1), nil)

		serv.EXPECT().GetUserByLogin("login", "pass").Return(int64(0), errors.New("Such user doesn't exist"))
		serv.EXPECT().Add(converters.ApiToServiceModel(req2)).Return(int64(2), nil)

		serv.EXPECT().UpdateUser(int64(2), upd).Return(nil)

		serv.EXPECT().GetUser(int64(3)).Return(nil, errors.New("Such user doesn't exist"))

		serv.EXPECT().GetUsers().Return(nil, errors.New("Db is empty"))

		serv.EXPECT().DeleteUser(int64(2)).Return(nil)

		res, _ := api.Auth(req1)
		assert.Equal(t, int64(1), res)

		_, err := api.Auth(req1)
		assert.Equal(t, nil, err)

		res, _ = api.Auth(req2)
		assert.Equal(t, int64(2), res)

		err = api.UpdateUser(int64(2), upd)
		assert.Equal(t, nil, err)

		_, err = api.GetUser(int64(3))
		assert.Equal(t, errors.New("Such user doesn't exist"), err)

		_, err = api.GetUsers()
		assert.Equal(t, errors.New("Db is empty"), err)

		err = api.DeleteUser(2)
		assert.Equal(t, nil, err)
	})
}
