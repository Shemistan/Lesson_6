package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("test sum", func(t *testing.T) {
		res := sum(1, 2)
		assert.Equal(t, res, 3)
	})

	t.Run("test sub", func(t *testing.T) {
		res := sum(1, 2)
		assert.Equal(t, res, 5)
	})

	t.Run("test sub", func(t *testing.T) {
		res := sum(1, 4)
		assert.Equal(t, res, 7)
	})

}
