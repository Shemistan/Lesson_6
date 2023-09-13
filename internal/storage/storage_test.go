package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/stretchr/testify/assert"
)

var tests = New()

type testAddUser struct {
	login    string
	password string
	exp      int64
}

func TestAuth(t *testing.T) {

	test := []testAddUser{
		{
			login:    "login",
			password: "password",
			exp:      1,
		},
		{
			login:    "login",
			password: "pass",
			exp:      0,
		},
		{
			login:    "login",
			password: "password",
			exp:      1,
		},
		{
			login:    "login1",
			password: "pass",
			exp:      2,
		},
	}
	for _, val := range test {
		res, err := tests.Auth(val.login, val.password)
		if err != nil {
			fmt.Println("got an error")
		} else {
			fmt.Println("pass")
		}
		assert.Equal(t, res, val.exp)
	}
}

type testUpdateUser struct {
	id      int64
	Name    string
	Surname string
	Active  bool
	Role    string
}

func TestUpdateUser(t *testing.T) {

	test := []testUpdateUser{
		{
			id:      1,
			Name:    "Dilmurod",
			Surname: "Toshtemirov",
			Active:  true,
			Role:    "User",
		},
	}
	for _, val := range test {
		tests.UpdateUser(val.id, val.Name, val.Surname, val.Active, val.Role)
	}
}

func TestGetUser(t *testing.T) {
	test := &models.User{
		Login:            "login",
		Password:         "password",
		Name:             "Dilmurod",
		Surname:          "Toshtemirov",
		Active:           true,
		Role:             "User",
		RegistrationDate: CurrentTime,
		UpdateDate:       time.Now().Format("02-01-2006 15:04"),
	}
	var id = 1
	res, err := tests.GetUser(int64(id))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("pass")
	}
	assert.Equal(t, res, test)

}

func TestGetUsers(t *testing.T) {
	test := []models.User{
		{
			Login:            "login",
			Password:         "password",
			Name:             "Dilmurod",
			Surname:          "Toshtemirov",
			Active:           true,
			Role:             "User",
			RegistrationDate: CurrentTime,
			UpdateDate:       time.Now().Format("02-01-2006 15:04"),
		},
		{
			Login:            "login1",
			Password:         "pass",
			Name:             "",
			Surname:          "",
			Active:           true,
			Role:             "User",
			RegistrationDate: CurrentTime,
			UpdateDate:       "",
		},
	}
	res, err := tests.GetUsers()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("pass")
	}
	//fmt.Printf("Actual data: %+v\n", res)
	assert.Equal(t, test, res)

}

func TestDeleteUser(t *testing.T) {
	id := 1
	var testirov error
	err := tests.DeleteUser(int64(id))
	assert.Equal(t, err, testirov)
}
