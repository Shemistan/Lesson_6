package api

import (
	"fmt"

	"github.com/Shemistan/Lesson_6/internal/convert"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
)

type IApi interface {
	Auth(s string)(int32,error)
	// Add(user *models.SUser)(bool,error)
	// Update(id *models.IdGenerate,s *models.SAuth)error
	// GetUser(id int32)*models.SUser
	// GetUsers()
	// DeleteUser(id int32)error
	// GetStatistics()
}

type SApi struct {
	repo service.IService
}

func NewIApi(repo service.IService) IApi{
	return &SApi{repo: repo}
}

func (s *SApi)Auth(jauth string)(int32,error){
	user1 := models.SAuth{}
	// test := convert.ApiAuthConvertFromoService(user)
	// fmt.Println("JSON-->", test)
	user1 = convert.ApiAuthConvertToService(jauth)
	if user1.Login != "" {
		idg, err := s.repo.Auth(&user1)
		if err != nil {
			fmt.Println("kakayato oshibka", err)
		}
		return idg, nil
	}
	return 0,nil
}

// func(s *SApi)Add(user *models.Sjson)(bool,error){

// }
