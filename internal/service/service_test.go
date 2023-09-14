package service

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var testStorage = storage.New()

var testService = New(testStorage)

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
		res, err := testService.Auth(val.login, val.password)
		if err != nil {
			fmt.Println("got an error")
		} else {
			fmt.Println("pass")
		}
		assert.Equal(t, res, val.exp)
	}
}

func TestUpdateUser(t *testing.T) {
	test := &models.User{
		Login:    "login",
		Password: "password",
		Name:     "Dilmuod",
		Surname:  "Toshtemirov",
		Active:   true,
		Role:     "user",
	}
	testService.UpdateUser(1, test)

}

func TestGetUser(t *testing.T) {
	test := &models.User{
		Login:            "login",
		Password:         "password",
		Name:             "Dilmuod",
		Surname:          "Toshtemirov",
		Active:           true,
		Role:             "user",
		RegistrationDate: time.Now().Format("02-01-2006 15:04"),
		UpdateDate:       time.Now().Format("02-01-2006 15:04"),
	}
	var id = 1
	res, err := testService.GetUser(int64(id))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("pass")
	}
	assert.Equal(t, res, test)

}

func TestGetUsers(t *testing.T) {
	test := []models.User{
		0: {
			Login:            "login",
			Password:         "password",
			Name:             "Dilmuod",
			Surname:          "Toshtemirov",
			Active:           true,
			Role:             "user",
			RegistrationDate: storage.CurrentTime,
			UpdateDate:       time.Now().Format("02-01-2006 15:04"),
		},
		1: {
			Login:            "login1",
			Password:         "pass",
			Name:             "",
			Surname:          "",
			Active:           true,
			Role:             "User",
			RegistrationDate: storage.CurrentTime,
			UpdateDate:       "",
		},
	}
	res, err := testService.GetUsers()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("pass")
	}
	//fmt.Printf("Actual data: %+v\n", res)

	assert.Equal(t, res, test)

}

func TestDeleteUser(t *testing.T) {
	id := 1
	var testirov error
	err := testService.DeleteUser(int64(id))
	assert.Equal(t, err, testirov)
}

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mock_storage.NewMockIStorage(ctrl)
	serv := New(db)
	var id = 0
	t.Run("req is nil", func(t *testing.T) {
		db.EXPECT().Auth("Dima", "Pass").Return(int64(1), nil)
		db.EXPECT().Auth("Dima", "Pas").Return(int64(0), errors.New("wrong password"))
		db.EXPECT().GetUser(int64(id)).Return(nil, errors.New("wrong id"))
		db.EXPECT().DeleteUser(int64(1)).Return(nil)
		db.EXPECT().DeleteUser(int64(0)).Return(errors.New("no such id"))
		db.EXPECT().GetUsers().Return(nil, errors.New("error getting users' info"))

		_, err := serv.Auth("Dima", "Pass")
		if err != nil {
			t.Errorf("expected nil, got err")
		}
		_, error := serv.Auth("Dima", "Pas")
		if error == nil {
			t.Errorf("expected error, got nil")
		}
		//serv.UpdateUser(1, test)
		_, err1 := serv.GetUser(int64(id))
		if err1 != nil {
			fmt.Println(err1)
		}

		err3 := serv.DeleteUser(1)
		if err3 != nil {
			fmt.Println(err3)
		}
		err4 := serv.DeleteUser(0)
		if err4 != nil {
			fmt.Println(err4)
		}
		_, err2 := serv.GetUsers()
		if err2 != nil {
			fmt.Println(err2)
		}
	})
}
