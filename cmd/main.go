package main

import (
	"time"

	"github.com/Shemistan/Lesson_6/internal/api"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

func main() {
	conn := storage.NewConnect()
	db := storage.New("localhost", 5432, 30, conn)
	serv := service.New(db)
	handler := api.New(serv)
	Server(handler)
}

func Server(handler api.IApi) {
	handler.Auth(&models.AddRequest{
		AuthParams: models.AuthData{
			Login: "login_1",
			Password: "password_1",
		},
		Date: models.UserDate{
			Name: "Alex",
			Surname: "Florida",
			Role: "User",
			Status: "Tourist",
			RegistrationDate: time.Now(),
			UpdateDate:       time.Now(),
		},
	})
}