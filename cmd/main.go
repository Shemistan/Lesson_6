package main

import (
	"github.com/Shemistan/Lesson_6/internal/api"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

func main() {
	storage := storage.New("localhost", 5432)
	service := service.New(storage)
	api := api.New(service)
	Server(api)
}


func Server(handler api.IApi) {
	handler.Register(&models.Request{
		AuthParams: models.AuthParams{
			Login: "July",
			Password: "142536sa",
		},
	})

	handler.Get(1)

	handler.Update(1, &models.UserUpdateRequest{
		Firstname: "Bogdan",
		Lastname: "Azimjanov",
	})

	handler.Get(1)

	handler.Update(1, &models.UserUpdateRequest{
		Firstname: "",
		Lastname: "Azimzhanov",
	})

	handler.Get(1)

	handler.GetAllUsers()

	handler.Register(&models.Request{
		AuthParams: models.AuthParams{
			Login:    "Jo",
			Password: "12345678",
		},
	})

	handler.Get(2)

	handler.GetAllUsers()

	handler.Update(2, &models.UserUpdateRequest{
		Firstname: "John",
		Lastname: "Doe",
	})

	handler.Get(2)

	handler.DeleteUser(2)

	handler.Get(2)

	handler.GetAllUsers()
}