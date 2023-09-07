package service

import (
	"errors"
	"fmt"
	"testing"

	mock_storage "github.com/Shemistan/Lesson_6/internal/storage/mocks"
	"github.com/golang/mock/gomock"
)



func TestStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stor := mock_storage.NewMockIStorage(ctrl)

	serv := NewIService(stor)
	_ = serv
	t.Run("request is nil", func(t *testing.T) {
		stor.EXPECT().Add(nil).Return(int32(0),errors.New("Error User not found"))

		_, err := serv.Add(nil)
		if err!=nil{
			errors.New("Error User not found2")
		}
	
	
		t.Run("test Update", func(t *testing.T) {
			stor.EXPECT().Update(int32(0),nil).Return(errors.New("Error User not found"))
			// stor.EXPECT().Update(int32(10),models.SUser{Login: "tttt", Surname: "sssss"}).Return(errors.New("Error USerID not found"))
			_, err := serv.UpdateUser(int32(0), nil)
			if err != nil {
				fmt.Println(err)
			}

				t.Run("test Get", func(t *testing.T) {
					stor.EXPECT().Get(int32(0)).Return(nil, errors.New("Error User ID not Found"))
					_,err := serv.GetUser(int32(0))
					if err != nil {
						fmt.Println(err)
					}
				})
		})

	})
}