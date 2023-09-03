package converters

import (
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/models"
)

func UserModelToDto(model models.User) dtos.UserDto {
	return dtos.UserDto{
		Id:               model.Id,
		Login:            model.Login,
		Name:             model.Name,
		Surname:          model.Surname,
		Status:           model.Status,
		Role:             model.Role,
		RegistrationDate: model.RegistrationDate,
		UpdateDate:       model.UpdateDate,
	}
}
