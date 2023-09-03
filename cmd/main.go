package main

import (
	"github.com/Shemistan/Lesson_6/api"
	"github.com/Shemistan/Lesson_6/api/dtos"
	"github.com/Shemistan/Lesson_6/models"
	"github.com/Shemistan/Lesson_6/services"
	"github.com/Shemistan/Lesson_6/storage"
	"log"
	"time"
)

const (
	host string = "127.0.0.1"
	port int    = 3232
	ttl  int    = 30
)

func main() {
	app := initApp()

	runCommands(app)
}

func runCommands(api api.IApi) {
	auth, err := api.Auth(dtos.AuthRequest{
		Login:    "denis01",
		Password: "12345",
	})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(auth)
	}

	auth2, err := api.Auth(dtos.AuthRequest{
		Login:    "denis02",
		Password: "12345",
	})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(auth2)
	}

	auth3, err := api.Auth(dtos.AuthRequest{
		Login:    "denis3",
		Password: "12345",
	})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(auth3)
	}

	update, err := api.UpdateUser(dtos.UpdateUserRequest{
		Id: 1,
		User: dtos.UserDto{
			Name:             "Daniel",
			Surname:          "Li",
			Status:           models.StatusBlocked,
			Role:             models.RoleAdmin,
			RegistrationDate: time.Time{},
			UpdateDate:       time.Time{},
		},
	})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(update)
	}

	getUser, err := api.GetUser(dtos.GetUserRequest{
		Id: 1,
	})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(getUser)
	}

	err = api.DeleteUser(dtos.DeleteUserRequest{
		Id: 1,
	})

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Success")
	}

	getUsers, err := api.GetUsers()

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(getUsers)
	}
}

func initApp() api.IApi {
	conn := storage.NewConn()
	appStorage := storage.New(host, port, ttl, conn)
	appService := services.New(appStorage)

	return api.New(appService)
}
