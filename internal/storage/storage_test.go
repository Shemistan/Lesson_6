package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("operation not valid Add", func(t *testing.T) {
		_, err := NewIStorage().Add(nil)
		assert.Equal(t, err.Error(), "Error User not found")
	})
}

func TestGetUser(t *testing.T) {
	t.Run("operation not valid GetUser", func(t *testing.T) {
		_, err := NewIStorage().Get(0)
		assert.Equal(t, err.Error(), "Error User ID not Found")
	})
}

func TestUpdatetUser(t *testing.T) {
	t.Run("operation not valid Update", func(t *testing.T) {
		err := NewIStorage().Update(100, nil)
		assert.Equal(t, err.Error(), "Error User not found")
	})
}

