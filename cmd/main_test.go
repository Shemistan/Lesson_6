package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/Shemistan/Lesson_6/internal/api"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
	"github.com/stretchr/testify/assert"

	"github.com/Shemistan/Lesson_6/internal/service"
)

type TestAuth struct {
	login            string
	password         string
	expectedResponse int64
}

/*
Zdes moi integratsionniye testi. Vozmojno ya napisal chto-to nepravilno, prosto ne sumel ponyat to, kak ih pisat nado.
V 6om uroke ne bilo primerov, a primer iz notion, bil ochen neponyaten
*/
func TestServer(t *testing.T) {
	testDB := storage.New()
	testService := service.New(testDB)
	testApi := api.New(testService)

	//auth first test
	var testAuthCase1 = TestAuth{
		login:            "Dilmurod",
		password:         "12345",
		expectedResponse: 1,
	}
	testAuthCase1Result, errCase1 := testApi.Auth(testAuthCase1.login, testAuthCase1.password)
	if errCase1 != nil {
		fmt.Printf("Auth error: %v\n", errCase1)
	}
	assert.Equal(t, testAuthCase1.expectedResponse, testAuthCase1Result, "Expected and actual responses are different")

	//auth second test
	var testAuthCase2 = TestAuth{
		login:            "Dilmurod",
		password:         "123",
		expectedResponse: 0,
	}
	testAuthCase2Result, errCase2 := testApi.Auth(testAuthCase2.login, testAuthCase2.password)
	if errCase2 != nil {
		fmt.Printf("Auth error: %v\n", errCase2)
	}
	assert.Equal(t, testAuthCase2.expectedResponse, testAuthCase2Result, "Expected and actual responses are different")

	//auth third test
	var testAuthCase3 = TestAuth{
		login:            "Dil",
		password:         "12345",
		expectedResponse: 2,
	}
	testAuthCase3Result, errCase3 := testApi.Auth(testAuthCase3.login, testAuthCase3.password)
	if errCase3 != nil {
		fmt.Printf("Auth error: %v\n", errCase3)
	}
	assert.Equal(t, testAuthCase3.expectedResponse, testAuthCase3Result, "Expected and actual responses are different")

	//update first test
	var testUpdateCase1 = models.Account{
		AuthData: models.AuthData{
			Login:    "Dilmurod",
			Password: "12345",
		},
		UserInfo: models.UserInfo{
			Name:    "Dilmurod",
			Surname: "Toshtemirov",
			Active:  true,
			Role:    "User",
		},
	}
	testApi.UpdateUser(1, testUpdateCase1)

	//update second test
	var testUpdateCase2 = models.Account{
		AuthData: models.AuthData{
			Login:    "Dil",
			Password: "12345",
		},
		UserInfo: models.UserInfo{
			Name:    "Dilmur",
			Surname: "Toshtemir",
			Active:  true,
			Role:    "Admin",
		},
	}
	testApi.UpdateUser(2, testUpdateCase2)

	//get user first test
	var testAllUsers = []models.User{
		{
			Login:            "Dilmurod",
			Password:         "12345",
			Name:             "Dilmurod",
			Surname:          "Toshtemirov",
			Active:           true,
			Role:             "User",
			RegistrationDate: time.Now().Format("02-01-2006 15:04"),
			UpdateDate:       time.Now().Format("02-01-2006 15:04"),
		},
		{
			Login:            "Dil",
			Password:         "12345",
			Name:             "Dilmur",
			Surname:          "Toshtemir",
			Active:           true,
			Role:             "Admin",
			RegistrationDate: time.Now().Format("02-01-2006 15:04"),
			UpdateDate:       time.Now().Format("02-01-2006 15:04"),
		},
	}
	testGetUserCase1Result, testGetUserCase1Error := testApi.GetUser(1)
	if testGetUserCase1Error != nil {
		fmt.Println(testGetUserCase1Error)
	}
	assert.Equal(t, &testAllUsers[0], testGetUserCase1Result, "Expected and actual responses are different")

	//get user 2nd test
	_, testGetUserCase2Error := testApi.GetUser(3)
	if testGetUserCase2Error != nil {
		fmt.Println(testGetUserCase2Error)
	}

	//get users first test
	testGetUsersCase1Result, testGetUsersCase1Error := testApi.GetUsers()
	if testGetUsersCase1Error != nil {
		fmt.Println(testGetUsersCase1Error)
	}
	assert.Equal(t, testAllUsers, testGetUsersCase1Result, "Expected and actual responses are different")

	//delete user first test
	testDeleteCase1Error := testApi.DeleteUser(1)
	if testDeleteCase1Error != nil {
		fmt.Println(testDeleteCase1Error)
	}

	//delete user second test
	testDeleteCase2Error := testApi.DeleteUser(2)
	if testDeleteCase2Error != nil {
		fmt.Println(testDeleteCase2Error)
	}

	//delete user third test
	testDeleteCase3Error := testApi.DeleteUser(3)
	if testDeleteCase3Error != nil {
		fmt.Println(testDeleteCase3Error)
	}

	//get users second test
	_, testGetUsersCase2Error := testApi.GetUsers()
	if testGetUsersCase2Error != nil {
		fmt.Println(testGetUsersCase2Error)
	}

	var testCache1 = map[string]int64{
		"Auth":          3,
		"UpdateUser":    2,
		"GetUser":       2,
		"GetUsers":      2,
		"DeleteUser":    3,
		"GetStatistics": 1,
		"DeletedUsers":  2,
	}

	testGetStatisticsCase1 := testApi.GetStatistics()
	assert.Equal(t, testCache1, testGetStatisticsCase1, "Expected and actual responses are different")
}
