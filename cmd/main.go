package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_6/internal/api"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
	"github.com/Shemistan/Lesson_6/internal/storage"
	"time"
)

func main() {
	// Читай конфиг файл и вытаскивам различные данные
	conn := storage.NewConnect()
	db := storage.New("localhost", 5432, 30, conn)

	serv := service.New(db)
	handlers := api.New(serv)

	Server(handlers)
}

func Server(h api.IApi) {
	_, err := h.Auth(&models.AuthRequest{
		Login:    "login_1",
		Password: "password_1",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = h.Auth(&models.AuthRequest{
		Login:    "login_1",
		Password: "password_1",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = h.Auth(&models.AuthRequest{
		Login:    "login_2",
		Password: "password_2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	errUpdate := h.UpdateUser(1, &models.UserRequest{
		Name:             "UpdateName",
		Surname:          "UpdateName",
		Status:           "active",
		Role:             "user",
		RegistrationDate: time.Now(),
		UpdateDate:       time.Now(),
	})
	if errUpdate != nil {
		fmt.Println("errUpdate")
		fmt.Println(errUpdate)
		return
	}

	_, errGetUser := h.GetUser(2)
	if errGetUser != nil {
		fmt.Println("errGetUser")
		fmt.Println(errGetUser)
		return
	}

	_, errGetUsers := h.GetUsers()
	if errGetUsers != nil {
		fmt.Println("errGetUsers")
		fmt.Println(errGetUsers)
		return
	}

	errDeleteUser := h.DeleteUser(2)
	if errDeleteUser != nil {
		fmt.Println("errDeleteUser")
		fmt.Println(errDeleteUser)
		return
	}

	_, errGetUsersAfterDelete := h.GetUsers()
	if errGetUsersAfterDelete != nil {
		fmt.Println("errGetUsersAfterDelete")
		fmt.Println(errGetUsersAfterDelete)
		return
	}

	h.GetStatistics()

}
