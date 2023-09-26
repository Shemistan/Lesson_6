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

func TestIntegrationRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockIStorage(ctrl)
	serv := New(storage)

	t.Run("Add user -> success", func(t *testing.T) {
		storage.EXPECT().Add(gomock.Any()).Return(int64(1), nil)

		fakeUserAuth := models.User{
			Login:            "July",
			HashedPassword:   "fshiuffhs554redyfsgf:reyfv",
			RegistrationDate: time.Now().String(),
		}

		id, _ := serv.Register(&fakeUserAuth)

		assert.Equal(t, int64(1), id)

	})

	t.Run("Add user -> nil", func(t *testing.T) {
		storage.EXPECT().Add(gomock.Any()).Return(int64(0), errors.New("user is nil"))

		_, err := serv.Register(nil)

		if err == nil {
			t.Error(err)
		}
	})

}

func TestIntegrationUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockIStorage(ctrl)
	serv := New(storage)

	fakeUpdatedUser := models.User{
		Firstname:   "John",
		Lastname:    "Doe",
		UpdatedDate: time.Now().String(),
	}

	t.Run("Update -> success", func(t *testing.T) {
		storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

		err := serv.UpdateUser(int64(1), &fakeUpdatedUser)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, nil, err)
	})

	t.Run("Update -> not found", func(t *testing.T) {
		storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("user with id: 1 not found"))

		err := serv.UpdateUser(int64(1), &fakeUpdatedUser)

		assert.Equal(t, "user with id: 1 not found", err.Error())
	})

	t.Run("Update -> ID < 0", func(t *testing.T) {
		storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("id should be greater than 0"))

		err := serv.UpdateUser(int64(-1), &fakeUpdatedUser)

		assert.Equal(t, "id should be greater than 0", err.Error())

	})
}

func TestIntegrationGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockIStorage(ctrl)
	serv := New(storage)

	fakeUser := models.User{
		Id:               2,
		Login:            "qwerty",
		Firstname:        "John",
		Lastname:         "Doe",
		Status:           "Active",
		HashedPassword:   "ksdjflsg8487465w3ur@kojfs125",
		RegistrationDate: time.Now().String(),
		UpdatedDate:      time.Now().String(),
	}

	t.Run("Get -> success", func(t *testing.T) {
		storage.EXPECT().Get(gomock.Any()).Return(&fakeUser, nil)

		user, _ := serv.GetUser(int64(10))

		assert.Equal(t, &fakeUser, user)
	})

	t.Run("Get -> not found", func(t *testing.T) {
		storage.EXPECT().Get(gomock.Any()).Return(&models.User{}, errors.New("user with id 10 doesn't exist"))

		_, err := serv.GetUser(int64(10))

		assert.Equal(t, "user with id 10 doesn't exist", err.Error())
	})

	t.Run("Get all", func(t *testing.T) {
		userList := []*models.User{&fakeUser, &fakeUser}
		storage.EXPECT().GetAll().Return(userList)

		users := serv.GetAllUsers()

		if len(users) == 0 {
			t.Error("something went wrong")
		}

		assert.Equal(t, len(userList), len(users))
	})
}

func TestIntegrationDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockIStorage(ctrl)
	serv := New(storage)

	t.Run("Delete -> success", func(t *testing.T) {
		storage.EXPECT().Delete(gomock.Any()).Return(int64(10), nil)

		id, _ := serv.DeleteUser(int64(10))

		assert.Equal(t, int64(10), id)
	})

	t.Run("Delete -> ID <= 0", func(t *testing.T) {
		storage.EXPECT().Delete(gomock.Any()).Return(int64(0), errors.New("invalid id"))

		_, err := serv.DeleteUser(int64(0))

		assert.Equal(t, "invalid id", err.Error())
	})

	t.Run("Delete -> ID doesn't exist", func(t *testing.T) {
		storage.EXPECT().Delete(gomock.Any()).Return(int64(0), errors.New("user with id 10 doesn't exist"))

		_, err := serv.DeleteUser(int64(0))

		assert.Equal(t, "user with id 10 doesn't exist", err.Error())
	})
}
