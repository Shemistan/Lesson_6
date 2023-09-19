package api

import (
	"errors"
	"testing"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockIStorage(ctrl)
	serv := service.New(storage)
	api := New(serv)

	tests := []TestCaseRegister{
		{
			name: "success",
			request: &models.Request{
				AuthParams: models.AuthParams{
					Login:    "Doe",
					Password: "123456789",
				},
			},
			response: 1,
			err:      nil,
		},
		{
			name: "fail",
			request: &models.Request{
				AuthParams: models.AuthParams{
					Login:    "John",
					Password: "qwerty987456321",
				},
			},
			response: 0,
			err:      errors.New("error text"),
		},
	}

	testsUpdate := []TestCaseUpdate{
		{
			name: "Update success",
			request: &models.UserUpdateRequest{
				Firstname: "John",
				Lastname:  "Doe",
			},
			response: nil,
			id:       1,
		},
		{
			name: "Update fail",
			id:   0,
			request: &models.UserUpdateRequest{
				Firstname: "",
				Lastname:  "",
			},
			response: errors.New("error text"),
		},
	}

	testsGet := []TestCaseGet {
		{
			name:     "Get user -> OK",
			id:       1,
			response: &models.User{
				Id:               1,
				Login:            "do",
				Firstname:        "John",
				Lastname:         "Doe",
				Status:           "Active",
				HashedPassword: "adshabsb2833465dhasd@#23shfsz",
				RegistrationDate: "19-09-2023",
				UpdatedDate:      "",
			},
			err:      nil,
		},
		{
			name:     "Get user -> FAIL",
			id:       0,
			response: &models.User{},
			err:      errors.New("error text"),
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			storage.EXPECT().Add(gomock.Any()).Return(c.response, c.err)
			res, err := api.Register(c.request)
			assert.Equal(t, c.response, res)
			assert.Equal(t, c.err, err)
		})
	}

	for _, c := range testsUpdate {
		t.Run(c.name, func(t *testing.T) {
			storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(c.response)

			resp := api.Update(int64(c.id), c.request)

			assert.Equal(t, resp, c.response)
		})
	}

	for _, c := range testsGet {
		t.Run(c.name, func(t *testing.T) {
			storage.EXPECT().Get(gomock.Any()).Return(c.response, c.err)

			resp, err := api.Get(int64(c.id))

			assert.Equal(t, c.response, resp)
			assert.Equal(t, c.err, err)
		})
	}

	t.Run("Get all users", func(t *testing.T) {
		storage.EXPECT().GetAll().Return([]*models.User{})

		res := api.GetAllUsers()

		assert.Equal(t, []*models.User{}, res)
	})

	t.Run("Delete -> OK", func(t *testing.T) {
		storage.EXPECT().Delete(gomock.Any()).Return(int64(1), nil)

		id, _ := api.DeleteUser(int64(1))

		assert.Equal(t, int64(1), id)
	})

	t.Run("Delete -> FAIL", func(t *testing.T) {
		storage.EXPECT().Delete(gomock.Any()).Return(int64(0), errors.New("error text"))

		_, err := api.DeleteUser(int64(0))

		assert.Equal(t, errors.New("error text"), err)
	})

}

type TestCaseRegister struct {
	name     string
	request  *models.Request
	response int64
	err      error
}

type TestCaseUpdate struct {
	name     string
	id       int64
	request  *models.UserUpdateRequest
	response error
}

type TestCaseGet struct {
	name string
	id int64
	response *models.User
	err error
}

