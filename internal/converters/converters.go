package converters

import (
	"github.com/Shemistan/Lesson_6/internal/models"
	"time"
)

func ApiAuthModelToServiceUserModel(id int, req models.AuthRequest) *models.User {
	res := &models.User{
		ID:               id,
		Login:            req.Login,
		Name:             req.Login,
		Surname:          req.Login,
		Status:           "active",
		Role:             "user",
		RegistrationDate: time.Now(),
		UpdateDate:       time.Now(),
	}
	return res
}
