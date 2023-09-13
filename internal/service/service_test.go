package service

import (
	"errors"
	"testing"

	"github.com/Shemistan/Lesson_6/internal/models"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock_storage.NewMockIStorage(ctrl)
	serv := New(storage)

	fakeUser := models.User{
		Id: 0,
		Login: "july",
		Firstname: "Bogdan",
		Lastname: "Azimjanov",
		Status: "Active",
		HashedPassword: "142536sa",
		RegistrationDate: "",
		UpdatedDate: "",
	}

	t.Run("Nil user", func(t *testing.T) {
		storage.EXPECT().Add(nil).Return(int64(0), errors.New("user is null"))

		_, err := serv.Register(nil)

		if assert.Error(t, err) {
			assert.Equal(t, "user is null", err.Error())
		}

	})

	t.Run("Create user", func(t *testing.T) {
		storage.EXPECT().Add(&fakeUser).Return(fakeUser.Id, nil)

		id, err := serv.Register(&fakeUser)

		if err != nil {
			t.Error("should return ID, returned error")
		}

		assert.Equal(t, int64(1), id)
	})
}


