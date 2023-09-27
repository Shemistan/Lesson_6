package main

import (
	"fmt"

	"github.com/Shemistan/Lesson_6/internal/api"
	"github.com/Shemistan/Lesson_6/internal/models"
)

func TestMain(handlers api.IApi) {

	authTest1 := &models.AddRequest{
		AuthParams: models.UserAuth{
			Login:    "login",
			Password: "password",
		},
		Data: models.UserData{
			Name:    "Username",
			Surname: "usersurname",
		},
	}

	authTest2 := &models.AddRequest{
		AuthParams: models.UserAuth{
			Login:    "login",
			Password: "password",
		},
		Data: models.UserData{
			Name:    "Username",
			Surname: "usersurname",
		},
	}

	authTest3 := &models.AddRequest{
		AuthParams: models.UserAuth{
			Login:    "anotherlogin",
			Password: "anotherpassword",
		},
		Data: models.UserData{
			Name:    "Username2",
			Surname: "usersurname2",
		},
	}

	updTestId1 := 2
	updTestId2 := 5
	updTest :=
		&models.UpdateUserData{
			Name:    "Name",
			Surname: "Surname",
			Status:  "Active",
			Role:    "User",
		}

	getUserId1 := 1
	getUserId2 := 5
	deleteUser := 2

	// db := storage.New()
	// serv := service.New(db)
	// handlers := api.New(serv)

	resAuth, err := handlers.Auth(authTest1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("AuthTest1: result = %v\n", resAuth)

	resAuth, err = handlers.Auth(authTest2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("AuthTest2: result = %v\n", resAuth)

	resAuth, err = handlers.Auth(authTest3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("AuthTest1: result = %v\n", resAuth)

	resStat := handlers.GetStatistics()
	fmt.Printf("Stat after Auth: %v\n", resStat)

	err = handlers.UpdateUser(int64(updTestId1), updTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("updated success")
	}

	err = handlers.UpdateUser(int64(updTestId2), updTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("updated success")
	}

	resStat = handlers.GetStatistics()
	fmt.Printf("Stat after update: %v\n", resStat)

	resGet, err := handlers.GetUser(int64(getUserId1))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("getTest1: result = %v\n", resGet)

	resGet, err = handlers.GetUser(int64(getUserId2))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("getTest2: result = %v\n", resGet)

	resStat = handlers.GetStatistics()
	fmt.Printf("Stat after Get: %v\n", resStat)

	resGets, err := handlers.GetUsers()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("GetUsersTest: result = %v\n", resGets)

	resStat = handlers.GetStatistics()
	fmt.Printf("Stat after GetAll: %v\n", resStat)

	err = handlers.DeleteUser(int64(deleteUser))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("deleted success")
	}

	err = handlers.DeleteUser(int64(deleteUser))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("deleted success")
	}

	resStat = handlers.GetStatistics()
	fmt.Printf("Final stat: %v\n", resStat)
}
