package converters

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/storage/models"
)

func AuthToUserModel(dto dtos.AuthRequest) models.User {
	return models.User{
		Name:     dto.Login,
		Password: dto.Password,
		Status:   models.StatusActive,
		Role:     models.RoleUser,
	}
}
