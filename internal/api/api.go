package api

import (
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
)

type IApi interface {
	Auth(s *models.SAuth)(int32,error)
	Add(user *models.SUser, id *models.IdGenerate)(bool,error)
	Update(id *models.IdGenerate,s *models.SAuth)error
	GetUser(id int32)*models.SUser
	GetUsers()
	DeleteUser(id int32)error
	GetStatistics()
}

type SApi struct {
	repo service.IService
}

// func(s *SApi)Add(user *models.Sjson, id *models.IdGenerate)(bool,error){

// }
