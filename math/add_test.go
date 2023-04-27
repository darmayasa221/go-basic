package math_test

import (
	"testing"

	"github.com/darmayasa221/go-basic/math"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Arrange
	x, y := 5, 10
	// Action
	add = math.Add(x, y)
	// Assert
	assert.Equal(t, 15, add, "add function should be return value type of int 15")
}
