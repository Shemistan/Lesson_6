package main

import (
	"fmt"
	"time"

	"github.com/nasrullonurullaev5/Lesson_6/internal/controller"
	"github.com/nasrullonurullaev5/Lesson_6/internal/model"
	"github.com/nasrullonurullaev5/Lesson_6/internal/service"
	"github.com/nasrullonurullaev5/Lesson_6/internal/storage"
)

func initializeUserController() *controller.UserController {
	db := storage.NewDataStorage() // Changed to NewDataStorage
	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)
	return userController
}

func getInput(message string, value interface{}) {
	fmt.Print(message)
	_, err := fmt.Scan(value)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	userController := initializeUserController()

	var name, surname, login, password string
	var choice, userID uint32
	var isAuthenticated bool

	fmt.Println("[1] Authentication\n[2] Registration\n[0] Exit")
	getInput("Choose an action: ", &choice)

start:

	if choice == 2 {
		fmt.Println("\t\tRegistration")
		getInput("Enter name: ", &name)
		getInput("Enter surname: ", &surname)
		if login == "" || password == "" {
			getInput("Enter login: ", &login)
			getInput("Enter password: ", &password)
		}
		id, err := userController.AddUser(name, surname, login, password) // Changed to AddUser
		if err != nil {
			fmt.Println(err)
		}
		userID = id
		isAuthenticated = true
	} else if choice == 1 {
		fmt.Println("\t\tAuthentication")
		getInput("Login: ", &login)
		getInput("Password: ", &password)

		userId, err := userController.Authenticate(login, password) // Changed to Authenticate
		if err != nil {
			if userId == 0 {
				fmt.Println(err)
			}
			if userId == -1 {
				fmt.Println("User does not exist\nCreating a new user")
				choice = 2
				goto start
			}
		}
		isAuthenticated = true
	} else if choice == 0 {
		return
	}

	if isAuthenticated {
		var user *model.User
		var err error

		fmt.Println("Fetching user with ID: ", userID)
		user, err = userController.GetUser(userID) // Changed to GetUser
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user)

		fmt.Println("\nFetching all users")
		users, _ := userController.GetAllUsers() // Changed to GetAllUsers
		if err != nil {
			fmt.Println(err)
		}

		for _, user := range users {
			fmt.Println(user)
		}

		time.Sleep(5 * time.Second)

		fmt.Println("\nUpdating user data")
		err = userController.UpdateUser(userID, &model.User{
			Name:         "Nasrullo",
			Surname:      "Nurullaev",
			Login:        "nasrullonurullaev",
			HashPassword: 0,
			Status:       "Status",
			Role:         "Role",
		})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(userController.GetUser(userID)) // Changed to GetUser

		fmt.Println("Deleting user with ID: ", userID)
		err = userController.DeleteUser(userID) // Changed to DeleteUser
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("\nFetching statistics")
		fmt.Println(userController.CollectStatistics()) // Changed to CollectStatistics
	}
}
