package api

import (
	"errors"
	"fmt"
	"testing"

	mock_service "github.com/Shemistan/Lesson_6/internal/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// test := &models.User{
	// 	Login:    "Dima",
	// 	Password: "Pass",
	// 	Name:     "Dilmuod",
	// 	Surname:  "Toshtemirov",
	// 	Active:   true,
	// 	Role:     "user",
	// }
	serv := mock_service.NewMockIService(ctrl)
	api := New(serv)
	var id = 0
	t.Run("req is nil", func(t *testing.T) {
		serv.EXPECT().Auth("Dima", "Pass").Return(int64(1), nil)
		serv.EXPECT().Auth("Dima", "Pas").Return(int64(0), errors.New("wrong password"))
		//db.EXPECT().UpdateUser(int(1), test.Name, test.Surname, test.Active, test.Role).Return()
		//t.Logf("Calling GetUser with id: %d", id)
		serv.EXPECT().GetUser(int64(id)).Return(nil, errors.New("wrong id"))
		serv.EXPECT().DeleteUser(int64(1)).Return(nil)
		serv.EXPECT().DeleteUser(int64(0)).Return(errors.New("no such id"))
		serv.EXPECT().GetUsers().Return(nil, errors.New("error getting users' info"))

		_, err := api.Auth("Dima", "Pass")
		if err != nil {
			t.Errorf("expected nil, got err")
		}
		_, error := api.Auth("Dima", "Pas")
		if error == nil {
			t.Errorf("expected error, got nil")
		}
		//serv.UpdateUser(1, test)
		_, err1 := api.GetUser(int64(id))
		if err1 != nil {
			fmt.Println(err1)
		}

		err3 := api.DeleteUser(1)
		if err3 != nil {
			fmt.Println(err3)
		}
		err4 := api.DeleteUser(0)
		if err4 != nil {
			fmt.Println(err4)
		}
		_, err2 := api.GetUsers()
		if err2 != nil {
			fmt.Println(err2)
		}
	})
}
