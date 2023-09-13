package api

import (
	"errors"
	"fmt"

	"github.com/Shemistan/Lesson_6/internal/convert"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
)

type IApi interface {
	Auth(s string)(string,error)
	Add(s string)(bool,error)
	Update(id int32, str string)error
	GetUser(id string)(string)
	GetUsers()
	DeleteUser(id string)error
	GetStatistics()
}

type SApi struct {
	repo service.IService
}

func NewIApi(repo service.IService) IApi{
	return &SApi{repo: repo}
}

func (s *SApi)GetStatistics(){
	s.repo.GetStatistics()
}

func(s *SApi)DeleteUser(idjson string)error{
	// id :=models.IdGenerate{}
	id,_ := convert.ApiIdConvertToService(idjson)
	idg := id.Id
	err := s.repo.DeleteUser(idg)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func(s *SApi)GetUsers(){
	mapa := s.repo.GetMap()
	for _, value := range mapa{
		value.PasswordHash = ""
		fmt.Println(convert.ApiUserConvertFromoService(*value))
	}
}

func(s *SApi)GetUser(idjson string)(string){
	// update1 := models.SUser{}
	// id :=models.IdGenerate{}
	id,_ := convert.ApiIdConvertToService(idjson)
	idg := id.Id
	user1, _  := s.repo.GetUser(idg)
	str := convert.ApiUserConvertFromoService(*user1)
	return str
}

func(s *SApi)Update(id int32,str string)error{
	// update1 := models.SUser{}
	update1, err := convert.ApiUserConvertToService(str)
	if err != nil {
		fmt.Println(err)
	}else{
		_, err := s.repo.UpdateUser(id, &update1)
		if err != nil {
			return errors.New("Error Cant Update user")
		}
		return nil
	}
	return errors.New("Error Cant Update user")
}

func(s *SApi)Add(str string)(bool,error){
	// user := models.SUser{}
	user, err := convert.ApiUserConvertToService(str)
	if err != nil {
		fmt.Println(err)
	}else{
		_,err := s.repo.Add(&user)
		if err != nil {
			return false, errors.New("Error Json Add to service")
		}
		return true, nil
	}
	return false, errors.New("Error Json Add to service")
}

func (s *SApi)Auth(jauth string)(string,error){
	// user1 := models.SAuth{}
	// test := convert.ApiAuthConvertFromoService(user)
	// fmt.Println("JSON-->", test)
	user1 := convert.ApiAuthConvertToService(jauth)
	var str string
	if user1.Login != "" {
		idg, err := s.repo.Auth(&user1)
		id := models.IdGenerate{}
		id.Id = idg
		str = convert.ApiIdConvertFromoService(id)
		if err != nil {
			fmt.Println("kakayato oshibka", err)
		}
		return str, nil
	}
	return str, nil
}
