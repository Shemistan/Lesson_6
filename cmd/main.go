package main

import (
	"fmt"
	"log"
	"time"

	"golang-api/internal/api"
	"golang-api/internal/models"
	"golang-api/internal/service"
	"golang-api/internal/storage"
)

func main() {
	conn := storage.NewConnect()
	db := storage.New("localhost", 8080, 30, conn)
	serv := service.New(db)
	handler := api.New(serv)
	Server(handler)
}

func Server(handler api.IApi) {
	_, err := handler.Auth(&models.AddRequest{
		AuthParams: models.AuthData{
			Login:    "user_1",
			Password: "user1_pass",
		},
	})

	if err != nil {
		log.Println(err.Error())
	}

	_, err = handler.Auth(&models.AddRequest{
		AuthParams: models.AuthData{
			Login:    "user1",
			Password: "user1_pass",
		},
	})

	if err != nil {
		log.Println(err.Error())
	}

	updateUser := handler.UpdateUser(1, &models.UserDate{
		Name:             "Faxa",
		Surname:          "Faxa",
		Status:           "online",
		Role:             "user",
		RegistrationDate: time.Time{},
		UpdateDate:       time.Time{},
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
