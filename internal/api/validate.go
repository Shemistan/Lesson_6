package api

import (
	"errors"

	"github.com/Shemistan/Lesson_6/internal/models"
)

func ValidateAuthRequest(req *models.AuthRequest) error {
	switch {
	case req == nil:
		return errors.New("request is nil")
	case req.Login == "" && req.Password == "":
		return errors.New("login and password is empty")
	case req.Login == "":
		return errors.New("login is empty")
	case req.Password == "":
		return errors.New("password is empty")
	default:
		return nil
	}
}

func ValidateUpdateUser(id int, req *models.UserRequest) error {
	switch {
	case req == nil:
		return errors.New("request is nil")
	case id == 0:
		return errors.New("no user with id equal to 0")
	case req.Name == "":
		return errors.New("name is empty")
	case req.Surname == "":
		return errors.New("surname is empty")
	case req.Role == "":
		return errors.New("role is empty")
	case req.Status == "":
		return errors.New("status is empty")
	case req.RegistrationDate.IsZero():
		return errors.New("registrationDate is empty")
	case req.UpdateDate.IsZero():
		return errors.New("updateDate is empty")
	default:
		return nil
	}
}
