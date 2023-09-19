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
		Login:    "manager_1",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	if err != nil {
		log.Fatalln("cannot add user")
	}

	_, err = handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_2",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	if err != nil {
		log.Fatalln("cannot add user")
	}
	user, err := handlers.GetUser(&models.GetUserRequest{
		UserId: 1,
	})

	if err != nil {
		log.Fatalln("cannot get user")
	}

	fmt.Printf("%v", user)

	err = handlers.UpdateUser(&models.UpdateUserRequest{
		Id:       int(user.Id),
		Login:    "updated_manager",
		Password: "12345678",
		Name:     "Daler_",
		Surname:  "Sultonov_",
	})

	if err != nil {
		log.Fatalln("cannot update user")
	}

	err = handlers.DeleteUser(&models.DeleteUserRequest{
		Id: 1,
	})

	if err != nil {
		log.Fatalln("cannot delete user")
	}

	_, err = handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_1",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	if err != nil {
		log.Fatalln("cannot add user")
	}

	_, err = handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_1",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	if err != nil {
		log.Fatalln("cannot add user")
	}

	_, err = handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_1",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	if err != nil {
		log.Fatalln("cannot add user")
	}

	users := handlers.GetUsers(&models.GetUsersRequest{})
	fmt.Printf("%v", users)

	statistics := handlers.GetStatustics()
	fmt.Println(statistics)
}
