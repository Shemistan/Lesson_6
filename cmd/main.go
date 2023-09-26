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
	authFirst, err := handler.Auth(&models.AddRequest{
		AuthParams: models.AuthData{
			Login: "nods1",
			Password: "nods123",
		},
	})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(authFirst)
	}

	authSecond, err := handler.Auth(&models.AddRequest{
		AuthParams: models.AuthData{
			Login: "nods2",
			Password: "nods123",
		},
	})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(authSecond)
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

	user, err := handler.GetUser(1)

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Get info about", user)
	}

	deleteUser := handler.DeleteUser(1)

	if deleteUser != nil {
		log.Println(err.Error())
	} else {
		log.Println("User deleted success")
	}

	getUsers, err := handler.GetUsers()

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Get all users:", getUsers)
	}

	handler.GetStatistics()
}