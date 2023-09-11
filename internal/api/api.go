package api

import (
	"github.com/Shemistan/Lesson_6/internal/converters"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
)

type IApi interface{
	Auth(req *models.AddRequest)(int, error)
}

func New(serv service.IService)IApi {
	return &api{serv: serv}
}

type api struct {
	serv service.IService
}

func (a *api) Auth(req *models.AddRequest)(int, error) {
	res, err := a.serv.Auth(converters.ApiAuthModelToServiceUserModel(*req))
	if err != nil {
		return 0, err
	}

	return res, nil
}