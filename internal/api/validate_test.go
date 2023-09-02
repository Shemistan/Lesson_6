package api

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/Lesson_6/internal/models"
)

func TestValidateAuthRequest(t *testing.T) {
	tests := []struct {
		input *models.AuthRequest
		error error
	}{
		{input: nil,
			error: errors.New("request is nil"),
		},
		{input: &models.AuthRequest{
			Login:    "Test",
			Password: "Test",
		},
			error: nil,
		},
		{input: &models.AuthRequest{
			Login:    "",
			Password: "Test",
		},
			error: errors.New("login is empty"),
		},
		{input: &models.AuthRequest{
			Login:    "Test",
			Password: "",
		},
			error: errors.New("password is empty"),
		},
		{input: &models.AuthRequest{
			Login:    "",
			Password: "",
		},
			error: errors.New("login and password is empty"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("validate auth fields", func(t *testing.T) {
			err := ValidateAuthRequest(tc.input)
			assert.Equal(t, tc.error, err)
		})
	}
}

func TestValidateUpdateUser(t *testing.T) {
	tests := []struct {
		input *models.UserRequest
		id    int
		error error
	}{
		{input: nil,
			id:    1,
			error: errors.New("request is nil"),
		},
		{input: &models.UserRequest{
			Name:             "Mike",
			Surname:          "Snow",
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
			id:    0,
			error: errors.New("no user with id equal to 0"),
		},
		{input: &models.UserRequest{
			Name:             "Mike",
			Surname:          "Snow",
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
			id:    1,
			error: nil,
		},
		{input: &models.UserRequest{
			Name:             "",
			Surname:          "Snow",
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
			id:    1,
			error: errors.New("name is empty"),
		},
		{input: &models.UserRequest{
			Name:             "Mike",
			Surname:          "",
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
			id:    1,
			error: errors.New("surname is empty"),
		},
		{input: &models.UserRequest{
			Name:             "Mike",
			Surname:          "Snow",
			Status:           "",
			Role:             "User",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
			id:    1,
			error: errors.New("status is empty"),
		},
		{input: &models.UserRequest{
			Name:             "Mike",
			Surname:          "Snow",
			Status:           "Active",
			Role:             "",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
			id:    1,
			error: errors.New("role is empty"),
		},
		{input: &models.UserRequest{
			Name:             "Mike",
			Surname:          "Snow",
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Time{},
			UpdateDate:       time.Now(),
		},
			id:    1,
			error: errors.New("registrationDate is empty"),
		},
		{input: &models.UserRequest{
			Name:             "Mike",
			Surname:          "Snow",
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Time{},
		},
			id:    1,
			error: errors.New("updateDate is empty"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run("validate update user fields", func(t *testing.T) {
			err := ValidateUpdateUser(tc.id, tc.input)
			assert.Equal(t, tc.error, err)
		})
	}
}
