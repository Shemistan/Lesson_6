package converters

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/storage/models"
)

func UserDtoToModel(dto dtos.UserDto) models.User {
	return models.User{
		Id:               dto.Id,
		Login:            dto.Login,
		Name:             dto.Name,
		Surname:          dto.Surname,
		Status:           dto.Status,
		Role:             dto.Role,
		RegistrationDate: dto.RegistrationDate,
		UpdateDate:       dto.UpdateDate,
	}
}
