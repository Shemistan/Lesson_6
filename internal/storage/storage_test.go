package storage

import (
	"errors"
	"github.com/Shemistan/Lesson_6/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	st := NewStorage()
	_, err := st.Add(&model.User{
		Name:             "test",
		Surname:          "test",
		Login:            "test",
		HashPassword:     123,
		Status:           "Active",
		Role:             "User",
		RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
		UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		t.Error(err)
	}

	_, err = st.Add(&model.User{
		Name:             "test1",
		Surname:          "test1",
		Login:            "test1",
		HashPassword:     1231,
		Status:           "Active",
		Role:             "User",
		RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
		UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		t.Error(err)
	}

	t.Run("Add", func(t *testing.T) {
		result, _ := st.Add(&model.User{
			Name:             "test2",
			Surname:          "test2",
			Login:            "test2",
			HashPassword:     1232,
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
			UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
		})
		assert.Equal(t, uint32(2), result)
	})

	t.Run("Get", func(t *testing.T) {
		_, err := st.Get(3)
		assert.Equal(t, errors.New("user not found"), err)
	})

	t.Run("Get", func(t *testing.T) {
		result, _ := st.Get(1)
		assert.Equal(t, &model.User{
			Name:             "test1",
			Surname:          "test1",
			Login:            "test1",
			HashPassword:     1231,
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
			UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
		}, result)
	})

	t.Run("GetAll", func(t *testing.T) {
		arr, _ := st.GetAll()
		assert.Equal(t, 3, len(arr))
	})

	t.Run("Update", func(t *testing.T) {
		err := st.Update(2, &model.User{
			Name:             "test1updated",
			Surname:          "test1",
			Login:            "test1",
			HashPassword:     1231,
			Status:           "Active",
			Role:             "User",
			RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
			UpdateDate:       time.Now().Format("2006-01-02 15:04:05"),
		})
		assert.Equal(t, nil, err)
	})

	t.Run("Delete", func(t *testing.T) {
		err := st.Delete(2)
		assert.Equal(t, nil, err)
	})

	t.Run("Auth", func(t *testing.T) {
		result, err := st.Auth("test", 123)
		assert.Equal(t, int32(0), result)
		assert.Equal(t, nil, err)
	})

	t.Run("Auth", func(t *testing.T) {
		result, err := st.Auth("test", 1231)
		assert.Equal(t, int32(-1), result)
		assert.Equal(t, errors.New("user not found"), err)
	})
}
