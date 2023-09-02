package converters

import (
	"time"

	"github.com/Shemistan/Lesson_6/internal/models"
)

func ApiAuthModelToServiceUserModel(req models.AuthRequest) *models.User {
	return &models.User{
		ID:               0,
		Login:            req.Login,
		Name:             req.Login,
		Surname:          req.Login,
		Status:           "active",
		Role:             "user",
		RegistrationDate: time.Now(),
		UpdateDate:       time.Now(),
	}
}
