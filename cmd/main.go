package main

import (
	"fmt"
	"golang-api/internal/api"
	"golang-api/internal/models"
	"golang-api/internal/service"
	"golang-api/internal/storage"
	"log"
)

func main() {
	st := storage.New()
	srv := service.New(st)
	handlers := api.New(srv)

	Server(handlers)
}

func Server(handlers api.IApi) {
	_, err := handlers.AddUser(&models.AddUserRequest{
		Login:    "user1",
		Password: "123456",
		Name:     "Bekzod",
		Surname:  "Askarov",
	})

	if err != nil {
		log.Fatalln("Cannot add user")
	}

	_, err = handlers.AddUser(&models.AddUserRequest{
		Login:    "user2",
		Password: "123456",
		Name:     "Doe",
		Surname:  "Jones",
	})

	if err != nil {
		log.Fatalln("Cannot add user")
	}
	user, err := handlers.GetUser(&models.GetUserRequest{
		UserId: 1,
	})

	if err != nil {
		log.Fatalln("Cannot get user")
	}

	fmt.Printf("%v", user)

	err = handlers.UpdateUser(&models.UpdateUserRequest{
		Id:       int(user.Id),
		Login:    "updated_user",
		Password: "12345678",
		Name:     "Bekzod_",
		Surname:  "Askarov_",
	})

	if err != nil {
		log.Fatalln("Cannot update user")
	}

	err = handlers.DeleteUser(&models.DeleteUserRequest{
		Id: 1,
	})

	if err != nil {
		log.Fatalln("Cannot delete user")
	}

	_, err = handlers.AddUser(&models.AddUserRequest{
		Login:    "user1",
		Password: "123456",
		Name:     "Bekzod",
		Surname:  "Askarov",
	})

	if err != nil {
		log.Fatalln("Cannot add user")
	}

	_, err = handlers.AddUser(&models.AddUserRequest{
		Login:    "user1",
		Password: "123456",
		Name:     "Bekzod",
		Surname:  "Askarov",
	})

	if err != nil {
		log.Fatalln("Cannot add user")
	}

	_, err = handlers.AddUser(&models.AddUserRequest{
		Login:    "user1",
		Password: "123456",
		Name:     "Bekzod",
		Surname:  "Askarov",
	})

	if err != nil {
		log.Fatalln("Cannot add user")
	}

	users := handlers.GetUsers(&models.GetUsersRequest{})
	fmt.Printf("%v", users)

	statistics := handlers.GetStatustics()
	fmt.Println(statistics)
}
