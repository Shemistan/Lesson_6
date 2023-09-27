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

type api struct {
	repo service.IService
}

func NewIApi(repo service.IService) IApi{
	return &api{repo: repo}
}

func (s *api)GetStatistics(){
	s.repo.GetStatistics()
}

func(s *api)DeleteUser(idjson string)error{
	id,_ := convert.ApiIdConvertToService(idjson)
	idg := id.Id
	err := s.repo.DeleteUser(idg)
	
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func(s *api)GetUsers(){
	mapa := s.repo.GetMap()

	for _, value := range mapa{
		value.PasswordHash = ""
		fmt.Println(convert.ApiUserConvertFromoService(*value))
	}
}

func(s *api)GetUser(idjson string)(string){
	id,_ := convert.ApiIdConvertToService(idjson)
	idg := id.Id
	user1, _  := s.repo.GetUser(idg)
	str := convert.ApiUserConvertFromoService(*user1)

	return str
}

func(s *api)Update(id int32,str string)error{
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

func(s *api)Add(str string)(bool,error){
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

func (s *api)Auth(jauth string)(string,error){
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
