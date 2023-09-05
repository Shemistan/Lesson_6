package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_6/controller"
	"github.com/Shemistan/Lesson_6/service"
	"github.com/Shemistan/Lesson_6/storage"
)

func Init() *controller.UserController {
	db := storage.NewStorage()
	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)
	return userController
}

func main() {
	userController := Init()

	var login, password string

	fmt.Print("Введите логин: ")
	fmt.Scan(&login)
	fmt.Print("Введите пароль: ")
	fmt.Scan(&password)
	userId, err := userController.Auth("Ilxom", "some_pass")
	if err != nil {
		panic(err)
	}

	user, err := userController.Get(userId)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	users, err := userController.GetAll()

	for _, user := range users {
		fmt.Println(user)
	}

	err = userController.Delete(0)
}
