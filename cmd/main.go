package main

import (
	"fmt"

	"golang-api/internal/api"
	"golang-api/internal/models"
	"golang-api/internal/service"
	"golang-api/internal/storage"
)

func main() {
	storage := storage.New("localhost", 5432)
	service := service.New(storage)
	api := api.New(service)
	Server(api)
}


func Server(handler api.IApi) {
	_, err := handler.Register(&models.Request{
		AuthParams: models.AuthParams{
			Login: "July",
			Password: "142536sa",
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = handler.Get(1)
	if err != nil {
		fmt.Println(err)
	}

	handler.GetStats()
	
	err = handler.Update(1, &models.UserUpdateRequest{
		Firstname: "Bogdan",
		Lastname: "Azimjanov",
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = handler.Get(1)
	if err != nil {
		fmt.Println(err)
	}

	err = handler.Update(1, &models.UserUpdateRequest{
		Firstname: "",
		Lastname: "Azimzhanov",
	})
	if err != nil {
		fmt.Println(err)
	}
	
	_, err = handler.Get(1)
	if err != nil {
		fmt.Println(err)
	}

	handler.GetAllUsers()

	_, err = handler.Register(&models.Request{
		AuthParams: models.AuthParams{
			Login:    "Jo",
			Password: "12345678",
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	handler.GetStats()

	_, err = handler.Get(2)
	if err != nil {
		fmt.Println(err)
	}

	handler.GetAllUsers()

	err = handler.Update(2, &models.UserUpdateRequest{
		Firstname: "John",
		Lastname: "Doe",
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = handler.Get(2)
	if err != nil {
		fmt.Println(err)
	}

	_, err = handler.DeleteUser(2)
	if err != nil {
		fmt.Println(err)
	}

	_, err = handler.Get(2)
	if err != nil {
		fmt.Println(err)
	}

	handler.GetAllUsers()
	handler.GetStats()
}