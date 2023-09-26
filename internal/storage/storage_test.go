package storage

import (
	"fmt"
	"testing"

	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("operation not valid Add", func(t *testing.T) {
		_, err := NewIStorage().Add(nil)
		assert.Equal(t, err.Error(), "Error User not found")
	})
	t.Run("operation not valid Add 2", func(t *testing.T) {
		v, err := NewIStorage().Add(&models.SUser{Login: "sss"})
		if v == 0 && err == nil {
			return
		}
	})
	t.Run("operation not valid Add 3", func(t *testing.T) {
		v, err := NewIStorage().Add(&models.SUser{Login: "sss"})
		if v > 0 && err == nil {
			return
		}
	})
}


func TestGetUser(t *testing.T) {
	t.Run("operation not valid GetUser", func(t *testing.T) {
		_, err := NewIStorage().Get(0)
		assert.Equal(t, err.Error(), "Error cant get user")
	})
	t.Run("operation not valid GetUser 2", func(t *testing.T) {
		_, err := NewIStorage().Get(0)
		if err == nil {
			fmt.Println("good")
		}
	})
}

func TestUpdatetUser(t *testing.T) {
	t.Run("operation not valid Update", func(t *testing.T) {
		err := NewIStorage().Update(100, nil)
		assert.Equal(t, err.Error(), "Error User not found")
	})
	t.Run("operation not valid Update 2", func(t *testing.T) {
		err := NewIStorage().Update(100, &models.SUser{Login: "ssss"})
		assert.Equal(t, err.Error(), "Error USerID not found")
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("operation not valid delete", func(t *testing.T) {
		err := NewIStorage().Delete(0)
		assert.Equal(t, err.Error(), "Error Cant Delete User")
		
	})
}

func TestGetAlleUser(t *testing.T) {
	t.Run("operation not valid Getall", func(t *testing.T) {
		NewIStorage().GetAll()
	})
}

func TestGetMap(t *testing.T) {
	t.Run("operation not valid GetMAp", func(t *testing.T) {
		v := NewIStorage().GetMap()
		if v != nil {
			return
		}
	})
}
