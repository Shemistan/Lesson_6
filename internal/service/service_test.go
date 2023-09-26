package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Shemistan/Lesson_6/internal/models"
	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
	"github.com/golang/mock/gomock"
)



func TestStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stor := mock_storage.NewMockIStorage(ctrl)

	serv := NewIService(stor)
	_ = serv
	t.Run("test add", func(t *testing.T) {
		t.Parallel()
		stor.EXPECT().Add(nil).Return(int32(0),errors.New("Error User not found"))

		_, err := serv.Add(nil)
		if err!=nil{
			fmt.Println(err)
		}
	})
	t.Run("test add 2", func(t *testing.T) {
		t.Parallel()
		stor.EXPECT().Add(&models.SUser{Login: "sss"}).Return(int32(0),nil)

		_, err := serv.Add(&models.SUser{Login: "sss"})
		if err!=nil{
			fmt.Println(err)
		}
	})

	t.Run("test GetUser ", func(t *testing.T) {
		stor.EXPECT().Get(int32(0)).Return(nil, errors.New("Error User ID not Found"))

		_,err := serv.GetUser(int32(0))
		if err != nil {
			fmt.Println(err)
		}
	})
	t.Run("test Update", func(t *testing.T) {
		stor.EXPECT().Update(int32(0),nil).Return(errors.New("Error User not found"))
		
		_, err := serv.UpdateUser(int32(0), nil)
		if err != nil {
			fmt.Println(err)
		}
	})

	t.Run("test Update", func(t *testing.T) {
		stor.EXPECT().Update(int32(0),nil).Return(errors.New("Error User not found"))
		
		v, err := serv.UpdateUser(int32(0), nil)
		if err != nil && v == false {
			fmt.Println(err)
		}else{
			fmt.Println(v,nil)
		}
	})
	t.Run("test Getstat", func(t *testing.T) {
		serv.GetStatistics()
	})
}

