package main

import (
	"fmt"
	"log"
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
	 _, err := handler.Auth(&models.AddRequest{
		AuthParams: models.AuthData{
			Login: "nods1",
			Password: "nods123",
		},
	})

	if err != nil {
		log.Println(err.Error())
	}

	_, err = handler.Auth(&models.AddRequest{
		AuthParams: models.AuthData{
			Login: "nods2",
			Password: "nods123",
		},
	})

	if err != nil {
		log.Println(err.Error())
	}

	updateUser := handler.UpdateUser(1, &models.UserDate{
		Name: "Nods",
		Surname: "Sulaymonov",
		Status: "online",
		Role: "user",
		RegistrationDate: time.Time{},
		UpdateDate: time.Time{},
	})

	if updateUser != nil {
		fmt.Println(updateUser)
		return
	}

	_, err = handler.GetUser(1)

	if err != nil {
		log.Println(err.Error())
	} 

	err = handler.DeleteUser(1)

	if err != nil {
		log.Println(err)
	} 

	_, err = handler.GetUsers()

	if err != nil {
		log.Println(err.Error())
	} 

	handler.GetStatistics()
}