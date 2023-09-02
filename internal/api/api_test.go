package api

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/Lesson_6/internal/models"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
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
		input *models.AuthRequest
		resp  int
		error error
	}{
		{input: &models.AuthRequest{
			Login:    "Test",
			Password: "Test",
		},
			resp:  1,
			error: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("tc.input.Login", func(t *testing.T) {
			storage.EXPECT().Auth(tc.input).Return(tc.resp, tc.error)
			res, err := api.Auth(tc.input)
			assert.Equal(t, tc.resp, res)
			assert.Equal(t, tc.error, err)
		})
	}

}
